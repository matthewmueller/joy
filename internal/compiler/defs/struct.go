package defs

import (
	"errors"
	"go/ast"
	"go/types"
	"strings"

	"github.com/fatih/structtag"

	"github.com/matthewmueller/golly/internal/compiler/def"
	"github.com/matthewmueller/golly/internal/compiler/index"

	"golang.org/x/tools/go/loader"
)

// Structer interface
type Structer interface {
	def.Definition
	Gen() *ast.GenDecl
	Node() *ast.TypeSpec
	Fields() []*field
	Field(name string) *field
	Methods() []def.Definition
	Implements() ([]Interfacer, error)
}

var _ Structer = (*structdef)(nil)

type structdef struct {
	exported  bool
	path      string
	name      string
	id        string
	index     *index.Index
	gen       *ast.GenDecl
	node      *ast.TypeSpec
	kind      types.Type
	tag       *structtag.Tag
	fields    []*field
	deps      []def.Definition
	processed bool
	omit      bool
}

// Struct fn
func Struct(index *index.Index, info *loader.PackageInfo, gn *ast.GenDecl, n *ast.TypeSpec) (def.Definition, error) {
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	name := n.Name.Name
	idParts := []string{packagePath, name}
	id := strings.Join(idParts, " ")

	// log.Infof("info.TypeOf(n.Name) %s", info.TypeOf(n.Name))

	return &structdef{
		exported: obj.Exported(),
		path:     packagePath,
		name:     name,
		id:       id,
		index:    index,
		node:     n,
		gen:      gn,
		kind:     info.TypeOf(n.Name),
	}, nil
}

func (d *structdef) process() (err error) {
	state, e := process(d.index, d, d.gen)
	if e != nil {
		return e
	}

	// copy state into function
	d.processed = true
	d.deps = state.deps
	d.omit = state.omit
	d.tag = state.tag
	d.fields = state.fields

	return nil
}

func (d *structdef) Dependencies() (deps []def.Definition, err error) {
	if d.processed {
		return d.deps, nil
	}
	err = d.process()
	if err != nil {
		return deps, err
	}
	return d.deps, nil
}

func (d *structdef) ID() string {
	return d.id
}

func (d *structdef) Name() string {
	if d.tag != nil {
		return d.tag.Name
	}
	return d.name
}

func (d *structdef) Kind() string {
	return "STRUCT"
}

func (d *structdef) OriginalName() string {
	return d.name
}

func (d *structdef) Path() string {
	return d.path
}

func (d *structdef) Exported() bool {
	return d.exported
}

func (d *structdef) Omitted() bool {
	if d.tag != nil && d.tag.HasOption("omit") {
		return true
	}
	return d.omit
}

func (d *structdef) Gen() *ast.GenDecl {
	return d.gen
}

func (d *structdef) Node() *ast.TypeSpec {
	return d.node
}

func (d *structdef) Type() types.Type {
	return d.kind
}

func (d *structdef) Fields() (fields []*field) {
	for _, f := range d.fields {
		fields = append(fields, f)
	}
	return fields
}

func (d *structdef) Field(name string) *field {
	for _, f := range d.fields {
		if f.name == name {
			return f
		}
	}
	return nil
}

func (d *structdef) Imports() map[string]string {
	return d.index.GetImports(d.path)
}

func (d *structdef) FromRuntime() bool {
	return false
}

// Methods functions
func (d *structdef) Methods() (methods []def.Definition) {
	for _, def := range d.index.All() {
		method, ok := def.(Methoder)
		if !ok {
			continue
		}
		if d.ID() == method.Recv().ID() {
			methods = append(methods, method)
		}
	}
	return methods
}

// Implements function
func (d *structdef) Implements() (ifaces []Interfacer, err error) {
	for _, def := range d.index.All() {
		iface, ok := def.(Interfacer)
		if !ok {
			continue
		}

		structType := types.NewPointer(d.Type())
		ifaceType, ok := iface.Type().(*types.Interface)
		if !ok {
			return ifaces, errors.New("interface type not an interface")
		}

		if types.Implements(structType, ifaceType) {
			ifaces = append(ifaces, iface)
		}
	}

	ifaces = uniqueInterfaces(ifaces)
	return ifaces, nil
}

func uniqueInterfaces(defs []Interfacer) []Interfacer {
	u := make([]Interfacer, 0, len(defs))
	seen := make(map[string]bool)

	for _, def := range defs {
		if _, ok := seen[def.ID()]; !ok {
			seen[def.ID()] = true
			u = append(u, def)
		}
	}

	return u
}
