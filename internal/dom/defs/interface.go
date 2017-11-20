package defs

import (
	"fmt"

	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
)

var _ def.Definition = (*Interface)(nil)

// Interface struct
type Interface struct {
	InterfaceName              string                        `xml:"name,attr"`
	Extends                    string                        `xml:"extends,attr"`
	Implements                 []string                      `xml:"implements,omitempty"`
	NoInterfaceObject          bool                          `xml:"no-interface-object,attr"`
	Element                    []*Element                    `xml:"element"`
	Methods                    []*Method                     `xml:"methods>method"`
	AnonymousMethods           []*Method                     `xml:"anonymous-methods>method"`
	Properties                 []*Property                   `xml:"properties>property"`
	Constants                  []*Constant                   `xml:"constants>constant"`
	Constructor                *Constructor                  `xml:"constructor,omitempty"`
	NamedConstructor           *NamedConstructor             `xml:"named-constructor"`
	Events                     []*Event                      `xml:"events>event"`
	AnonymousContentAttributes []*AnonymousContentAttributes `xml:"anonymous-content-attributes>parsedattribute"`

	Index index.Index
}

// ID fnw
func (d *Interface) ID() string {
	return d.InterfaceName
}

// Name fn
func (d *Interface) Name() string {
	return d.InterfaceName
}

// Kind fn
func (d *Interface) Kind() string {
	return "INTERFACE"
}

// Parents fn
func (d *Interface) Parents() (parents []def.Definition, err error) {
	if d.Extends != "" && d.Extends != "Object" {
		parent, isset := d.Index[d.Extends]
		if !isset {
			return parents, fmt.Errorf("extends doesn't exist %s on %s", d.Extends, d.InterfaceName)
		}
		parents = append(parents, parent)
	}

	for _, imp := range d.Implements {
		parent, isset := d.Index[imp]
		if !isset {
			return parents, fmt.Errorf("implements doesn't exist %s on %s", imp, d.InterfaceName)
		}
		parents = append(parents, parent)
	}

	return parents, nil
}

// // Ancestors fn
// func (d *Interface) Ancestors() []def.Definition {
// 	return nil
// }

// Children fn
func (d *Interface) Children() (defs []def.Definition, err error) {
	if d.Constructor != nil {
		for _, param := range d.Constructor.Params {
			if def := d.Index.Find(param.Type); def != nil {
				defs = append(defs, def)
			}
		}
	}

	if d.NamedConstructor != nil {
		for _, param := range d.NamedConstructor.Params {
			if def := d.Index.Find(param.Type); def != nil {
				defs = append(defs, def)
			}
		}
	}

	for _, method := range d.Methods {
		for _, param := range method.Params {
			if def := d.Index.Find(param.Type); def != nil {
				defs = append(defs, def)
			}
		}
		if def := d.Index.Find(method.Type); def != nil {
			defs = append(defs, def)
		}
	}

	for _, prop := range d.Properties {
		if prop.Type == "EventHandler" {
			def, err := d.findEvent(prop.EventHandler)
			if err != nil {
				return defs, err
			} else if def != nil {
				defs = append(defs, def)
			}
			continue
		}

		if def := d.Index.Find(prop.Type); def != nil {
			defs = append(defs, def)
		}
	}

	for _, event := range d.Events {
		if def := d.Index.Find(event.Type); def != nil {
			defs = append(defs, def)
		}
	}

	return defs, nil
}

// find the event, traversing up if necessary
func (d *Interface) findEvent(name string) (def.Definition, error) {
	for _, event := range d.Events {
		if event.Name == name {
			if e, isset := d.Index[event.Type]; isset {
				return e, nil
			}
		}
	}

	parents, err := d.Parents()
	if err != nil {
		return nil, err
	}

	for _, parent := range parents {
		if t, ok := parent.(*Interface); ok {
			return t.findEvent(name)
		}
	}

	// return the default event
	return d.Index["Event"], nil
}

// Constant struct
type Constant struct {
	Name         string `xml:"name,attr"`
	Type         string `xml:"type,attr"`
	TypeOriginal string `xml:"type-original,attr"`
	Value        string `xml:"value,attr"`
}

// Constructor struct
type Constructor struct {
	Params []Param `xml:"param"`
}

// NamedConstructor struct
type NamedConstructor struct {
	Name   string  `xml:"name,attr"`
	Params []Param `xml:"param"`
}

// AnonymousContentAttributes struct
type AnonymousContentAttributes struct {
	Name        string `xml:"name,attr"`
	EnumValues  string `xml:"enum-values,attr"`
	ValueSyntax string `xml:"value-syntax,attr"`
}

// Event struct
type Event struct {
	Name        string `xml:"name,attr"`
	Bubbles     bool   `xml:"bubbles,attr"`
	Cancelable  bool   `xml:"cancelable,attr"`
	Tags        string `xml:"tags,attr"`
	Dispatch    string `xml:"dispatch,attr"`
	Follows     string `xml:"follows,attr"`
	Precedes    string `xml:"precedes,attr"`
	SkipsWindow bool   `xml:"skips-window,attr"`
	Type        string `xml:"type,attr"`
	Aliases     string `xml:"aliases,attr"`
}

// Element struct
type Element struct {
	Name            string `xml:"name,attr"`
	HTMLSelfClosing bool   `xml:"html-self-closing,attr"`
	Namespace       string `xml:"namespace,attr"`
}

// Method struct
type Method struct {
	Name                     string  `xml:"name,attr"`
	Type                     string  `xml:"type,attr"`
	Creator                  bool    `xml:"creator,attr"`
	Setter                   bool    `xml:"setter,attr"`
	DoNotCheckDomainSecurity bool    `xml:"do-not-check-domain-security,attr"`
	Params                   []Param `xml:"param"`

	Comment string
}

// Property struct
type Property struct {
	Name                              string `xml:"name,attr"`
	Type                              string `xml:"type,attr"`
	ReadOnly                          bool   `xml:"read-only,attr"`
	ContentAttribute                  string `xml:"content-attribute,attr"`
	ContentAttributeReflects          bool   `xml:"content-attribute-reflects,attr"`
	ContentAttributeValueSyntax       string `xml:"content-attribute-value-syntax,attr"`
	ContentAttributeEnumValues        string `xml:"content-attribute-enum-values,attr"`
	EventHandler                      string `xml:"event-handler,attr"`
	TreatNullAs                       string `xml:"treat-null-as,attr"`
	Nullable                          bool   `xml:"nullable,attr"`
	Replaceable                       bool   `xml:"replaceable,attr"`
	PropertyDescriptorNotConfigurable bool   `xml:"property-descriptor-not-configurable,attr"`
	DoNotCheckDomainSecurity          bool   `xml:"do-not-check-domain-security,attr"`
	Unforgeable                       bool   `xml:"unforgeable,attr"`

	Comment string
}

// Param struct
type Param struct {
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr,omitempty"`
	Optional bool   `xml:"optional,attr,omitempty"`
	Variadic bool   `xml:"variadic,attr,omitempty"`
}
