package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
)

var _ Callback = (*cb)(nil)

// NewCallback create a callback
func NewCallback(index index.Index, data *raw.Callback) Callback {
	return &cb{
		data:  data,
		index: index,
	}
}

// Callback interface
type Callback interface {
	def.Definition
}

// cb struct
type cb struct {
	data *raw.Callback

	index index.Index
}

// ID cb
func (d *cb) ID() string {
	return d.data.Name
}

// Name cb
func (d *cb) Name() string {
	return d.data.Name
}

// Kind cb
func (d *cb) Kind() string {
	return "CALLBACK"
}

// // Parents cb
// func (d *Dictionary) Parents() []def.Definition {
// 	return nil
// }

// // Ancestors cb
// func (d *Dictionary) Ancestors() []def.Definition {
// 	return nil
// }

// Children cb
func (d *cb) Children() (defs []def.Definition, err error) {
	for _, param := range d.data.Params {
		if def := d.index.Find(param.Type); def != nil {
			defs = append(defs, def)
		}
	}
	return defs, nil
}

func (d *cb) Generate() (string, error) {
	return "", nil
}
