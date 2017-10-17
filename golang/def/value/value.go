package value

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/index"

	"golang.org/x/tools/go/loader"
)

// Value interface
type Value interface {
	def.Definition
}

var _ Value = (*valuedef)(nil)

type valuedef struct {
	exported bool
	path     string
	name     string
	id       string
	index    *index.Index
	node     *ast.ValueSpec
	kind     types.Type
}

// NewValue fn
func NewValue(index *index.Index, info *loader.PackageInfo, n *ast.ValueSpec) (def.Definition, error) {
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
		// kind:     VALUE,
	}, nil
}

func (d *valuedef) ID() string {
	return d.id
}

func (d *valuedef) Name() string {
	return d.name
}

func (d *valuedef) Path() string {
	return d.path
}

func (d *valuedef) Dependencies() ([]def.Definition, error) {
	return nil, nil
}

func (d *valuedef) Exported() bool {
	return d.exported
}

func (d *valuedef) Omitted() bool {
	return false
}

func (d *valuedef) Node() ast.Node {
	return d.node
}

func (d *valuedef) Type() types.Type {
	return d.kind
}
