package methods

import (
	"github.com/matthewmueller/golly/compiler/decl"
	"github.com/matthewmueller/golly/compiler/index"
	"github.com/matthewmueller/golly/jsast"
)

var _ decl.Declaration = (*Methods)(nil)

// New Methods
func New(idx *index.Index) *Methods {
	return &Methods{}
}

// Methods struct
type Methods struct {
}

// ID string
func (d *Methods) ID() string {
	return ""
}

// Translate fn
func (d *Methods) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (d *Methods) Dependencies() (deps []decl.Declaration, err error) {
	return deps, nil
}
