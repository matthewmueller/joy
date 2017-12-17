package indexer

import (
	"fmt"
	"go/ast"
	"path"
	"strings"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/compiler/defs"
	"github.com/matthewmueller/joy/internal/compiler/index"
	"github.com/matthewmueller/joy/internal/compiler/util"
	"golang.org/x/tools/go/loader"
)

// Indexer struct
type Indexer struct {
	index *index.Index
}

// New fn
func New(program *loader.Program) (idx *index.Index, err error) {
	// defer log.Trace("index").Stop(&err)

	index := &Indexer{
		index: index.New(program),
	}

	for _, info := range program.AllPackages {
		pkgpath := info.Pkg.Path()
		// ignore the joy/macro package path when indexing
		// TODO: this is a HUGE HACK and should be eliminated
		// altogether. It shouldn't matter whether we include macro
		// in the index, but right now index.DefinitionOf will return
		// macro.File rather than the underlying type so rewrites
		// currently get messed up
		if strings.HasSuffix(pkgpath, "joy/macro") {
			continue
		}

		if e := index.pkg(info); e != nil {
			return nil, e
		}
	}

	for id, def := range index.index.All() {
		log.Debugf("index key: %s (%T)", id, def)
	}

	return index.index, nil
}

func (indexer *Indexer) pkg(info *loader.PackageInfo) (err error) {
	for _, file := range info.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			recurse, e := indexer.inspect(info, node)
			if e != nil {
				err = e
				return false
			}
			return recurse
		})
	}
	return err
}

func (indexer *Indexer) inspect(info *loader.PackageInfo, node ast.Node) (recurse bool, err error) {
	switch t := node.(type) {
	case *ast.FuncDecl:
		if t.Recv != nil {
			m, e := defs.Method(indexer.index, info, t)
			if e != nil {
				return false, e
			}

			indexer.index.AddDefinition(m)
			return false, nil
		}

		d, e := defs.Function(indexer.index, info, t)
		if e != nil {
			return false, e
		}

		indexer.index.AddDefinition(d)
		return false, nil
	case *ast.GenDecl:
		for _, spec := range t.Specs {
			e := indexer.spec(info, t, spec)
			if e != nil {
				return false, e
			}
		}
		return false, nil
	default:
		return true, nil
	}
}

func (indexer *Indexer) spec(info *loader.PackageInfo, n *ast.GenDecl, spec ast.Spec) error {
	switch t := spec.(type) {
	case *ast.ValueSpec:
		return indexer.valueSpec(info, n, t)
	case *ast.ImportSpec:
		return indexer.importSpec(info, n, t)
	case *ast.TypeSpec:
		return indexer.typeSpec(info, n, t)
	default:
		return fmt.Errorf("unhandled spec: %T", spec)
	}
}

func (indexer *Indexer) typeSpec(info *loader.PackageInfo, gn *ast.GenDecl, n *ast.TypeSpec) error {
	switch t := n.Type.(type) {
	case *ast.StructType:
		st, e := defs.Struct(indexer.index, info, gn, n)
		if e != nil {
			return e
		}
		indexer.index.AddDefinition(st)
		return nil
	case *ast.InterfaceType:
		iface, e := defs.Interface(indexer.index, info, gn, n)
		if e != nil {
			return e
		}
		indexer.index.AddDefinition(iface)

		methods := util.MethodsFromInterface(t, iface.Path(), iface.Name())
		for _, method := range methods {
			indexer.index.Link(method, iface)
		}

		return nil
	case *ast.Ident, *ast.SelectorExpr, *ast.MapType, *ast.ArrayType:
		typedef, e := defs.Type(indexer.index, info, gn, n)
		if e != nil {
			return e
		}
		indexer.index.AddDefinition(typedef)
		return nil
	default:
		return fmt.Errorf("unhandled typeSpec: %T", n.Type)
	}
}

func (indexer *Indexer) importSpec(info *loader.PackageInfo, gn *ast.GenDecl, n *ast.ImportSpec) error {
	// trim the "" of package path's
	deppath := strings.Trim(n.Path.Value, `"`)

	// create a package path
	packagePath := info.Pkg.Path()

	// alias
	if n.Name != nil {
		indexer.index.AddImport(packagePath, n.Name.Name, deppath)
		return nil
	}

	// use base name
	name := path.Base(deppath)
	indexer.index.AddImport(packagePath, name, deppath)
	return nil
}

func (indexer *Indexer) valueSpec(info *loader.PackageInfo, gn *ast.GenDecl, n *ast.ValueSpec) error {
	d, e := defs.Value(indexer.index, info, gn, n)
	if e != nil {
		return e
	}
	indexer.index.AddDefinition(d)
	return nil
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
