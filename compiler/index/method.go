package index

import (
	"go/ast"

	"github.com/matthewmueller/golly/jsast"
)

var _ Declaration = (*Method)(nil)

// Method struct
type Method struct {
	Node *ast.FuncDecl
}

// ID string
func (d *Method) ID() string {
	return ""
}

// Translate fn
func (d *Method) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (d *Method) Dependencies() (deps []Declaration, err error) {
	return deps, nil
}
