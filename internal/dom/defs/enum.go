package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
	"github.com/matthewmueller/golly/internal/gen"
)

var _ Enum = (*enum)(nil)

// NewEnum create a callback
func NewEnum(index index.Index, data *raw.Enum) Enum {
	return &enum{
		data:  data,
		index: index,
	}
}

// Enum interface
type Enum interface {
	def.Definition
}

// Enum struct
type enum struct {
	data *raw.Enum
	pkg  string
	file string

	index index.Index
}

// ID fn
func (d *enum) ID() string {
	return d.data.Name
}

// Name fn
func (d *enum) Name() string {
	return d.data.Name
}

// Kind fn
func (d *enum) Kind() string {
	return "ENUM"
}

func (d *enum) Type(caller string) (string, error) {
	if caller == d.pkg {
		return gen.Pointer(gen.Capitalize(d.data.Name)), nil
	}
	return gen.Pointer(d.pkg + "." + gen.Capitalize(d.data.Name)), nil
}

func (d *enum) SetPackage(pkg string) {
	d.pkg = pkg
}
func (d *enum) GetPackage() string {
	return d.pkg
}

func (d *enum) SetFile(file string) {
	d.file = file
}
func (d *enum) GetFile() string {
	return d.file
}

// Children fn
func (d *enum) Dependencies() (defs []def.Definition, err error) {
	return defs, nil
}

// Generate fn
func (d *enum) Generate() (string, error) {
	data := struct {
		Package string
		Name    string
	}{
		Package: d.pkg,
		Name:    gen.Capitalize(d.data.Name),
	}

	return gen.Generate("enum/"+d.data.Name, data, `
		package {{ .Package }}

		type {{ .Name }} string
	`)
}
