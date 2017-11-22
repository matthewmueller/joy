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

	def.Definition
}

// iface struct
type iface struct {
	data *raw.Interface

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

// Children fn
func (d *iface) Children() (defs []def.Definition, err error) {
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

	for _, prop := range d.data.Properties {
		if prop.Type == "EventHandler" {
			def, err := d.findEvent(prop.EventHandler)
			if err != nil {
				return defs, err
			} else if def != nil {
				defs = append(defs, def)
			}
			continue
		}

		if def := d.index.Find(prop.Type); def != nil {
			defs = append(defs, def)
		}
	}

	for _, event := range d.data.Events {
		if def := d.index.Find(event.Type); def != nil {
			defs = append(defs, def)
		}
	}

	return defs, nil
}

// find the event, traversing up if necessary
func (d *iface) findEvent(name string) (def.Definition, error) {
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
			return t.findEvent(name)
		}
	}

	// return the default event
	return d.index["Event"], nil
}

type interfaceData struct {
	Name       string
	Extends    string
	Implements []string
	Methods    []string
	Properties []string
}

type methodData struct {
	Recv   string
	Name   string
	Params []gen.Vartype
	Result gen.Vartype
}

// Generate fn
func (d *iface) Generate() (string, error) {
	name := d.data.Name
	data := interfaceData{}
	data.Name = name

	// pkgname := gen.Lowercase(d.InterfaceName)

	imps, err := d.ImplementedBy()
	if err != nil {
		return "", errors.Wrap(err, "implemented by")
	}

	if len(imps) > 0 {

		log.Infof("implemented=%s", d.Name())
	}

	// for _, method := range d.Methods {
	// 	m, e := method.Generate()
	// }

	return "", nil

	// 	recv := i
	// 	if implementor != nil {
	// 		recv = implementor
	// 	}

	// 	if d.Extends != "" && d.Extends != "Object" {
	// 		data.Extends = gen.Pointer(findPackage(idx, pkgname, d.Extends))
	// 	}

	// 	for _, imp := range d.Implements {
	// 		data.Implements = append(data.Implements, gen.Pointer(findPackage(idx, pkgname, imp)))
	// 	}

	// 	for _, method := range d.Methods {
	// 		if method == nil {
	// 			continue
	// 		}

	// 		m, e := method.Generate(idx, recv)
	// 		if e != nil {
	// 			return "", e
	// 		}
	// 		data.Methods = append(data.Methods, m)
	// 	}

	// 	for _, property := range d.Properties {
	// 		if property == nil {
	// 			continue
	// 		}

	// 		m, e := property.Generate(idx, recv)
	// 		if e != nil {
	// 			return "", e
	// 		}
	// 		data.Properties = append(data.Properties, m)
	// 	}

	// 	return gen.Generate("interface/"+d.Name, data, `
	// type {{ .Name }} struct {
	// 	{{ .Extends }}

	// 	{{- range .Implements }}
	// 	{{ . }}
	// 	{{- end }}
	// }

	// {{ range .Methods -}}
	// {{ . }}
	// {{- end }}

	// {{ range .Properties -}}
	// {{ . }}
	// {{- end }}
	// `)
}
