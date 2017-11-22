package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
)

var _ CallbackInterface = (*cbiface)(nil)

// NewCallbackInterface create a callback
func NewCallbackInterface(index index.Index, data *raw.CallbackInterface) CallbackInterface {
	return &cbiface{
		data:  data,
		index: index,
	}
}

// CallbackInterface interface
type CallbackInterface interface {
	def.Definition
}

// cbiface struct
type cbiface struct {
	data *raw.CallbackInterface

	index index.Index
}

// ID fn
func (d *cbiface) ID() string {
	return d.data.Name
}

// Name fn
func (d *cbiface) Name() string {
	return d.data.Name
}

// Kind fn
func (d *cbiface) Kind() string {
	return "CALLBACK_INTERFACE"
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
func (d *cbiface) Children() (defs []def.Definition, err error) {
	for _, method := range d.data.Methods {
		for _, param := range method.Params {
			if def := d.index.Find(param.Type); def != nil {
				defs = append(defs, def)
			}
		}
		if def := d.index.Find(method.Type); def != nil {
			defs = append(defs, def)
		}
	}
	return defs, nil
}

// Generate fn
func (d *cbiface) Generate() (string, error) {
	return "", nil
}
