package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
)

var _ def.Definition = (*Callback)(nil)

// Callback struct
type Callback struct {
	CallbackName string  `xml:"name,attr"`
	Callback     bool    `xml:"callback,attr,omitempty"`
	Type         string  `xml:"type,attr"`
	Params       []Param `xml:"param"`

	Index index.Index
}

// ID fn
func (d *Callback) ID() string {
	return d.CallbackName
}

// Name fn
func (d *Callback) Name() string {
	return d.CallbackName
}

// Kind fn
func (d *Callback) Kind() string {
	return "CALLBACK"
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
func (d *Callback) Children() (defs []def.Definition, err error) {
	return defs, nil
}
