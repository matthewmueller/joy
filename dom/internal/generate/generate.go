package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/apex/log"
	"github.com/knq/snaker"
	"github.com/matthewmueller/golly/dom/internal/generate/builtin"
)

var typemap = map[string]string{
	"DOMString": "string",
	"boolean":   "bool",
}

// CallbackFunction struct
type CallbackFunction struct {
	XMLName  xml.Name `xml:"callback-function"`
	Name     string   `xml:"name,attr"`
	Callback bool     `xml:"callback,attr,omitempty"`
	Type     string   `xml:"type,attr"`
	Params   []Param  `xml:"param"`
}

// Param struct
type Param struct {
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr,omitempty"`
	Optional bool   `xml:"optional,attr,omitempty"`
	Variadic bool   `xml:"variadic,attr,omitempty"`
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

// Interface struct
type Interface struct {
	Name                       string                       `xml:"name,attr"`
	Extends                    string                       `xml:"extends,attr"`
	Implements                 []string                     `xml:"implements,omitempty"`
	NoInterfaceObject          bool                         `xml:"no-interface-object,attr"`
	Element                    []Element                    `xml:"element"`
	Methods                    []Method                     `xml:"methods>method"`
	AnonymousMethods           []Method                     `xml:"anonymous-methods>method"`
	Properties                 []Property                   `xml:"properties>property"`
	Constants                  []Constant                   `xml:"constants>constant"`
	Constructor                Constructor                  `xml:"constructor,omitempty"`
	NamedConstructor           NamedConstructor             `xml:"named-constructor"`
	Events                     []Event                      `xml:"events>event"`
	AnonymousContentAttributes []AnonymousContentAttributes `xml:"anonymous-content-attributes>parsedattribute"`
}

// Package fn
func (i *Interface) Package() string {
	return lowercase(i.Name)
}

// Type as a string
func (i *Interface) Type() string {
	return i.Name
}

// Generate the type
func (i *Interface) Generate(index *index) (string, error) {
	return generate(i.Name, i, `
	{{ .Name }}	
`)
}

// Member struct
type Member struct {
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr,omitempty"`
	Required bool   `xml:"required,attr,omitempty"`
	Default  string `xml:"default,attr,omitempty"`
	Nullable bool   `xml:"nullable,attr,omitempty"`
}

// Dictionary struct
type Dictionary struct {
	Name    string   `xml:"name,attr"`
	Extends string   `xml:"extends,attr"`
	Members []Member `xml:"members>member"`
}

// Enum struct
type Enum struct {
	Name   string   `xml:"name,attr"`
	Values []string `xml:"value"`
}

// Typedef struct
type Typedef struct {
	NewType string `xml:"new-type,attr"`
	Type    string `xml:"type,attr"`
}

// DOM struct
type DOM struct {
	WebIDL             xml.Name            `xml:"webidl-xml"`
	CallbackFunctions  []*CallbackFunction `xml:"callback-functions>callback-function"`
	CallbackInterfaces []*Interface        `xml:"callback-interfaces>interface"`
	Dictionaries       []*Dictionary       `xml:"dictionaries>dictionary"`
	Enums              []*Enum             `xml:"enums>enum"`
	Interfaces         []*Interface        `xml:"interfaces>interface"`
	MixinInterfaces    []*Interface        `xml:"mixin-interfaces>interface"`
	TypeDefs           []*Typedef          `xml:"typedefs>typedef"`
}

type index struct {
	CallbackFunctions  map[string]*CallbackFunction
	CallbackInterfaces map[string]*Interface
	Dictionaries       map[string]*Dictionary
	Enums              map[string]*Enum
	Interfaces         map[string]*Interface
	MixinInterfaces    map[string]*Interface
	TypeDefs           map[string]*Typedef
}

func main() {
	src, err := ioutil.ReadFile("./browser.webidl.xml")
	if err != nil {
		log.WithError(err).Errorf("error reading file")
		return
	}

	var dom DOM
	decoder := xml.NewDecoder(bytes.NewBuffer(src))
	if e := decoder.Decode(&dom); e != nil {
		log.WithError(e).Errorf("error decoding")
		return
	}

	// 1. index
	idx := &index{
		CallbackFunctions:  map[string]*CallbackFunction{},
		CallbackInterfaces: map[string]*Interface{},
		Dictionaries:       map[string]*Dictionary{},
		Enums:              map[string]*Enum{},
		Interfaces:         map[string]*Interface{},
		MixinInterfaces:    map[string]*Interface{},
		TypeDefs:           map[string]*Typedef{},
	}
	for _, node := range dom.CallbackFunctions {
		idx.CallbackFunctions[node.Name] = node
	}
	for _, node := range dom.CallbackInterfaces {
		idx.CallbackInterfaces[node.Name] = node
	}
	for _, node := range dom.Dictionaries {
		idx.Dictionaries[node.Name] = node
	}
	for _, node := range dom.Enums {
		idx.Enums[node.Name] = node
	}
	for _, node := range dom.Interfaces {
		idx.Interfaces[node.Name] = node
	}
	for _, node := range dom.MixinInterfaces {
		idx.MixinInterfaces[node.Name] = node
	}
	for _, node := range dom.TypeDefs {
		idx.TypeDefs[node.NewType] = node
	}

	// 2. generate

	d := &data{}
	for _, iface := range dom.Interfaces {
		if iface.Name != "Event" {
			continue
		}

		str, err := iface.Generate(idx)
		if err != nil {
			log.WithError(err).Errorf("error generating")
		}
		println(str)

		// pkg := &Package{
		// 	Name: lowercase(iface.Name),
		// }

		// for _, prop := range iface.Properties {
		// 	// if prop.Name != "document" {
		// 	// 	continue
		// 	// }

		// 	name := capitalize(prop.Name)

		// 	kind := prop.Type
		// 	if t, isset := typemap[kind]; isset {
		// 		kind = t
		// 	}

		// 	// builtinType := true
		// 	// if t, isset := builtin.Types[kind]; !isset {
		// 	// 	builtinType = false
		// 	// }

		// 	variable := &Variable{
		// 		Name: name,
		// 		Type: kind,
		// 		// BuiltinType: builtinType,
		// 		Rewrite: lowercase(iface.Name) + "." + prop.Name,
		// 	}
		// 	pkg.Variables = append(pkg.Variables, variable)
		// }

		// d.Packages = append(d.Packages, pkg)
	}

	for _, pkg := range d.Packages {
		output, err := render(pkg)
		if err != nil {
			log.WithError(err).Errorf("error generating")
			return
		}
		println(output)
	}

	// pretty.Print(dom)

	// 	cwd, err := os.Getwd()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Infof("cwd=%s", cwd)
}

type data struct {
	Packages []*Package
}

// Package struct
type Package struct {
	Name      string
	Variables []*Variable
}

// Variable struct
type Variable struct {
	Name        string
	Rewrite     string
	Type        string
	BuiltinType bool
}

func generate(name string, data interface{}, tpl string) (string, error) {
	// tpl, err := template.New(name).Funcs(templateMap).Parse(string(raw))
	t, err := template.New(name).Funcs(fns()).Parse(tpl)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	if e := t.Execute(&b, data); e != nil {
		return "", e
	}

	return string(b.Bytes()), nil
}

func render(data *Package) (string, error) {
	// tpl, err := template.New(name).Funcs(templateMap).Parse(string(raw))
	tpl, err := template.New("interfaces").Funcs(fns()).Parse(tpl())
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	if e := tpl.Execute(&b, data); e != nil {
		return "", e
	}

	return string(b.Bytes()), nil
}

func capitalize(s string) string {
	return snaker.SnakeToCamelIdentifier(snaker.CamelToSnake(s))
}

func lowercase(s string) string {
	if builtin.Identifiers[s] != "" {
		return builtin.Identifiers[s]
	}
	return strings.ToLower(snaker.SnakeToCamelIdentifier(s))
}

func pointer(s string) string {
	if _, isset := builtin.Types[s]; isset {
		return s
	}
	return "*" + lowercase(s) + "." + s
}

func dereference(s string) string {
	if t, isset := builtin.Types[s]; isset {
		switch t {
		case 0:
			return "0"
		case "":
			return "\"\""
		case nil:
			return "nil"
		case false:
			return "false"
		}
	}
	return "&" + lowercase(s) + "." + s + "{}"
}

func fns() template.FuncMap {
	return template.FuncMap{
		"capitalize":  capitalize,
		"lowercase":   lowercase,
		"pointer":     pointer,
		"dereference": dereference,
	}
}

func tpl() string {
	return `
package {{ .Name }}

{{ range .Variables }}
var {{ .Name }} = func() {{ pointer .Type }} {
	js.Rewrite("{{ .Rewrite }}")
	return {{ dereference .Type }}
}()
{{ end }}
`
}
