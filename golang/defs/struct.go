package defs

import (
	"errors"
	"go/ast"
	"go/types"
	"strings"

	"github.com/fatih/structtag"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/index"
	"github.com/matthewmueller/golly/golang/util"

	"golang.org/x/tools/go/loader"
)

// Structer interface
type Structer interface {
	def.Definition
	Node() *ast.TypeSpec
	Fields() []*field
	Field(name string) *field
	OriginalName() string
	Methods() []def.Definition
	Implements() ([]def.Definition, error)
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
	if d.tag != nil {
		return d.tag.HasOption("omit")
	}
	return false
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
func (d *structdef) Methods() (defs []def.Definition) {
	return defs
}

// Implements function
func (d *structdef) Implements() (defs []def.Definition, err error) {
	for _, def := range d.index.All() {
		iface, ok := def.(Interfacer)
		if !ok {
			continue
		}

		structType := types.NewPointer(d.Type())
		ifaceType, ok := iface.Type().(*types.Interface)
		if !ok {
			return defs, errors.New("interface type not an interface")
		}

		if types.Implements(structType, ifaceType) {
			defs = append(defs, iface)
		}
	}

	defs = util.Unique(defs)
	return defs, nil
}
