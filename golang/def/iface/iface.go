package iface

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/def/method"
	"github.com/matthewmueller/golly/golang/index"

	"golang.org/x/tools/go/loader"
)

// Interface method
type Interface interface {
	def.Definition
	ImplementedBy(method string) []method.Method
	Node() *ast.TypeSpec
}

var _ Interface = (*ifacedef)(nil)

type ifacedef struct {
	exported  bool
	path      string
	name      string
	id        string
	index     *index.Index
	methods   []string
	node      *ast.TypeSpec
	kind      *types.Interface
	processed bool
	deps      []def.Definition
	imports   map[string]string
}

// NewInterface fn
func NewInterface(index *index.Index, info *loader.PackageInfo, n *ast.TypeSpec) (def.Definition, error) {
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

	return &ifacedef{
		exported: obj.Exported(),
		path:     packagePath,
		name:     n.Name.Name,
		id:       id,
		index:    index,
		node:     n,
		methods:  methods,
		kind:     kind,
	}, nil
}

func (d *ifacedef) process() (err error) {
	seen := map[string]bool{}
	_ = seen

	ast.Inspect(d.node, func(n ast.Node) bool {
		switch t := n.(type) {
		}

		return true
	})

	d.processed = true
	return err
}

func (d *ifacedef) ID() string {
	return d.id
}

func (d *ifacedef) Name() string {
	return d.name
}

func (d *ifacedef) Path() string {
	return d.path
}

func (d *ifacedef) Dependencies() ([]def.Definition, error) {
	return nil, nil
}

func (d *ifacedef) Exported() bool {
	return d.exported
}
func (d *ifacedef) Omitted() bool {
	return false
}

func (d *ifacedef) Node() *ast.TypeSpec {
	return d.node
}

func (d *ifacedef) Type() types.Type {
	return d.kind
}

// TODO: optimize
func (d *ifacedef) ImplementedBy(meth string) (defs []method.Method) {
	for _, n := range d.index.All() {
		m, ok := n.(method.Method)
		if !ok {
			continue
		}

		if m.Name() != meth {
			continue
		}

		if types.Implements(m.Recv().Type(), d.kind) {
			defs = append(defs, m)
		}
	}

	return defs
}

func (d *ifacedef) Imports() map[string]string {
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

func (d *ifacedef) FromRuntime() bool {
	return false
}
