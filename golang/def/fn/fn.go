package fn

import (
	"fmt"
	"go/ast"
	"go/types"
	"strconv"
	"strings"

	"github.com/matthewmueller/golly/golang/def/jsfile"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/index"

	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/golang/util"

	"golang.org/x/tools/go/loader"
)

// Function interface
type Function interface {
	def.Definition
	IsAsync() (bool, error)
	Node() *ast.FuncDecl
}

type funcdef struct {
	info      *loader.PackageInfo
	index     *index.Index
	id        string
	path      string
	name      string
	kind      types.Type
	node      *ast.FuncDecl
	exported  bool
	params    []string
	tag       *structtag.Tag
	runtime   bool
	processed bool
	deps      []def.Definition
	async     bool
	imports   map[string]string
}

var _ Function = (*funcdef)(nil)

// NewFunction fn
func NewFunction(index *index.Index, info *loader.PackageInfo, n *ast.FuncDecl) (def.Definition, error) {
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	name := n.Name.Name
	idParts := []string{packagePath, name}
	id := strings.Join(idParts, " ")

	var params []string
	for _, param := range n.Type.Params.List {
		for _, ident := range param.Names {
			params = append(params, ident.Name)
		}
	}

	// if it's a method don't export,
	// if it's the main() function
	// export either way
	exported := obj.Exported()
	if n.Recv != nil {
		exported = false
	} else if name == "main" {
		exported = true
	}

	fromRuntime := false
	runtimePath, e := util.RuntimePath()
	if e != nil {
		return nil, e
	}
	if packagePath == runtimePath {
		fromRuntime = true
	}

	tag, e := util.JSTag(n.Doc)
	if e != nil {
		return nil, e
	}

	return &funcdef{
		index:    index,
		info:     info,
		id:       id,
		exported: exported,
		path:     packagePath,
		name:     name,
		node:     n,
		params:   params,
		kind:     info.TypeOf(n.Name),
		runtime:  fromRuntime,
		tag:      tag,
		imports:  map[string]string{},
	}, nil
}

func (d *funcdef) process() (err error) {
	seen := map[string]bool{}

	ast.Inspect(d.node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.FuncDecl:
			tag, e := util.JSTag(t.Doc)
			if e != nil {
				err = e
				return false
			}
			d.tag = tag

		case *ast.SelectorExpr:
			def, e := d.index.DefinitionOf(d.Path(), t)
			if e != nil {
				err = e
				return false
			} else if def != nil && !seen[def.ID()] {
				d.deps = append(d.deps, def)
				seen[def.ID()] = true

				// check if it points to something async
				if e := d.maybeAsync(def); e != nil {
					err = e
					return false
				}
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
				if e := d.maybeAsync(def); e != nil {
					err = e
					return false
				}
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

func (d *funcdef) ID() string {
	return d.id
}

func (d *funcdef) Name() string {
	if d.tag != nil {
		return d.tag.Name
	}
	return d.name
}

func (d *funcdef) Path() string {
	return d.path
}

func (d *funcdef) Dependencies() (defs []def.Definition, err error) {
	if d.processed {
		return d.deps, nil
	}
	e := d.process()
	if e != nil {
		return defs, e
	}
	return d.deps, nil
}

func (d *funcdef) Exported() bool {
	return d.exported
}

func (d *funcdef) Omitted() bool {
	if d.tag == nil {
		return false
	}
	return d.tag.HasOption("omit")
}

func (d *funcdef) Node() *ast.FuncDecl {
	return d.node
}

func (d *funcdef) Type() types.Type {
	return d.kind
}

func (d *funcdef) IsAsync() (bool, error) {
	if d.processed {
		return d.async, nil
	}
	e := d.process()
	if e != nil {
		return false, e
	}
	return d.async, nil
}

func (d *funcdef) Imports() map[string]string {
	// combine def imports with file imports
	imports := map[string]string{}
	for alias, path := range d.imports {
		imports[alias] = path
	}
	for alias, path := range d.index.GetImports(d.path) {
		imports[alias] = path
	}
	return imports
}

func (d *funcdef) FromRuntime() bool {
	return d.runtime
}

func (d *funcdef) maybeAsync(def def.Definition) error {
	if d.async || d.ID() == def.ID() {
		return nil
	}

	fn, ok := def.(Function)
	if !ok {
		return nil
	}

	async, e := fn.IsAsync()
	if e != nil {
		return e
	}
	d.async = async

	return nil
}
