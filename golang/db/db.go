package db

import (
	"fmt"
	"go/ast"
	"path"
	"strings"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/golang/def/fn"
	"github.com/matthewmueller/golly/golang/def/iface"
	"github.com/matthewmueller/golly/golang/def/method"
	"github.com/matthewmueller/golly/golang/def/struc"
	"github.com/matthewmueller/golly/golang/def/value"
	"github.com/matthewmueller/golly/golang/index"
	"github.com/matthewmueller/golly/golang/util"
	"golang.org/x/tools/go/loader"
)

// DB struct
type DB struct {
	index *index.Index
	// imports map[string]map[string]string
	// index   map[string]def.Definition
}

// New fn
func New(program *loader.Program) (*index.Index, error) {
	db := &DB{
		index: index.New(program),
	}

	jsPath, e := util.JSSourcePath()
	if e != nil {
		return nil, e
	}

	for _, info := range program.AllPackages {
		// ignore the golly/js package path
		if info.Pkg.Path() == jsPath {
			continue
		}

		if e := db.pkg(info); e != nil {
			return nil, e
		}
	}

	for id, def := range db.index.All() {
		log.Debugf("index key: %s (%T)", id, def)
	}

	return db.index, nil
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
			m, e := method.NewMethod(db.index, info, t)
			if e != nil {
				return false, e
			}

			db.index.AddDefinition(m)
			return false, nil
		}

		d, e := fn.NewFunction(db.index, info, t)
		if e != nil {
			return false, e
		}

		db.index.AddDefinition(d)
		return false, nil
	case *ast.ValueSpec:
		d, e := value.NewValue(db.index, info, t)
		if e != nil {
			return false, e
		}
		db.index.AddDefinition(d)

		// for _, name := range t.Names {
		// 	if name.Name == "_" {
		// 		continue
		// 	}
		// }
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
		st, e := struc.NewStruct(db.index, info, n)
		if e != nil {
			return e
		}
		db.index.AddDefinition(st)
		return nil
	case *ast.InterfaceType:
		iface, e := iface.NewInterface(db.index, info, n)
		if e != nil {
			return e
		}
		db.index.AddDefinition(iface)

		// link the methods to the interface declaration
		for _, m := range t.Methods.List {
			for _, name := range m.Names {
				mid := iface.Path() + " " + name.Name + " " + iface.Name()
				db.index.Link(mid, iface)
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

	// alias
	if n.Name != nil {
		db.index.AddImport(packagePath, n.Name.Name, deppath)
		return
	}

	// use base name
	name := path.Base(deppath)
	db.index.AddImport(packagePath, name, deppath)
	return
}

// // DefinitionOf fn
// func (db *DB) DefinitionOf(info *loader.PackageInfo, n ast.Node) (d def.Definition, err error) {
// 	return db.index.DefinitionOf(info, n)
// }

// // TypeOf fn
// func (db *DB) TypeOf(info *loader.PackageInfo, n ast.Node) (kind types.Type, err error) {
// 	return db.index.TypeOf(info, n)
// }

// // Mains gets all the main functions
// func (db *DB) Mains() (defs []def.Definition) {
// 	return db.index.Mains()
// 	// for _, info := range db.program.InitialPackages() {
// 	// 	p := info.Pkg.Path()
// 	// 	def := db.index.Get(p + " main")
// 	// 	if def == nil {
// 	// 		continue
// 	// 	}
// 	// 	defs = append(defs, def)
// 	// }
// 	// return defs
// }
