package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
)

var _ Enum = (*enum)(nil)

// NewEnum create a callback
func NewEnum(index index.Index, data *raw.Enum) Enum {
	return &enum{
		data:  data,
		index: index,
	}
}

// Enum interface
type Enum interface {
	def.Definition
}

// Enum struct
type enum struct {
	data *raw.Enum

	index index.Index
}

// ID fn
func (d *enum) ID() string {
	return d.data.Name
}

// Name fn
func (d *enum) Name() string {
	return d.data.Name
}

// Kind fn
func (d *enum) Kind() string {
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
func (d *enum) Dependencies() (defs []def.Definition, err error) {
	return defs, nil
}

// Generate fn
func (d *enum) Generate() (string, error) {
	// gen.Generate()
	return "", nil
}
