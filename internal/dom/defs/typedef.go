package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
)

var _ Typedef = (*typedef)(nil)

// NewTypedef fn
func NewTypedef(index index.Index, data *raw.Method) Typedef {
	return &method{
		index: index,
		data:  data,
	}
}

// Typedef interface
type Typedef interface {
	def.Definition
}

type typedef struct {
	data *raw.TypeDef

	index index.Index
}

// ID fn
func (d *typedef) ID() string {
	return d.data.NewType
}

// Name fn
func (d *typedef) Name() string {
	return d.data.NewType
}

// Kind fn
func (d *typedef) Kind() string {
	return "TYPEDEF"
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
func (d *typedef) Children() (defs []def.Definition, err error) {
	if def := d.index.Find(d.data.Type); def != nil {
		defs = append(defs, def)
	}
	return defs, nil
}

func (d *typedef) Generate() (string, error) {
	return "", nil
}
