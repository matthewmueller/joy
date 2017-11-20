package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
)

var _ def.Definition = (*Enum)(nil)

// Enum struct
type Enum struct {
	EnumName string   `xml:"name,attr"`
	Values   []string `xml:"value"`

	Index index.Index
}

// ID fn
func (d *Enum) ID() string {
	return d.EnumName
}

// Name fn
func (d *Enum) Name() string {
	return d.EnumName
}

// Kind fn
func (d *Enum) Kind() string {
	return "ENUM"
}

// // Parents fn
// func (d *Enum) Parents() []def.Definition {
// 	return nil
// }

// // Ancestors fn
// func (d *Enum) Ancestors() []def.Definition {
// 	return nil
// }

// Children fn
func (d *Enum) Children() (defs []def.Definition, err error) {
	return defs, nil
}
