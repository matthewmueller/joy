package index

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/util"
	"golang.org/x/tools/go/loader"
)

// Index type
type Index struct {
	program *loader.Program
	defs    map[string]def.Definition
	imports map[string]map[string]string
}

// New index
func New(program *loader.Program) *Index {
	return &Index{
		program: program,
		defs:    map[string]def.Definition{},
		imports: map[string]map[string]string{},
	}
}

// AddDefinition adds a definition to the index
func (i *Index) AddDefinition(d def.Definition) {
	i.defs[d.ID()] = d
}

// AddImport fn
func (i *Index) AddImport(parentPath string, alias, depPath string) {
	if i.imports[parentPath] == nil {
		i.imports[parentPath] = map[string]string{}
	}
	i.imports[parentPath][alias] = depPath
}

// All gets all definitions from the index
// TODO: remove
func (i *Index) All() map[string]def.Definition {
	return i.defs
}

// Get all definitions from the index
func (i *Index) Get(id string) def.Definition {
	return i.defs[id]
}

// Link is like a symlink to a definition
func (i *Index) Link(alias string, d def.Definition) {
	i.defs[alias] = d
}

// Mains gets all the main functions
func (i *Index) Mains() (defs []def.Definition) {
	for _, info := range i.program.InitialPackages() {
		p := info.Pkg.Path()
		def := i.defs[p+" main"]
		if def == nil {
			continue
		}
		defs = append(defs, def)
	}
	return defs
}

// DefinitionOf fn
// Will return nil when ident is:
// - package name
// - basic type
// - local variable that points to a basic type
// - function parameters & results
func (i *Index) DefinitionOf(info *loader.PackageInfo, n ast.Node) (def.Definition, error) {
	if s, ok := n.(*ast.SelectorExpr); ok {
		d, e := i.DefinitionOf(info, s.Sel)
		if e != nil {
			return nil, e
		} else if d != nil {
			return d, nil
		}
		return i.DefinitionOf(info, s.X)
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
	d := i.defs[pkg.Path()+" "+obj.Name()]
	if d != nil {
		return d, nil
	}

	// lookup using the type. Useful for:
	// - local variables
	// - methods (funcs w/ receivers)
	ids, err := i.typeToDef(obj.Name(), obj.Type())
	if err != nil {
		return nil, err
	}

	id := strings.Join(ids, " ")
	return i.Get(id), nil
}

// TypeOf fn
func (i *Index) TypeOf(packagePath string, n ast.Node) (types.Type, error) {
	info := i.program.Package(packagePath)

	id, e := util.GetIdentifier(n)
	if e != nil {
		return nil, e
	}

	return info.TypeOf(id), nil
}

func (i *Index) typeToDef(name string, kind types.Type) (arr []string, err error) {
	switch t := kind.(type) {
	case *types.Basic:
		return arr, nil
	case *types.Interface:
		if t.Empty() {
			return arr, nil
		}
		return arr, fmt.Errorf("typeToDef: unhandled non-interface %s", t)
	case *types.Slice:
		return i.typeToDef(name, t.Elem())
	case *types.Signature:
		recv := t.Recv()
		// TODO: clean this up
		if recv != nil {
			a, e := i.typeToDef(name, recv.Type())
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
		return i.typeToDef(name, t.Elem())
	case *types.Map:
		// TODO: is this always the case?
		return arr, nil
	default:
		return arr, fmt.Errorf("typeToDef: unhandled type %T", kind)
	}

	// log.Infof("ident %+v", i/d)
	return arr, nil
}

// func (db *DB) TypeOf(info *loader.PackageInfo, n ast.Node) ()
