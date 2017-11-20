package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
)

var _ def.Definition = (*TypeDef)(nil)

// TypeDef struct
type TypeDef struct {
	NewType string `xml:"new-type,attr"`
	Type    string `xml:"type,attr"`

	Index index.Index
}

// ID fn
func (d *TypeDef) ID() string {
	return d.NewType
}

// Name fn
func (d *TypeDef) Name() string {
	return d.NewType
}

// Kind fn
func (d *TypeDef) Kind() string {
	return "TYPEDEF"
}

// // Parents fn
// func (d *Dictionary) Parents() []def.Definition {
// 	return nil
// }

// // Ancestors fn
// func (d *Dictionary) Ancestors() []def.Definition {
// 	return nil
// }

// Children fn
func (d *TypeDef) Children() (defs []def.Definition, err error) {
	return defs, nil
}
