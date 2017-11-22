package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
)

var _ TypeDef = (*typedef)(nil)

// NewTypeDef fn
func NewTypeDef(index index.Index, data *raw.TypeDef) TypeDef {
	return &typedef{
		index: index,
		data:  data,
	}
}

// TypeDef interface
type TypeDef interface {
	def.Definition
}

type typedef struct {
	data *raw.TypeDef
	pkg  string
	file string

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

func (d *typedef) Type() (string, error) {
	return d.index.Coerce(d.data.NewType)
}

func (d *typedef) SetPackage(pkg string) {
	d.pkg = pkg
}
func (d *typedef) GetPackage() string {
	return d.pkg
}

func (d *typedef) SetFile(file string) {
	d.file = file
}
func (d *typedef) GetFile() string {
	return d.file
}

// // Parents fn
// func (d *Dictionary) Parents() []def.Definition {
// 	return nil
// }

// // Ancestors fn
// func (d *Dictionary) Ancestors() []def.Definition {
// 	return nil
// }

// Dependencies fn
func (d *typedef) Dependencies() (defs []def.Definition, err error) {
	if def := d.index.Find(d.data.Type); def != nil {
		defs = append(defs, def)
	}
	return defs, nil
}

func (d *typedef) Generate() (string, error) {
	return "", nil
}
