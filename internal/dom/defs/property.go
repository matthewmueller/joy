package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
)

var _ Property = (*prop)(nil)

// NewProperty creates a new property
func NewProperty(index index.Index, data *raw.Property, receiver def.Definition) Property {
	return &prop{
		index: index,
		data:  data,
		recv:  receiver,
	}
}

// Property interface
type Property interface {
	def.Definition
}

// prop struct
type prop struct {
	data *raw.Property

	index   index.Index
	comment string
	recv    def.Definition
}

func (d *prop) ID() string {
	return d.recv.Name() + " " + d.data.Name
}

func (d *prop) Name() string {
	return d.data.Name
}

func (d *prop) Kind() string {
	return "PROPERTY"
}

// Children fn
func (d *prop) Children() (defs []def.Definition, err error) {
	// if def := d.index.Find(d.data.Type); def != nil {
	// 	defs = append(defs, def)
	// }
	return defs, nil
}

func (d *prop) Generate() (string, error) {
	return "", nil
}
