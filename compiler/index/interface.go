package index

import (
	"github.com/matthewmueller/golly/jsast"
)

var _ Declaration = (*Interface)(nil)

// Interface struct
type Interface struct {
	// Node *ast.In
}

// ID string
func (d *Interface) ID() string {
	return ""
}

// Translate fn
func (d *Interface) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (d *Interface) Dependencies() (deps []Declaration, err error) {
	return deps, nil
}
