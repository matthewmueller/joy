package interfaces

import (
	"github.com/matthewmueller/golly/compiler/decl"
	"github.com/matthewmueller/golly/compiler/index"
	"github.com/matthewmueller/golly/jsast"
)

var _ decl.Declaration = (*Interfaces)(nil)

// New Interfaces
func New(idx *index.Index) *Interfaces {
	return &Interfaces{}
}

// Interfaces struct
type Interfaces struct {
}

// ID string
func (d *Interfaces) ID() string {
	return ""
}

// Translate fn
func (d *Interfaces) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (d *Interfaces) Dependencies() (deps []decl.Declaration, err error) {
	return deps, nil
}
