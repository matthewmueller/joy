package defs

import (
	"fmt"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
	"github.com/matthewmueller/golly/internal/gen"
	"github.com/pkg/errors"
)

var _ Interface = (*iface)(nil)

// NewInterface creates an interface
func NewInterface(index index.Index, data *raw.Interface) Interface {
	return &iface{
		data:  data,
		index: index,
	}
}

// Interface interface
type Interface interface {
	ImplementedBy() ([]def.Definition, error)
	Properties() []def.Definition
	Methods() []def.Definition
	FindEvent(name string) (def.Definition, error)

	def.Definition
}

// iface struct
type iface struct {
	data *raw.Interface
	pkg  string
	file string

	index         index.Index
	implementedBy []def.Definition
	methods       []def.Definition
	properties    []def.Definition
}

// ID fn
func (d *iface) ID() string {
	return d.data.Name
}

// Name fn
func (d *iface) Name() string {
	return d.data.Name
}

// Kind fn
func (d *iface) Kind() string {
	return "INTERFACE"
}

func (d *iface) Type() (string, error) {
	return d.index.Coerce(d.data.Name)
}

func (d *iface) SetPackage(pkg string) {
	d.pkg = pkg
}
func (d *iface) GetPackage() string {
	return d.pkg
}

func (d *iface) SetFile(file string) {
	d.file = file
}
func (d *iface) GetFile() string {
	return d.file
}

// ImplementedBy fn
// TODO: fix, this is really inefficient
func (d *iface) ImplementedBy() (imps []def.Definition, err error) {
	if d.implementedBy != nil {
		return d.implementedBy, nil
	}

	for _, def := range d.index {
		if t, ok := def.(*iface); ok {
			parents, err := t.Parents()
			if err != nil {
				return imps, err
			}

			for _, p := range parents {
				if p.ID() == d.ID() {
					imps = append(imps, t)
				}
			}
		}
	}

	d.implementedBy = imps
	return imps, nil
}

func (d *iface) Methods() (methods []def.Definition) {
	for _, method := range d.data.Methods {
		methods = append(methods, NewMethod(d.index, method, d))
	}
	return methods
}

func (d *iface) Properties() (props []def.Definition) {
	for _, prop := range d.data.Properties {
		props = append(props, NewProperty(d.index, prop, d))
	}
	return props
}

// Parents fn
func (d *iface) Parents() (parents []def.Definition, err error) {
	if d.data.Extends != "" && d.data.Extends != "Object" {
		parent, isset := d.index[d.data.Extends]
		if !isset {
			return parents, fmt.Errorf("extends doesn't exist %s on %s", d.data.Extends, d.data.Name)
		}
		parents = append(parents, parent)
	}

	for _, imp := range d.data.Implements {
		parent, isset := d.index[imp]
		if !isset {
			return parents, fmt.Errorf("implements doesn't exist %s on %s", imp, d.data.Name)
		}
		parents = append(parents, parent)
	}

	return parents, nil
}

// Dependencies fn
func (d *iface) Dependencies() (defs []def.Definition, err error) {
	// imps, err := d.ImplementedBy()
	// if err != nil {
	// 	return defs, err
	// }
	// if len(imps) > 0 {
	// 	return defs, nil
	// }

	if d.data.Constructor != nil {
		for _, param := range d.data.Constructor.Params {
			if def := d.index.Find(param.Type); def != nil {
				defs = append(defs, def)
			}
		}
	}

	if d.data.NamedConstructor != nil {
		for _, param := range d.data.NamedConstructor.Params {
			if def := d.index.Find(param.Type); def != nil {
				defs = append(defs, def)
			}
		}
	}

	for _, method := range d.Methods() {
		deps, err := method.Dependencies()
		if err != nil {
			return defs, errors.Wrap(err, "method deps")
		}
		defs = append(defs, deps...)
	}

	for _, prop := range d.Properties() {
		deps, err := prop.Dependencies()
		if err != nil {
			return defs, errors.Wrap(err, "method deps")
		}
		defs = append(defs, deps...)
	}

	// for _, event := range d.data.Events {
	// 	if def := d.index.Find(event.Type); def != nil {
	// 		defs = append(defs, def)
	// 	}
	// }

	return defs, nil
}

// Generate fn
func (d *iface) Generate() (string, error) {
	data := struct {
		Package    string
		Name       string
		Embeds     []string
		Methods    []string
		Properties []string
	}{
		Package: d.pkg,
		Name:    gen.Capitalize(d.data.Name),
	}

	imps, err := d.ImplementedBy()
	if err != nil {
		return "", errors.Wrap(err, "implemented by")
	}
	// ignore for now
	if len(imps) > 0 {
		log.Infof("use interface=%s", d.Name())
		return gen.Generate("interface/"+d.data.Name, data, `
		package {{ .Package }}

		type {{ .Name }} interface {

		}
		`)
	}

	// Handle embeds
	parents, err := d.Parents()
	if err != nil {
		return "", errors.Wrapf(err, "error getting parents")
	}
	for _, parent := range parents {
		data.Embeds = append(data.Embeds, parent.Name())
	}

	// handle methods
	for _, method := range d.Methods() {
		m, err := method.Generate()
		if err != nil {
			return "", errors.Wrapf(err, "error generating method")
		}
		data.Methods = append(data.Methods, m)
	}

	// handle properties
	for _, property := range d.Properties() {
		m, err := property.Generate()
		if err != nil {
			return "", errors.Wrapf(err, "error generating property")
		}
		data.Properties = append(data.Properties, m)
	}

	return gen.Generate("interface/"+d.data.Name, data, `
		package {{ .Package }}
		
		type {{ .Name }} struct {
			{{- range .Embeds }}
			{{ . }}
			{{- end }}
		}

		{{ range .Methods -}}
		{{ . }}
		{{- end }}

		{{ range .Properties -}}
		{{ . }}
		{{- end }}
	`)
}

// find the event, traversing up if necessary
func (d *iface) FindEvent(name string) (def.Definition, error) {
	for _, event := range d.data.Events {
		if event.Name == name {
			if e, isset := d.index[event.Type]; isset {
				return e, nil
			}
		}
	}

	parents, err := d.Parents()
	if err != nil {
		return nil, err
	}

	for _, parent := range parents {
		if t, ok := parent.(*iface); ok {
			return t.FindEvent(name)
		}
	}

	// return the default event
	return d.index["Event"], nil
}
