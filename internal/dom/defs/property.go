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
	GenerateInterface() (string, error)
	GenerateAs(recv Interface) (string, error)

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

func (d *prop) Type(caller string) (string, error) {
	if caller == d.pkg {
		return gen.Pointer(gen.Capitalize(d.data.Name)), nil
	}
	return gen.Pointer(d.pkg + "." + gen.Capitalize(d.data.Name)), nil
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
// NOTE: Since this is called via
func (d *prop) Dependencies() (defs []def.Definition, err error) {
	if d.data.Type == "EventHandler" {
		def, err := d.findEvent(d.data.EventHandler)
		if err != nil {
			return defs, errors.Wrapf(err, "event handler error")
		} else if def == nil {
			return defs, nil
		}
		return append(defs, def), nil
	}

	if def := d.index.Find(d.data.Type); def != nil {
		defs = append(defs, def)
	}

	return defs, nil
}

func (d *prop) Generate() (string, error) {
	return d.generate(d.recv)
}

func (d *prop) GenerateAs(recv Interface) (string, error) {
	return d.generate(recv)
}

func (d *prop) generate(recv Interface) (string, error) {
	data := struct {
		Recv    string
		Name    string
		Result  gen.Vartype
		Comment string
	}{
		Recv:    gen.Pointer(recv.Name()),
		Name:    d.data.Name,
		Comment: d.data.Comment,
	}

	if d.data.Type == "EventHandler" {
		def, err := d.findEvent(d.data.EventHandler)
		if err != nil {
			return "", err
		}

		t, err := def.Type(d.pkg)
		if err != nil {
			return "", errors.Wrapf(err, "error getting type")
		}

		// event handlers are functions
		// TODO: this seems fragile
		t = "func(" + t + ")"
		data.Result = gen.Vartype{
			Var:      gen.Lowercase(d.data.Name),
			Optional: d.data.Nullable,
			Type:     t,
		}
	} else {
		t, err := d.index.Coerce(d.pkg, d.data.Type)
		if err != nil {
			return "", errors.Wrapf(err, "error coercing")
		}

		data.Result = gen.Vartype{
			Var:  gen.Identifier(d.data.Name),
			Type: t,
		}
	}

	// only one known instance of this (property that returns a Promise that returns undefined)
	// so we'll just special case this for now
	async := strings.Contains(d.data.Type, "Promise<")
	results := []string{}

	if async && data.Result.Type == "" {
		getter, e := gen.Generate("property_getter/"+d.data.Name, data, `
			// {{ capitalize .Name }} prop {{ .Comment }}
			// js:"{{ .Name }}"
			func ({{ .Recv }}) {{ capitalize .Name }}() {
				js.Rewrite("await $_.{{ .Name }}")
			}
		`)
		if e != nil {
			return "", e
		}
		results = append(results, getter)
	} else {
		getter, e := gen.Generate("property_getter/"+d.data.Name, data, `
		// {{ capitalize .Name }} prop {{ .Comment }}
		// js:"{{ .Name }}"
		func ({{ .Recv }}) {{ capitalize .Name }}() ({{ .Result.Var }} {{ .Result.Type }}) {
			js.Rewrite("$_.{{ .Name }}")
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
		// Set{{ capitalize .Name }} prop {{ .Comment }}
		// js:"{{ .Name }}"
		func ({{ .Recv }}) Set{{ capitalize .Name }} ({{ .Result.Var }} {{ .Result.Type }}) {
			js.Rewrite("$_.{{ .Name }} = $1", {{ .Result.Var }})
		}
		`)
		if e != nil {
			return "", e
		}
		results = append(results, setter)
	}

	return strings.Join(results, "\n"), nil
}

func (d *prop) GenerateInterface() (string, error) {
	data := struct {
		Name    string
		Result  gen.Vartype
		Comment string
	}{
		Name:    d.data.Name,
		Comment: d.data.Comment,
	}

	if d.data.Type == "EventHandler" {
		def, err := d.findEvent(d.data.EventHandler)
		if err != nil {
			return "", err
		}

		t, err := def.Type(d.pkg)
		if err != nil {
			return "", errors.Wrapf(err, "error getting type")
		}

		// event handlers are functions
		// TODO: this seems fragile
		t = "func(" + t + ")"

		data.Result = gen.Vartype{
			Var:      gen.Lowercase(d.data.Name),
			Optional: d.data.Nullable,
			Type:     t,
		}
	} else {
		t, err := d.index.Coerce(d.pkg, d.data.Type)
		if err != nil {
			return "", errors.Wrapf(err, "error coercing")
		}

		data.Result = gen.Vartype{
			Var:  gen.Identifier(d.data.Name),
			Type: t,
		}
	}

	// only one known instance of this (property that returns a Promise that returns undefined)
	// so we'll just special case this for now
	async := strings.Contains(d.data.Type, "Promise<")
	results := []string{}

	if async && data.Result.Type == "" {
		getter, e := gen.Generate("property_getter/"+d.data.Name, data, `
			// {{ capitalize .Name }} prop
			// js:"{{ .Name }}"
			{{ capitalize .Name }}()
		`)
		if e != nil {
			return "", e
		}
		results = append(results, getter)
	} else {
		getter, e := gen.Generate("property_getter/"+d.data.Name, data, `
		// {{ capitalize .Name }} prop {{ .Comment }}
		// js:"{{ .Name }}"
		{{ capitalize .Name }}() ({{ .Result.Var }} {{ .Result.Type }})
		`)
		if e != nil {
			return "", e
		}
		results = append(results, getter)
	}

	if !d.data.ReadOnly {
		setter, e := gen.Generate("property_setter/"+d.data.Name, data, `
		// Set{{ capitalize .Name }} prop {{ .Comment }}
		// js:"{{ .Name }}"
		Set{{ capitalize .Name }} ({{ .Result.Var }} {{ .Result.Type }})
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
