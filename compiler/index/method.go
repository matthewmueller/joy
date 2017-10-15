package index

import (
	"go/ast"

	"golang.org/x/tools/go/loader"

	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/compiler/types"
	"github.com/matthewmueller/golly/jsast"
)

var _ types.Declaration = (*method)(nil)

// Method fn
func (i *Index) Method(info *loader.PackageInfo, n *ast.FuncDecl) error {
	id := info.ObjectOf(n.Name).String()

	fn := &method{
		info:  info,
		index: i,
		id:    id,
		node:  n,
	}

	// add to the index
	i.decls[id] = fn

	return nil
}

// method struct
type method struct {
	info  *loader.PackageInfo
	index *Index
	id    string
	node  *ast.FuncDecl

	name     string
	omit     bool
	exported bool
	async    bool

	// Just names for now because
	// we only need the names right
	// now for js.Rewrite.
	params []string

	// This is nonnil when there is a comment above
	// the declaration
	// e.g. // js:"..."
	tag *structtag.Tag

	// Rewrite *Rewrite
}

// ID string
func (d *method) ID() string {
	return ""
}

// Name of the declaration
func (d *method) Name() string {
	return ""
}

// Exported fn
// TODO
func (d *method) Exported() bool {
	return true
}

// Path string
func (d *method) Path() string {
	return d.info.Pkg.Path()
}

// Translate fn
func (d *method) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (d *method) Dependencies() (deps []types.Declaration, err error) {
	return deps, nil
}
