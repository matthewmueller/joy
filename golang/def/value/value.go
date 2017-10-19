package value

import (
	"fmt"
	"go/ast"
	"go/types"
	"strconv"
	"strings"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/def/jsfile"
	"github.com/matthewmueller/golly/golang/index"
	"github.com/matthewmueller/golly/golang/util"

	"golang.org/x/tools/go/loader"
)

// Value interface
type Value interface {
	def.Definition
	Node() *ast.ValueSpec
	// IsAsync() bool, error
}

var _ Value = (*valuedef)(nil)

type valuedef struct {
	exported  bool
	path      string
	name      string
	id        string
	index     *index.Index
	node      *ast.ValueSpec
	kind      types.Type
	processed bool
	deps      []def.Definition
	imports   map[string]string
	async     bool
}

// NewValue fn
func NewValue(index *index.Index, info *loader.PackageInfo, gn *ast.GenDecl, n *ast.ValueSpec) (def.Definition, error) {
	packagePath := info.Pkg.Path()
	names := []string{}
	exported := false

	for _, ident := range n.Names {
		obj := info.ObjectOf(ident)
		if obj.Exported() {
			exported = true
		}
		names = append(names, ident.Name)
	}

	name := strings.Join(names, ",")
	idParts := []string{packagePath, name}
	id := strings.Join(idParts, " ")

	return &valuedef{
		exported: exported,
		path:     packagePath,
		name:     name,
		id:       id,
		index:    index,
		node:     n,
		imports:  map[string]string{},
	}, nil
}

func (d *valuedef) process() (err error) {
	seen := map[string]bool{}

	ast.Inspect(d.node, func(n ast.Node) bool {
		switch t := n.(type) {

		case *ast.SelectorExpr:
			def, e := d.index.DefinitionOf(d.Path(), t)
			if e != nil {
				err = e
				return false
			} else if def != nil && !seen[def.ID()] {
				d.deps = append(d.deps, def)
				seen[def.ID()] = true

				// check if it points to something async
				// if e := d.maybeAsync(def); e != nil {
				// 	err = e
				// 	return false
				// }
			}

		case *ast.CallExpr:
			fn, e := util.ExprToString(t.Fun)
			if e != nil {
				err = e
				return false
			}
			if fn == "js.RawFile" && len(t.Args) > 0 {
				lit, ok := t.Args[0].(*ast.BasicLit)
				if !ok {
					err = fmt.Errorf("fn process: expected rawfile to have basiclit argument, but got %T", t.Args[0])
					return false
				}
				path, e := strconv.Unquote(lit.Value)
				if e != nil {
					err = e
					return false
				}
				def, e := jsfile.NewFile(d.path, path)
				if e != nil {
					err = e
					return false
				}

				if !seen[def.ID()] {
					d.index.AddDefinition(def)
					d.deps = append(d.deps, def)
					seen[def.ID()] = true
				}
			}

		case *ast.Ident:
			def, e := d.index.DefinitionOf(d.Path(), t)
			if e != nil {
				err = e
				return false
			} else if def != nil && !seen[def.ID()] {
				d.deps = append(d.deps, def)
				seen[def.ID()] = true

				// check if it points to something async
				// if e := d.maybeAsync(def); e != nil {
				// 	err = e
				// 	return false
				// }
			}

		case *ast.ChanType:
			deps, e := d.index.Runtime("Deferred", "Channel", "Send", "Recv")
			if e != nil {
				err = e
				return false
			}
			d.deps = append(d.deps, deps...)
			for _, dep := range deps {
				d.imports["runtime"] = dep.Path()
				seen[dep.ID()] = true
			}

			// TODO: more specific later
			d.async = true
		}

		return true
	})

	d.processed = true
	return err
}

func (d *valuedef) ID() string {
	return d.id
}

func (d *valuedef) Name() string {
	return d.name
}

func (d *valuedef) Kind() string {
	return "VALUE"
}

func (d *valuedef) Path() string {
	return d.path
}

func (d *valuedef) Dependencies() (defs []def.Definition, err error) {
	if d.processed {
		return d.deps, nil
	}
	e := d.process()
	if e != nil {
		return defs, e
	}
	return d.deps, nil
}

func (d *valuedef) Exported() bool {
	return d.exported
}

func (d *valuedef) Omitted() bool {
	return false
}

func (d *valuedef) Node() *ast.ValueSpec {
	return d.node
}

func (d *valuedef) Type() types.Type {
	return d.kind
}

func (d *valuedef) Imports() map[string]string {
	return d.index.GetImports(d.path)
}

func (d *valuedef) FromRuntime() bool {
	return false
}
