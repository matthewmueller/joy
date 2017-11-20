package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
)

var _ def.Definition = (*Dictionary)(nil)

// Dictionary struct
type Dictionary struct {
	DictionaryName string `xml:"name,attr"`
	Extends        string `xml:"extends,attr"`
	Members        []struct {
		Name     string `xml:"name,attr"`
		Type     string `xml:"type,attr,omitempty"`
		Required bool   `xml:"required,attr,omitempty"`
		Default  string `xml:"default,attr,omitempty"`
		Nullable bool   `xml:"nullable,attr,omitempty"`
	} `xml:"members>member"`

	Index index.Index
}

// ID fn
func (d *Dictionary) ID() string {
	return d.DictionaryName
}

// Name fn
func (d *Dictionary) Name() string {
	return d.DictionaryName
}

// Kind fn
func (d *Dictionary) Kind() string {
	return "DICTIONARY"
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
func (d *Dictionary) Children() (defs []def.Definition, err error) {
	return defs, nil
}
