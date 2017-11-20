package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
)

var _ def.Definition = (*CallbackInterface)(nil)

// CallbackInterface struct
type CallbackInterface struct {
	CallbackInterfaceName string   `xml:"name,attr"`
	Extends               string   `xml:"extends,attr"`
	Methods               []Method `xml:"methods>method"`

	Index index.Index
}

// ID fn
func (d *CallbackInterface) ID() string {
	return d.CallbackInterfaceName
}

// Name fn
func (d *CallbackInterface) Name() string {
	return d.CallbackInterfaceName
}

// Kind fn
func (d *CallbackInterface) Kind() string {
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
func (d *CallbackInterface) Children() (defs []def.Definition, err error) {
	return defs, nil
}
