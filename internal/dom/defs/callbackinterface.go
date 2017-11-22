package defs

import (
	"fmt"

	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
	"github.com/matthewmueller/golly/internal/gen"
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
	pkg  string
	file string

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

func (d *cbiface) Type() (string, error) {
	return d.index.Coerce(d.data.Name)
}

func (d *cbiface) SetPackage(pkg string) {
	d.pkg = pkg
}
func (d *cbiface) GetPackage() string {
	return d.pkg
}

func (d *cbiface) SetFile(file string) {
	d.file = file
}
func (d *cbiface) GetFile() string {
	return d.file
}

// Children fn
func (d *cbiface) Dependencies() (defs []def.Definition, err error) {
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
	data := struct {
		Name   string
		Params []gen.Vartype
		Result gen.Vartype
	}{
		Name: gen.Capitalize(d.data.Name),
	}

	if len(d.data.Methods) != 1 {
		return "", fmt.Errorf("callback_interface: expected %s to only have 1 method, but it has %d methods", d.data.Name, len(d.data.Methods))
	}

	method := d.data.Methods[0]
	for _, param := range method.Params {
		t, e := d.index.Coerce(param.Type)
		if e != nil {
			return "", e
		}
		data.Params = append(data.Params, gen.Vartype{
			Var:      gen.Identifier(param.Name),
			Optional: param.Optional,
			Type:     t,
		})
	}

	t, e := d.index.Coerce(method.Type)
	if e != nil {
		return "", e
	}
	data.Result = gen.Vartype{
		Var:  gen.Identifier(method.Name),
		Type: t,
	}

	if t == "void" {
		return gen.Generate("callback_interface/"+d.data.Name, data, `func ({{ joinvt .Params }})`)
	}

	return gen.Generate("callback_interface/"+d.data.Name, data, `func ({{ joinvt .Params }}) ({{ vt .Result }})`)
}
