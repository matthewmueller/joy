package defs

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/loader"

	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/index"
	"github.com/matthewmueller/golly/golang/util"
)

// Functioner interface
type Functioner interface {
	def.Definition
	IsAsync() (bool, error)
	IsVariadic() bool
	Node() *ast.FuncDecl
	Rewrite() def.Rewrite
	Params() []string
	Results() []def.FunctionResult
}

var _ Functioner = (*functions)(nil)

type functions struct {
	info      *loader.PackageInfo
	index     *index.Index
	id        string
	path      string
	name      string
	kind      types.Type
	node      *ast.FuncDecl
	exported  bool
	tag       *structtag.Tag
	runtime   bool
	processed bool
	deps      []def.Definition
	rewrite   def.Rewrite
	results   []def.FunctionResult
	async     bool
	imports   map[string]string
	omit      bool
	params    []string
	variadic  bool
}

type result struct {
	name string
	def  def.Definition
}

func (r *result) Name() string {
	return r.name
}

func (r *result) Definition() def.Definition {
	return r.def
}

// Function fn
func Function(index *index.Index, info *loader.PackageInfo, n *ast.FuncDecl) (def.Definition, error) {
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	name := n.Name.Name
	idParts := []string{packagePath, name}
	id := strings.Join(idParts, " ")

	// TODO: scoping
	// scope := info.Scopes[n.Type]

	var params []string
	var variadic bool
	for _, param := range n.Type.Params.List {
		for _, ident := range param.Names {
			params = append(params, ident.Name)
		}
		if _, ok := param.Type.(*ast.Ellipsis); ok {
			variadic = true
		}
	}

	var results []def.FunctionResult
	if n.Type.Results != nil {
		for _, r := range n.Type.Results.List {
			def, err := index.DefinitionOf(packagePath, r.Type)
			if err != nil {
				return nil, err
			}

			if len(r.Names) == 0 {
				results = append(results, &result{
					def: def,
				})
			}

			for _, id := range r.Names {
				results = append(results, &result{
					name: id.Name,
					def:  def,
				})
			}
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

	return &functions{
		index:    index,
		info:     info,
		id:       id,
		exported: exported,
		path:     packagePath,
		name:     name,
		node:     n,
		kind:     info.TypeOf(n.Name),
		runtime:  fromRuntime,
		imports:  map[string]string{},
		params:   params,
		results:  results,
		variadic: variadic,
		tag:      tag,
	}, nil
}

func (d *functions) process() (err error) {
	state, e := process(d.index, d, d.node)
	if e != nil {
		return e
	}

	// copy state into function
	d.processed = true
	d.async = state.async
	d.deps = state.deps
	d.imports = state.imports
	d.omit = state.omit
	d.rewrite = state.rewrite

	if d.tag == nil {
		d.tag = state.tag
	}

	return nil
}

func (d *functions) ID() string {
	return d.id
}

func (d *functions) Name() string {
	if d.tag != nil {
		return d.tag.Name
	}
	return d.name
}

func (d *functions) Path() string {
	return d.path
}

func (d *functions) Dependencies() (deps []def.Definition, err error) {
	if d.processed {
		return d.deps, nil
	}
	e := d.process()
	if e != nil {
		return deps, e
	}

	return d.deps, nil
}

func (d *functions) Exported() bool {
	return d.exported
}

func (d *functions) Omitted() bool {
	if d.tag != nil && d.tag.HasOption("omit") {
		return true
	}
	return d.omit
}

func (d *functions) Node() *ast.FuncDecl {
	return d.node
}

func (d *functions) Type() types.Type {
	return d.kind
}

func (d *functions) Kind() string {
	return "FUNCTION"
}

func (d *functions) IsAsync() (bool, error) {
	if d.processed {
		return d.async, nil
	}
	e := d.process()
	if e != nil {
		return false, e
	}
	return d.async, nil
}

// Rewrite fn
func (d *functions) Rewrite() def.Rewrite {
	return d.rewrite
}

// Params fn
func (d *functions) Params() []string {
	return d.params
}

func (d *functions) Imports() map[string]string {
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

func (d *functions) FromRuntime() bool {
	return d.runtime
}

func (d *functions) maybeAsync(def def.Definition) error {
	if d.async || d.ID() == def.ID() {
		return nil
	}

	fn, ok := def.(Functioner)
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

func (d *functions) IsVariadic() bool {
	return d.variadic
}

func (d *functions) Results() []def.FunctionResult {
	return d.results
}
