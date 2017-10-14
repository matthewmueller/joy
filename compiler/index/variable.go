package index

import (
	"go/ast"

	"github.com/matthewmueller/golly/jsast"
)

var _ Declaration = (*Variable)(nil)

// Variable struct
type Variable struct {
	Node *ast.ValueSpec
}

// ID string
func (d *Variable) ID() string {
	return ""
}

// Translate fn
func (d *Variable) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (d *Variable) Dependencies() (deps []Declaration, err error) {
	return deps, nil
}
