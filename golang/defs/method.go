package defs

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/index"
	"github.com/matthewmueller/golly/golang/util"
	"golang.org/x/tools/go/loader"
)

// Methoder interface
type Methoder interface {
	Functioner
	Recv() def.Definition
	OriginalName() string
}

var _ Methoder = (*methods)(nil)

type methods struct {
	index     *index.Index
	id        string
	path      string
	name      string
	kind      types.Type
	node      *ast.FuncDecl
	exported  bool
	params    []string
	results   []def.FunctionResult
	recv      string
	async     bool
	processed bool
	deps      []def.Definition
	tag       *structtag.Tag
	runtime   bool
	imports   map[string]string
	rewrite   def.Rewrite
	omit      bool
	variadic  bool
}

// Method fn
func Method(index *index.Index, info *loader.PackageInfo, n *ast.FuncDecl) (def.Definition, error) {
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	name := n.Name.Name
	idParts := []string{packagePath, name}
	var recv string

	// if it has a receiver, include that in the ID
	if n.Recv != nil && len(n.Recv.List) > 0 {
		id, e := util.GetIdentifier(n.Recv.List[0].Type)
		if e != nil {
			return nil, e
		}
		idParts = append(idParts, id.Name)
		recv = packagePath + " " + id.Name
	}

	id := strings.Join(idParts, " ")

	// loop over params and check if method is variadic
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

	// check if this declaration is from the runtime
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
	// log.Infof("before id=%s tag=%s", id, tag)
	// async := false
	// if tag != nil && tag.HasOption("async") {
	// 	async = true
	// }

	return &methods{
		id:       id,
		index:    index,
		exported: exported,
		path:     packagePath,
		name:     name,
		node:     n,
		kind:     info.TypeOf(n.Name),
		recv:     recv,
		runtime:  fromRuntime,
		params:   params,
		results:  results,
		variadic: variadic,
		imports:  map[string]string{},
		tag:      tag,
		// async:    async,
	}, nil
}

func (d *methods) process() (err error) {
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

func (d *methods) Dependencies() (deps []def.Definition, err error) {
	if d.processed {
		return d.deps, nil
	}
	err = d.process()
	if err != nil {
		return deps, err
	}
	return d.deps, nil
}

func (d *methods) IsAsync() (bool, error) {
	if d.processed {
		return d.async, nil
	}
	e := d.process()
	if e != nil {
		return false, e
	}
	return d.async, nil
}

func (d *methods) ID() string {
	return d.id
}

func (d *methods) Name() string {
	if d.tag != nil {
		return d.tag.Name
	}
	return d.name
}

func (d *methods) OriginalName() string {
	return d.name
}

func (d *methods) Path() string {
	return d.path
}

func (d *methods) Exported() bool {
	return d.exported
}

func (d *methods) Omitted() bool {
	if d.tag != nil && d.tag.HasOption("omit") {
		return true
	}

	// also omit if the receiver has been omitted
	if d.Recv().Omitted() {
		return true
	}

	return d.omit
}

func (d *methods) Node() *ast.FuncDecl {
	return d.node
}

func (d *methods) Type() types.Type {
	return d.kind
}

func (d *methods) Kind() string {
	return "METHOD"
}

func (d *methods) Recv() def.Definition {
	return d.index.Get(d.recv)
}

func (d *methods) Imports() map[string]string {
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

func (d *methods) FromRuntime() bool {
	return d.runtime
}

// Rewrite fn
func (d *methods) Rewrite() def.Rewrite {
	return d.rewrite
}

// Params fn
func (d *methods) Params() []string {
	return d.params
}

func (d *methods) maybeAsync(def def.Definition) error {
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

func (d *methods) IsVariadic() bool {
	return d.variadic
}

func (d *methods) Results() []def.FunctionResult {
	return d.results
}
