package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
	"github.com/pkg/errors"
)

var _ Property = (*prop)(nil)

// NewProperty creates a new property
func NewProperty(index index.Index, data *raw.Property, receiver Interface) Property {
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
	recv    Interface
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

// Dependencies fn
func (d *prop) Dependencies() (defs []def.Definition, err error) {
	if d.data.Type == "EventHandler" {
		def, err := d.recv.FindEvent(d.data.EventHandler)
		if err != nil {
			return defs, err
		} else if def != nil {
			return append(defs, def), nil
		}

		// generic event
		def, err = d.recv.FindEvent("Event")
		if err != nil {
			return defs, err
		} else if def == nil {
			return defs, errors.New("property/dependencies: Event shouldn't be nil")
		}

		return append(defs, def), nil
	}

	if def := d.index.Find(d.data.Type); def != nil {
		defs = append(defs, def)
	}

	return defs, nil
}

func (d *prop) Generate() (string, error) {
	return "", nil
}
