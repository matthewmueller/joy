package functions

import (
	"github.com/matthewmueller/golly/compiler/decl"
	"github.com/matthewmueller/golly/compiler/index"
	"github.com/matthewmueller/golly/jsast"
)

var _ decl.Declaration = (*Functions)(nil)

// New fn
func New(idx *index.Index) *Functions {
	return &Functions{}
}

// Functions struct
type Functions struct {
}

// ID string
func (d *Functions) ID() string {
	return ""
}

// Translate fn
func (d *Functions) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (d *Functions) Dependencies() (deps []decl.Declaration, err error) {
	return deps, nil
}
