package defs

import (
	"strings"

	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
	"github.com/matthewmueller/golly/internal/gen"
	"github.com/pkg/errors"
)

var _ Property = (*prop)(nil)

// NewProperty creates a new property
func NewProperty(index index.Index, data *raw.Property, receiver Interface) Property {
	return &prop{
		index: index,
		data:  data,
		recv:  receiver,
	}
}

// Property interface
type Property interface {
	def.Definition
}

// prop struct
type prop struct {
	data *raw.Property
	pkg  string
	file string

	index   index.Index
	comment string
	recv    Interface
}

func (d *prop) ID() string {
	return d.recv.ID() + " " + d.data.Name
}

func (d *prop) Name() string {
	return d.data.Name
}

func (d *prop) Kind() string {
	return "PROPERTY"
}

func (d *prop) Type() (string, error) {
	return d.index.Coerce(d.data.Name)
}

func (d *prop) SetPackage(pkg string) {
	d.pkg = pkg
}
func (d *prop) GetPackage() string {
	return d.pkg
}

func (d *prop) SetFile(file string) {
	d.file = file
}
func (d *prop) GetFile() string {
	return d.file
}


// Dependencies fn
func (d *prop) Dependencies() (defs []def.Definition, err error) {
	if d.data.Type == "EventHandler" {
		def, err := d.findEvent(d.data.EventHandler)
		if err != nil {
			return append(defs, def), nil
		}
	}

	if def := d.index.Find(d.data.Type); def != nil {
		defs = append(defs, def)
	}

	return defs, nil
}

func (d *prop) Generate() (string, error) {
	data := struct {
		Recv   string
		Name   string
		Result gen.Vartype
	}{
		Recv: gen.Pointer(d.recv.Name()),
		Name: gen.Capitalize(d.data.Name),
	}

	if d.data.Type == "EventHandler" {
		def, err := d.findEvent(d.data.EventHandler)
		if err != nil {
			return "", err
		}

		t, err := def.Type()
		if err != nil {
			return "", errors.Wrapf(err, "error getting type")
		}

		data.Result = gen.Vartype{
			Var:      gen.Identifier(def.Name()),
			Optional: d.data.Nullable,
			Type:     t,
		}
	}

	t, err := d.index.Coerce(d.data.Type)
	if err != nil {
		return "", errors.Wrapf(err, "error coercing")
	}

	data.Result = gen.Vartype{
		Var:  gen.Identifier(d.data.Name),
		Type: t,
	}

	// only one known instance of this (property that returns a Promise that returns undefined)
	// so we'll just special case this for now
	async := strings.Contains(d.data.Type, "Promise<")
	results := []string{}

	if async && data.Result.Type == "" {
		getter, e := gen.Generate("property_getter/"+d.data.Name, data, `
			func ({{ .Recv }}) Get{{ capitalize .Name }}() {
				js.Rewrite("await $<.{{ .Name }}")
			}
		`)
		if e != nil {
			return "", e
		}
		results = append(results, getter)
	} else {
		getter, e := gen.Generate("property_getter/"+d.data.Name, data, `
		func ({{ .Recv }}) Get{{ capitalize .Name }}() ({{ .Result.Var }} {{ .Result.Type }}) {
			js.Rewrite("$<.{{ .Name }}")
			return {{ .Result.Var }}
		}
		`)
		if e != nil {
			return "", e
		}
		results = append(results, getter)
	}

	if !d.data.ReadOnly {
		setter, e := gen.Generate("property_setter/"+d.data.Name, data, `
		func ({{ .Recv }}) Set{{ capitalize .Name }} ({{ .Result.Var }} {{ .Result.Type }}) {
			js.Rewrite("$<.{{ .Name }} = $1", {{ .Result.Var }})
		}
		`)
		if e != nil {
			return "", e
		}
		results = append(results, setter)
	}

	return strings.Join(results, "\n"), nil
}

func (d *prop) findEvent(name string) (def.Definition, error) {
	// find event
	def, err := d.recv.FindEvent(d.data.EventHandler)
	if err != nil {
		return nil, err
	} else if def != nil {
		return def, nil
	}

	// generic event
	def, err = d.recv.FindEvent("Event")
	if err != nil {
		return nil, err
	} else if def == nil {
		return nil, errors.New("property/findEvent: Event shouldn't be nil")
	}

	return def, nil
}
