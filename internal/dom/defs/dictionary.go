package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
)

var _ Dictionary = (*dict)(nil)

// NewDictionary create a callback
func NewDictionary(index index.Index, data *raw.Dictionary) Dictionary {
	return &dict{
		data:  data,
		index: index,
	}
}

// Dictionary interface
type Dictionary interface {
	def.Definition
}

// dict struct
type dict struct {
	data *raw.Dictionary

	index index.Index
}

// ID fn
func (d *dict) ID() string {
	return d.data.Name
}

// Name fn
func (d *dict) Name() string {
	return d.data.Name
}

// Kind fn
func (d *dict) Kind() string {
	return "DICTIONARY"
}

// // Parents fn
// func (d *dict) Parents() []def.Definition {
// 	return nil
// }

// // Ancestors fn
// func (d *dict) Ancestors() []def.Definition {
// 	return nil
// }

// Children fn
func (d *dict) Dependencies() (defs []def.Definition, err error) {
	for _, member := range d.data.Members {
		if def := d.index.Find(member.Type); def != nil {
			defs = append(defs, def)
		}
	}
	return defs, nil
}

// Generate fn
func (d *dict) Generate() (string, error) {
	return "", nil
}
