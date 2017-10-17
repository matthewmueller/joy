package struc

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/index"

	"golang.org/x/tools/go/loader"
)

// Struct interface
type Struct interface {
	def.Definition
	Node() *ast.TypeSpec
	Fields() []string
}

var _ Struct = (*structdef)(nil)

type structdef struct {
	exported bool
	path     string
	name     string
	id       string
	index    *index.Index
	node     *ast.TypeSpec
	kind     types.Type
}

// NewStruct fn
func NewStruct(index *index.Index, info *loader.PackageInfo, n *ast.TypeSpec) (def.Definition, error) {
	// typeof := info.TypeOf(n.Name)
	obj := info.ObjectOf(n.Name)
	packagePath := obj.Pkg().Path()
	idParts := []string{packagePath, n.Name.Name}
	id := strings.Join(idParts, " ")

	return &structdef{
		exported: obj.Exported(),
		path:     packagePath,
		name:     n.Name.Name,
		id:       id,
		index:    index,
		node:     n,
		kind:     info.TypeOf(n.Name),
	}, nil
}

func (d *structdef) ID() string {
	return d.id
}

func (d *structdef) Name() string {
	return d.name
}

func (d *structdef) Path() string {
	return d.path
}

func (d *structdef) Dependencies() ([]def.Definition, error) {
	return nil, nil
}

func (d *structdef) Exported() bool {
	return d.exported
}

func (d *structdef) Omitted() bool {
	return false
}

func (d *structdef) Node() *ast.TypeSpec {
	return d.node
}

func (d *structdef) Type() types.Type {
	return d.kind
}

func (d *structdef) Fields() (fields []string) {
	st := d.node.Type.(*ast.StructType)
	for _, field := range st.Fields.List {
		for _, ident := range field.Names {
			fields = append(fields, ident.Name)
		}
	}
	return fields
}
