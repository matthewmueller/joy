package defs

import (
	"fmt"
	"strings"

	"github.com/matthewmueller/joy/internal/dom/def"
	"github.com/matthewmueller/joy/internal/dom/index"
	"github.com/matthewmueller/joy/internal/dom/raw"
	"github.com/matthewmueller/joy/internal/gen"
	"github.com/pkg/errors"
)

var _ Interface = (*iface)(nil)

// NewInterface creates an interface
func NewInterface(index index.Index, data *raw.Interface) Interface {
	d := &iface{
		data:  data,
		index: index,
	}

	methodMap := map[string]bool{}
	for _, method := range data.Methods {
		methodMap[method.Name] = true
		d.methods = append(d.methods, NewMethod(index, method, d))
	}

	for _, property := range data.Properties {
		if methodMap[property.Name] {
			property.Name = "get" + property.Name
		}
		d.properties = append(d.properties, NewProperty(index, property, d))
	}

	return d
}

// Interface interface
type Interface interface {
	ImplementedBy() ([]def.Definition, error)
	Properties() []Property
	Methods() []Method
	FindEvent(name string) (def.Definition, error)
	Ancestors() (ancestors []def.Definition, err error)
	Implements() (defs []def.Definition, err error)

	def.Definition
}

// iface struct
type iface struct {
	data *raw.Interface
	pkg  string
	file string

	index         index.Index
	implementedBy []def.Definition
	methods       []Method
	properties    []Property
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

func (d *iface) Type(caller string) (string, error) {
	imps, err := d.ImplementedBy()
	if err != nil {
		return "", errors.Wrapf(err, "implemented by")
	}

	if caller == d.pkg || d.pkg == "" {
		if len(imps) > 0 || d.data.NoInterfaceObject {
			return gen.Capitalize(d.data.Name), nil
		}
		return gen.Pointer(gen.Capitalize(d.data.Name)), nil
	}

	if len(imps) > 0 || d.data.NoInterfaceObject {
		return d.pkg + "." + gen.Capitalize(d.data.Name), nil
	}

	return gen.Pointer(d.pkg + "." + gen.Capitalize(d.data.Name)), nil
}

func (d *iface) SetPackage(pkg string) {
	d.pkg = pkg
	for _, method := range d.Methods() {
		method.SetPackage(pkg)
	}
	for _, property := range d.Properties() {
		property.SetPackage(pkg)
	}
}
func (d *iface) GetPackage() string {
	return d.pkg
}

func (d *iface) SetFile(file string) {
	d.file = file
	for _, method := range d.Methods() {
		method.SetFile(file)
	}
	for _, property := range d.Properties() {
		property.SetFile(file)
	}
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
			parents, err := t.Ancestors()
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

func (d *iface) Methods() (methods []Method) {
	return d.methods
}

func (d *iface) Properties() (props []Property) {
	return d.properties
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

	return parents, nil
}

// Dependencies fn
func (d *iface) Dependencies() (defs []def.Definition, err error) {
	// Extends
	if d.data.Extends != "" {
		if def := d.index.Find(d.data.Extends); def != nil {
			defs = append(defs, def)
		}
	}

	// Implements
	for _, imp := range d.data.Implements {
		if def := d.index.Find(imp); def != nil {
			defs = append(defs, def)
		}
	}

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

	return defs, nil
}

func (d *iface) Ancestors() (ancestors []def.Definition, err error) {
	if d.data.Extends != "" && d.data.Extends != "Object" {
		def := d.index.Find(d.data.Extends)
		if def == nil {
			return ancestors, fmt.Errorf("extends '%s' not found", d.data.Extends)
		}
		ancestors = append(ancestors, def)

		switch t := def.(type) {
		case Interface:
			a, err := t.Ancestors()
			if err != nil {
				return ancestors, errors.Wrapf(err, "error getting ancestors")
			}
			ancestors = append(ancestors, a...)
		}
	}

	return ancestors, nil
}

// Implements
func (d *iface) Implements() (defs []def.Definition, err error) {
	for _, imp := range d.data.Implements {
		def := d.index.Find(imp)
		if def == nil {
			return defs, fmt.Errorf("implements '%s' not found", imp)
		}
		defs = append(defs, def)
	}

	return defs, nil
}

// Generate fn
func (d *iface) Generate() (string, error) {
	imps, err := d.ImplementedBy()
	if err != nil {
		return "", errors.Wrap(err, "implemented by")
	}

	// ignore for now
	if len(imps) > 0 || d.data.NoInterfaceObject {
		return d.generateInterface()
	}

	data := struct {
		Package     string
		Name        string
		Type        string
		Implements  []string
		Methods     []string
		Properties  []string
		Constructor struct {
			Name    string
			Params  []gen.Vartype
			Rewrite string
		}
	}{
		Package: d.pkg,
		Name:    d.data.Name,
	}

	// Get the type
	t, err := d.Type(d.pkg)
	if err != nil {
		return "", err
	}
	data.Type = t

	methodMap := map[string]bool{}

	// add the methods of this struct
	for _, method := range d.Methods() {
		if methodMap[method.Name()] {
			continue
		}
		methodMap[method.Name()] = true

		m, err := method.GenerateAs(d)
		if err != nil {
			return "", errors.Wrapf(err, "error generating method %s", method.ID())
		}
		data.Methods = append(data.Methods, m)
	}

	// add the properties of this struct
	for _, property := range d.Properties() {
		if methodMap[property.Name()] {
			continue
		}
		methodMap[property.Name()] = true

		m, err := property.GenerateAs(d)
		if err != nil {
			return "", errors.Wrapf(err, "error generating property %s", property.ID())
		}
		data.Properties = append(data.Properties, m)
	}

	// handle the implements of this struct
	for _, imp := range d.data.Implements {
		def := d.index.Find(imp)
		if def == nil {
			return "", fmt.Errorf("implements missing %s", imp)
		}

		t, ok := def.(Interface)
		if !ok {
			return "", fmt.Errorf("implements not an interface %s", imp)
		}

		for _, method := range t.Methods() {
			if methodMap[method.Name()] {
				continue
			}
			methodMap[method.Name()] = true

			m, err := method.GenerateAs(d)
			if err != nil {
				return "", errors.Wrapf(err, "error generating method")
			}
			data.Methods = append(data.Methods, m)
		}

		for _, property := range t.Properties() {
			if methodMap[property.Name()] {
				continue
			}
			methodMap[property.Name()] = true

			m, err := property.GenerateAs(d)
			if err != nil {
				return "", errors.Wrapf(err, "error generating property")
			}
			data.Properties = append(data.Properties, m)
		}
	}

	// Handle extends and implements for all ancestors
	// TODO: cleanup this is a total mess
	ancestors, err := d.Ancestors()
	if err != nil {
		return "", errors.Wrapf(err, "generating ancestors")
	}
	for _, ancestor := range ancestors {
		switch t := ancestor.(type) {
		case Interface:
			kind, err := t.Type(d.pkg)
			if err != nil {
				return "", errors.Wrapf(err, "error getting implements type")
			}
			data.Implements = append(data.Implements, kind)

			for _, method := range t.Methods() {
				if methodMap[method.Name()] {
					continue
				}
				methodMap[method.Name()] = true

				m, err := method.GenerateAs(d)
				if err != nil {
					return "", errors.Wrapf(err, "error generating method %s", method.ID())
				}
				data.Methods = append(data.Methods, m)
			}
			for _, property := range t.Properties() {
				if methodMap[property.Name()] {
					continue
				}
				methodMap[property.Name()] = true

				m, err := property.GenerateAs(d)
				if err != nil {
					return "", errors.Wrapf(err, "error generating property %s", property.ID())
				}
				data.Properties = append(data.Properties, m)
			}

			// also implement all implementations that the ancestors implement
			imps, err := t.Implements()
			if err != nil {
				return "", errors.Wrapf(err, "error getting implements for interface %s", d.ID())
			}
			for _, def := range imps {
				switch t := def.(type) {
				case Interface:
					kind, err := t.Type(d.pkg)
					if err != nil {
						return "", errors.Wrapf(err, "error getting implements type")
					}
					data.Implements = append(data.Implements, kind)

					for _, method := range t.Methods() {
						if methodMap[method.Name()] {
							continue
						}
						methodMap[method.Name()] = true

						m, err := method.GenerateAs(d)
						if err != nil {
							return "", errors.Wrapf(err, "error generating method %s", method.ID())
						}
						data.Methods = append(data.Methods, m)
					}
					for _, property := range t.Properties() {
						if methodMap[property.Name()] {
							continue
						}
						methodMap[property.Name()] = true

						m, err := property.GenerateAs(d)
						if err != nil {
							return "", errors.Wrapf(err, "error generating property %s", property.ID())
						}
						data.Properties = append(data.Properties, m)
					}
				default:
					return "", fmt.Errorf("unhandled ancestor type %T for %s", t, d.ID())
				}
			}
		default:
			return "", fmt.Errorf("unhandled ancestor type %T for %s", t, d.ID())
		}
	}

	// handle the constructor if there is one
	if d.data.Constructor != nil {
		// New() function
		if d.pkg == d.file {
			data.Constructor.Name = "New"
		} else {
			data.Constructor.Name = "New" + gen.Capitalize(d.data.Name)
		}

		for _, param := range d.data.Constructor.Params {
			t, err := d.index.Coerce(d.pkg, param.Type)
			if err != nil {
				return "", errors.Wrapf(err, "constructor param")
			}

			data.Constructor.Params = append(data.Constructor.Params, gen.Vartype{
				Var:      gen.Lowercase(param.Name),
				Optional: param.Optional,
				Type:     t,
			})
		}

		// add in the rewrites
		switch d.file {
		case "window":
			data.Constructor.Rewrite = "window"
		default:
			ints := strings.Join(gen.Sequence(len(d.data.Constructor.Params)), ", ")
			data.Constructor.Rewrite = "new " + d.data.Name + "(" + ints + ")"
		}
	}

	return gen.Generate("struct/"+d.data.Name, data, `
		{{ range .Implements -}}
		var _ {{ . }} = (*{{ capitalize $.Name }})(nil)
		{{ end }}

		{{ if .Constructor.Name -}}
		// {{ .Constructor.Name }} fn
		func {{ .Constructor.Name }}({{ joinvt .Constructor.Params }}) {{ .Type }} {
			macro.Rewrite("{{ .Constructor.Rewrite }}", {{ joinv .Constructor.Params }})
			return &{{ capitalize .Name }}{}
		}
		{{- end }}

		// {{ capitalize .Name }} struct
		// js:"{{ capitalize .Name }},omit"
		type {{ capitalize .Name }} struct {
		}

		{{ range .Methods -}}
		{{ . }}
		{{- end }}

		{{ range .Properties -}}
		{{ . }}
		{{- end }}
	`)
}

func (d *iface) generateInterface() (string, error) {
	data := struct {
		Package    string
		Name       string
		Extends    []string
		Methods    []string
		Properties []string
		Rewrites   []string
	}{
		Package: d.pkg,
		Name:    gen.Capitalize(d.data.Name),
	}

	// Handle embeds
	parents, err := d.Parents()
	if err != nil {
		return "", errors.Wrapf(err, "error getting parents")
	}
	for _, parent := range parents {
		t, err := parent.Type(d.pkg)
		if err != nil {
			return "", errors.Wrapf(err, "parent type")
		}
		data.Extends = append(data.Extends, t)
	}

	// handle the implements too
	for _, imp := range d.data.Implements {
		def := d.index.Find(imp)
		if def == nil {
			return "", fmt.Errorf("implements missing %s", imp)
		}

		t, ok := def.(Interface)
		if !ok {
			return "", fmt.Errorf("implements not an interface %s", imp)
		}

		for _, method := range t.Methods() {
			m, err := method.GenerateInterfaceAs(d)
			if err != nil {
				return "", errors.Wrapf(err, "error generating method")
			}
			data.Methods = append(data.Methods, m)
		}

		for _, property := range t.Properties() {
			m, err := property.GenerateInterfaceAs(d)
			if err != nil {
				return "", errors.Wrapf(err, "error generating property")
			}
			data.Properties = append(data.Properties, m)
		}
	}

	// handle methods
	for _, method := range d.Methods() {
		m, err := method.GenerateInterfaceAs(d)
		if err != nil {
			return "", errors.Wrapf(err, "error generating method")
		}
		data.Methods = append(data.Methods, m)
	}

	// handle properties
	for _, property := range d.Properties() {
		m, err := property.GenerateInterfaceAs(d)
		if err != nil {
			return "", errors.Wrapf(err, "error generating property")
		}
		data.Properties = append(data.Properties, m)
	}

	return gen.Generate("interface/"+d.data.Name, data, `
		// {{ .Name }} interface
		// js:"{{ .Name }}"
		type {{ .Name }} interface {
			{{- range .Extends }}
			{{ . }}
			{{- end }}
		
			{{ range .Methods -}}
			{{ . }}
			{{- end }}
	
			{{ range .Properties -}}
			{{ . }}
			{{- end }}
		}
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
