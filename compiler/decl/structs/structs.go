package structs

import (
	"github.com/matthewmueller/golly/compiler/decl"
	"github.com/matthewmueller/golly/compiler/index"
	"github.com/matthewmueller/golly/jsast"
)

var _ decl.Declaration = (*Structs)(nil)

// New Structs
func New(idx *index.Index) *Structs {
	return &Structs{}
}

// Structs struct
type Structs struct {
}

// ID string
func (d *Structs) ID() string {
	return ""
}

// Translate fn
func (d *Structs) Translate() (jsast.IStatement, error) {
	return nil, nil
}

// Dependencies fn
func (d *Structs) Dependencies() (deps []decl.Declaration, err error) {
	return deps, nil
}
