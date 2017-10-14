package index

import (
	"go/ast"

	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/jsast"
)

var _ Declaration = (*function)(nil)

// function declaration type
type function struct {
	from string
	id   string
	node *ast.FuncDecl

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

func (f *function) process() {

}

// ID string
func (f *function) ID() string {
	return f.id
}

// Translate fn
func (f *function) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (f *function) Dependencies() (deps []Declaration, err error) {
	return deps, nil
}
