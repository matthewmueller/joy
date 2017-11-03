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
	recv      string
	async     bool
	processed bool
	edges     []def.Edge
	tag       *structtag.Tag
	runtime   bool
	imports   map[string]string
	rewrite   *rewrite
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
		imports:  map[string]string{},
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
	d.edges = state.edges.Edges()
	d.imports = state.imports
	d.omit = state.omit
	d.rewrite = state.rewrite
	d.tag = state.tag
	d.variadic = state.variadic

	return nil
}

func (d *methods) Dependencies() (edges []def.Edge, err error) {
	if d.processed {
		return d.edges, nil
	}
	err = d.process()
	if err != nil {
		return edges, err
	}
	return d.edges, nil
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
func (d *methods) Rewrite(arguments []string) (string, error) {
	if !d.processed {
		if e := d.process(); e != nil {
			return "", e
		}
	}

	if d.rewrite == nil {
		return "", nil
	}

	return d.rewrite.Rewrite(arguments)
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

func (d *methods) IsVariadic() (bool, error) {
	if d.processed {
		return d.variadic, nil
	}
	e := d.process()
	if e != nil {
		return false, e
	}
	return d.variadic, nil
}
