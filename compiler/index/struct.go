package index

import (
	"go/ast"

	"github.com/matthewmueller/golly/jsast"
)

var _ Declaration = (*Struct)(nil)

// Struct struct
type Struct struct {
	Node *ast.StructType
}

// ID string
func (d *Struct) ID() string {
	return ""
}

// Translate fn
func (d *Struct) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (d *Struct) Dependencies() (deps []Declaration, err error) {
	return deps, nil
}
