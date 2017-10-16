package db

import (
	"fmt"
	"go/ast"
	"go/types"
	"path"
	"strings"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/util"
	"golang.org/x/tools/go/loader"
)

// DB struct
type DB struct {
	imports map[string]map[string]string
	index   map[string]def.Definition
}

type functionDef struct {
	id   string
	name string
	node *ast.FuncDecl
}

// New fn
func New(program *loader.Program) (*DB, error) {
	db := &DB{
		imports: map[string]map[string]string{},
		index:   map[string]def.Definition{},
	}

	for _, info := range program.AllPackages {
		if e := db.pkg(info); e != nil {
			return nil, e
		}
	}

	for id := range db.index {
		log.Infof("key %s", id)
	}

	return db, nil
}

func (db *DB) pkg(info *loader.PackageInfo) (err error) {
	for _, file := range info.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			recurse, e := db.inspect(info, node)
			if e != nil {
				err = e
				return false
			}
			return recurse
		})
	}
	return err
}

func (db *DB) inspect(info *loader.PackageInfo, node ast.Node) (recurse bool, err error) {
	switch t := node.(type) {
	case *ast.FuncDecl:
		if t.Recv != nil {
			m, e := def.NewMethod(db.index, info, t)
			if e != nil {
				return false, e
			}
			db.index[m.ID()] = m
			return false, nil
		}

		d, e := def.NewFunction(db.index, info, t)
		if e != nil {
			return false, e
		}
		db.index[d.ID()] = d
		return false, nil
	case *ast.ValueSpec:
		for _, name := range t.Names {
			if name.Name == "_" {
				continue
			}
			d, e := def.NewValue(db.index, info, name)
			if e != nil {
				return false, e
			}
			db.index[d.ID()] = d
		}
		return false, nil
	case *ast.TypeSpec:
		e := db.parseType(info, t)
		if e != nil {
			return false, e
		}
		return false, nil
	case *ast.ImportSpec:
		db.parseImport(info, t)
		return false, nil
	default:
		return true, nil
	}
}

func (db *DB) parseType(info *loader.PackageInfo, n *ast.TypeSpec) error {
	switch t := n.Type.(type) {
	case *ast.StructType:
		st, e := def.NewStruct(db.index, info, n)
		if e != nil {
			return e
		}
		db.index[st.ID()] = st
		return nil
	case *ast.InterfaceType:
		iface, e := def.NewInterface(db.index, info, n)
		if e != nil {
			return e
		}
		db.index[iface.ID()] = iface

		// link the methods to the interface declaration
		for _, m := range t.Methods.List {
			for _, name := range m.Names {
				mid := iface.Path() + " " + name.Name + " " + iface.Name()
				db.index[mid] = iface
			}
		}
		return nil
	default:
		return fmt.Errorf("unhandled %T", n.Type)
	}
}

func (db *DB) parseImport(info *loader.PackageInfo, n *ast.ImportSpec) {
	// trim the "" of package path's
	deppath := strings.Trim(n.Path.Value, `"`)

	// create a package path
	packagePath := info.Pkg.Path()
	if db.imports[packagePath] == nil {
		db.imports[packagePath] = map[string]string{}
	}

	// alias
	if n.Name != nil {
		db.imports[packagePath][n.Name.Name] = deppath
		return
	}

	// use base name
	name := path.Base(deppath)
	db.imports[packagePath][name] = deppath
	return
}

// DefinitionOf fn
// Will return nil when ident is:
// - package name
// - basic type
// - local variable that points to a basic type
// - function parameters & results
func (db *DB) DefinitionOf(info *loader.PackageInfo, n ast.Node) (def.Definition, error) {
	if s, ok := n.(*ast.SelectorExpr); ok {
		d, e := db.DefinitionOf(info, s.Sel)
		if e != nil {
			return nil, e
		} else if d != nil {
			return d, nil
		}
		return db.DefinitionOf(info, s.X)
	}

	ident, e := util.GetIdentifier(n)
	if e != nil {
		return nil, e
	}

	// built-in or package name
	obj := info.ObjectOf(ident)
	if obj == nil {
		return nil, nil
	}

	// package is nil for types (e.g int)
	// & built-in functions (e.g. append)
	pkg := obj.Pkg()
	if pkg == nil {
		return nil, nil
	}

	// first try getting the definition from
	// the package and object name
	d := db.index[pkg.Path()+" "+obj.Name()]
	if d != nil {
		return d, nil
	}

	// lookup using the type. Useful for:
	// - local variables
	// - methods (funcs w/ receivers)
	ids, err := db.typeToDef(obj.Name(), obj.Type())
	if err != nil {
		return nil, err
	}

	id := strings.Join(ids, " ")
	return db.index[id], nil
}

func (db *DB) typeToDef(name string, kind types.Type) (arr []string, err error) {
	switch t := kind.(type) {
	case *types.Basic:
		return arr, nil
	case *types.Interface:
		if t.Empty() {
			return arr, nil
		}
		return arr, fmt.Errorf("typeToDef: unhandled non-interface %s", t)
	case *types.Slice:
		return db.typeToDef(name, t.Elem())
	case *types.Signature:
		recv := t.Recv()
		// TODO: clean this up
		if recv != nil {
			a, e := db.typeToDef(name, recv.Type())
			if e != nil {
				return arr, e
			} else if len(a) == 2 {
				return []string{a[0], name, a[1]}, nil
			}
			return arr, nil
		}
	case *types.Named:
		obj := t.Obj()
		if obj == nil {
			return arr, fmt.Errorf("typeToDef: unhandled unnamed types.Named: %s", t.String())
		}
		// pkg will be nil for Error in error.Error()
		if obj.Pkg() == nil {
			return arr, nil
		}
		return append(arr, obj.Pkg().Path(), obj.Name()), nil
	case *types.Pointer:
		return db.typeToDef(name, t.Elem())
	default:
		log.Infof("typeToDef: unhandled type %T", kind)
	}

	// log.Infof("ident %+v", i/d)
	return arr, nil
}
