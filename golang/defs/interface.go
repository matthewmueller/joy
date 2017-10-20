package defs

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/index"

	"golang.org/x/tools/go/loader"
)

// Interfacer method
type Interfacer interface {
	def.Definition
	ImplementedBy(method string) []Methoder
	DependenciesOf(method string) ([]def.Definition, error)
	Node() *ast.TypeSpec
}

var _ Interfacer = (*interfaces)(nil)

type interfaces struct {
	exported   bool
	path       string
	name       string
	id         string
	index      *index.Index
	methods    []string
	node       *ast.TypeSpec
	kind       *types.Interface
	processed  bool
	methodDeps map[string]def.Definition
	imports    map[string]string
}

// Interface fn
func Interface(index *index.Index, info *loader.PackageInfo, gn *ast.GenDecl, n *ast.TypeSpec) (def.Definition, error) {
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	idParts := []string{packagePath, n.Name.Name}
	id := strings.Join(idParts, " ")
	methods := []string{}

	iface := n.Type.(*ast.InterfaceType)
	for _, list := range iface.Methods.List {
		for _, method := range list.Names {
			methods = append(methods, method.Name)
		}
	}

	kind, ok := info.TypeOf(n.Type).(*types.Interface)
	if !ok {
		return nil, fmt.Errorf("NewInterface: expected n.Type to be an interface")
	}

	return &interfaces{
		exported: obj.Exported(),
		path:     packagePath,
		name:     n.Name.Name,
		id:       id,
		index:    index,
		node:     n,
		methods:  methods,
		kind:     kind,
		imports:  map[string]string{},
	}, nil
}

func (d *interfaces) process() (err error) {
	state, e := process(d.index, d, d.node)
	if e != nil {
		return e
	}

	// copy state into function
	d.processed = true
	d.imports = state.imports

	return nil
}

// interfaces don't include dependencies on their own
func (d *interfaces) Dependencies() (defs []def.Definition, err error) {
	if !d.processed {
		d.process()
	}
	return defs, nil
}

func (d *interfaces) ID() string {
	return d.id
}

func (d *interfaces) Name() string {
	return d.name
}

func (d *interfaces) Path() string {
	return d.path
}

func (d *interfaces) Exported() bool {
	return false
}
func (d *interfaces) Omitted() bool {
	return true
}

func (d *interfaces) Node() *ast.TypeSpec {
	return d.node
}

func (d *interfaces) Kind() string {
	return "INTERFACE"
}

func (d *interfaces) Type() types.Type {
	return d.kind
}

// TODO: optimize
func (d *interfaces) ImplementedBy(m string) (defs []Methoder) {
	for _, def := range d.index.All() {
		method, ok := def.(Methoder)
		if !ok {
			continue
		}

		if method.Name() != m {
			continue
		}

		// TODO: pointer is required for the deep interfaces test and
		// it works for the other tests, but will it always work?
		if types.Implements(types.NewPointer(method.Recv().Type()), d.kind) {
			defs = append(defs, method)
		}
	}

	return defs
}

func (d *interfaces) DependenciesOf(m string) (defs []def.Definition, err error) {
	return defs, nil
}

func (d *interfaces) Imports() map[string]string {
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

func (d *interfaces) FromRuntime() bool {
	return false
}
