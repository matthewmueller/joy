package defs

import (
	"github.com/matthewmueller/joy/internal/dom/def"
	"github.com/matthewmueller/joy/internal/dom/index"
	"github.com/matthewmueller/joy/internal/dom/raw"
	"github.com/matthewmueller/joy/internal/gen"
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
	pkg  string
	file string

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

func (d *cb) Type(caller string) (string, error) {
	data := struct {
		Params []gen.Vartype
		Result gen.Vartype
	}{}

	for _, param := range d.data.Params {
		t, err := d.index.Coerce(d.pkg, param.Type)
		if err != nil {
			return "", err
		}
		data.Params = append(data.Params, gen.Vartype{
			Var:      gen.Identifier(param.Name),
			Optional: param.Optional,
			Type:     t,
		})
	}

	t, err := d.index.Coerce(d.pkg, d.data.Type)
	if err != nil {
		return "", err
	}
	data.Result = gen.Vartype{
		Var:  gen.Identifier(d.data.Name),
		Type: t,
	}

	if t == "" {
		return gen.Generate("callback_fn/"+d.data.Name, data, `func ({{ joinvt .Params }})`)
	}

	return gen.Generate("callback_fn/"+d.data.Name, data, `func ({{ joinvt .Params }}) ({{ vt .Result }})`)
}

func (d *cb) SetPackage(pkg string) {
	d.pkg = pkg
}
func (d *cb) GetPackage() string {
	return d.pkg
}

func (d *cb) SetFile(file string) {
	d.file = file
}
func (d *cb) GetFile() string {
	return d.file
}

func (d *cb) Dependencies() (defs []def.Definition, err error) {
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
