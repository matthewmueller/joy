package generator

import "encoding/xml"

// Interface struct
type Interface struct {
	Name                       string                        `xml:"name,attr"`
	Extends                    string                        `xml:"extends,attr"`
	Implements                 []string                      `xml:"implements,omitempty"`
	NoInterfaceObject          bool                          `xml:"no-interface-object,attr"`
	Element                    []*Element                    `xml:"element"`
	Methods                    []*Method                     `xml:"methods>method"`
	AnonymousMethods           []*Method                     `xml:"anonymous-methods>method"`
	Properties                 []*Property                   `xml:"properties>property"`
	Constants                  []*Constant                   `xml:"constants>constant"`
	Constructor                Constructor                   `xml:"constructor,omitempty"`
	NamedConstructor           NamedConstructor              `xml:"named-constructor"`
	Events                     []*Event                      `xml:"events>event"`
	AnonymousContentAttributes []*AnonymousContentAttributes `xml:"anonymous-content-attributes>parsedattribute"`
}

type addedType struct {
	Kind           string `json:"kind,omitempty"`
	Name           string `json:"name,omitempty"`
	Type           string `json:"type,omitempty"`
	Flavor         string `json:"flavor,omitempty"`
	Interface      string `json:"interface,omitempty"`
	Readonly       bool   `json:"readonly,omitempty"`
	ExposeGlobally bool   `json:"expose_globally,omitempty"`
	Extends        string `json:"extends,omitempty"`
	Properties     []struct {
		Name     string `json:"name,omitempty"`
		Kind     string `json:"kind,omitempty"`
		Readonly bool   `json:"readonly,omitempty"`
	} `json:"properties,omitempty"`
	Methods []struct {
		Name       string   `json:"name,omitempty"`
		Signatures []string `json:"signatures,omitempty"`
	} `json:"methods,omitempty"`
	Indexer []struct {
		Signatures []string `json:"signatures,omitempty"`
	} `json:"indexer,omitempty"`
}

// CallbackFunction struct
type CallbackFunction struct {
	XMLName  xml.Name `xml:"callback-function"`
	Name     string   `xml:"name,attr"`
	Callback bool     `xml:"callback,attr,omitempty"`
	Type     string   `xml:"type,attr"`
	Params   []Param  `xml:"param"`
}

// Typedef struct
type Typedef struct {
	NewType string `xml:"new-type,attr"`
	Type    string `xml:"type,attr"`
}

// Enum struct
type Enum struct {
	Name   string   `xml:"name,attr"`
	Values []string `xml:"value"`
}

// Dictionary struct
type Dictionary struct {
	Name    string   `xml:"name,attr"`
	Extends string   `xml:"extends,attr"`
	Members []Member `xml:"members>member"`
}

// Member struct
type Member struct {
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr,omitempty"`
	Required bool   `xml:"required,attr,omitempty"`
	Default  string `xml:"default,attr,omitempty"`
	Nullable bool   `xml:"nullable,attr,omitempty"`
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

// CallbackInterface struct
type CallbackInterface struct {
	Name    string    `xml:"name,attr"`
	Extends string    `xml:"extends,attr"`
	Methods []*Method `xml:"methods>method"`
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
}

// Param struct
type Param struct {
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr,omitempty"`
	Optional bool   `xml:"optional,attr,omitempty"`
	Variadic bool   `xml:"variadic,attr,omitempty"`
}
