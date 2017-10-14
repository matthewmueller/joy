package variables

import (
	"github.com/matthewmueller/golly/compiler/decl"
	"github.com/matthewmueller/golly/compiler/index"
	"github.com/matthewmueller/golly/jsast"
)

var _ decl.Declaration = (*Variables)(nil)

// New Variables
func New(idx *index.Index) *Variables {
	return &Variables{}
}

// Variables struct
type Variables struct {
}

// ID string
func (d *Variables) ID() string {
	return ""
}

// Translate fn
func (d *Variables) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (d *Variables) Dependencies() (deps []decl.Declaration, err error) {
	return deps, nil
}
