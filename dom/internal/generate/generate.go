package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/knq/snaker"
	"github.com/matthewmueller/golly/dom/internal/generate/builtin"
	"github.com/matthewmueller/golly/dom/internal/generate/formatter"
)

var resequence = regexp.MustCompile(`sequence<([\w<>]+)>`)
var repromise = regexp.MustCompile(`Promise<([\w<>]+)>`)

var typemap = map[string]string{
	"DOMString":        "string",
	"USVString":        "string",
	"IDBKeyPath":       "string",
	"AAGUID":           "string",
	"EndOfStreamError": "string",
	"ReadyState":       "string",

	"boolean":                "bool",
	"bool":                   "bool",
	"Boolean":                "bool",
	"MSAttestationStatement": "bool",

	"unsigned long":      "uint",
	"unsigned long long": "uint64",
	"unsigned short":     "uint8",
	"Uint8Array":         "[]uint8",
	"Uint8ClampedArray":  "[]uint8",

	"short":               "int8",
	"long":                "int",
	"long long":           "int64",
	"DOMTimeStamp":        "int",
	"DOMHighResTimeStamp": "int",
	"Int32Array":          "[]int32",

	"float":               "float32",
	"double":              "float32",
	"unrestricted double": "float32",
	"UnrestrictedDouble":  "float32",
	"Float32Array":        "[]float32",

	"void": "void",

	"object":       "interface{}",
	"any":          "interface{}",
	"Dictionary":   "interface{}",
	"payloadType":  "interface{}",
	"Entry":        "interface{}",
	"Transferable": "interface{}",
	"BodyInit":     "interface{}",

	"ArrayBuffer":     "[]byte",
	"ArrayBufferView": "[]byte",
	"BufferSource":    "[]byte",
	"octet":           "byte",

	"EventHandler": "EventHandler",

	"RequestInfo": "*Request",

	"function": "func()",
	"Function": "func()",

	"Date": "time.Time",

	"WebKitEntry[]": "[]*WebKitEntry",
}

type vartype struct {
	Var  string
	Type string
}

// Type transforms the XML type to a Go type
func Type(idx *index, kind string) (string, error) {
	kind = strings.TrimSpace(kind)
	isSlice := false

	// TODO: handle unions better
	if strings.Contains(kind, " or ") {
		return "interface{}", nil
	}

	matches := repromise.FindStringSubmatch(kind)
	if len(matches) > 1 {
		kind = matches[1]
	}

	matches = resequence.FindStringSubmatch(kind)
	if len(matches) > 1 {
		isSlice = true
		kind = matches[1]
	}

	if typemap[kind] == "" {
		kind = pointer(kind)
	} else {
		kind = typemap[kind]
	}

	if isSlice {
		return "[]" + kind, nil
	}
	return kind, nil
}

func isAsync(kind string) bool {
	return strings.Contains(kind, "Promise<")
}

// CallbackFunction struct
type CallbackFunction struct {
	XMLName  xml.Name `xml:"callback-function"`
	Name     string   `xml:"name,attr"`
	Callback bool     `xml:"callback,attr,omitempty"`
	Type     string   `xml:"type,attr"`
	Params   []Param  `xml:"param"`
}

// callbackFunctionData struct
type callbackFunctionData struct {
	Params []string
	Result string
}

// Generate fn
func (c *CallbackFunction) Generate(idx *index) (string, error) {
	data := callbackFunctionData{}

	for _, param := range c.Params {
		p, e := param.Generate(idx)
		if e != nil {
			return "", e
		}
		data.Params = append(data.Params, p)
	}

	t, e := Type(idx, c.Type)
	if e != nil {
		return "", e
	}
	data.Result = t

	if t == "void" {
		return generate("callback_fn/"+c.Name, data, `func ({{ join .Params }})`)
	}

	return generate("callback_fn/"+c.Name, data, `func ({{ join .Params }}) ({{ .Result }})`)
}

// Param struct
type Param struct {
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr,omitempty"`
	Optional bool   `xml:"optional,attr,omitempty"`
	Variadic bool   `xml:"variadic,attr,omitempty"`
}

type paramData struct {
	Name string
	Type string
}

// Generate fn
func (p *Param) Generate(idx *index) (string, error) {
	data := paramData{}
	data.Name = name(p.Name)

	t, e := Type(idx, p.Type)
	if e != nil {
		return "", e
	}
	data.Type = t

	return generate("param/"+p.Name, data, `{{ .Name }} {{ .Type }}`)
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

type methodData struct {
	Recv   string
	Name   string
	Params []vartype
	Result vartype
}

// Generate fn
func (m *Method) Generate(idx *index, recv *Interface) (string, error) {
	data := methodData{}
	data.Recv = pointer(recv.Name)
	data.Name = m.Name

	for _, param := range m.Params {
		// handle callback interfaces
		if idx.CallbackInterfaces[param.Type] != nil {
			fn := idx.CallbackInterfaces[param.Type]
			t, e := fn.Generate(idx)
			if e != nil {
				return "", e
			}

			data.Params = append(data.Params, vartype{
				Var:  name(param.Name),
				Type: t,
			})
			continue
		}

		// handle callback functions
		if idx.CallbackFunctions[param.Type] != nil {
			fn := idx.CallbackFunctions[param.Type]
			t, e := fn.Generate(idx)
			if e != nil {
				return "", e
			}

			data.Params = append(data.Params, vartype{
				Var:  name(param.Name),
				Type: t,
			})
			continue
		}

		t, e := Type(idx, param.Type)
		if e != nil {
			return "", e
		}

		data.Params = append(data.Params, vartype{
			Var:  name(param.Name),
			Type: t,
		})
	}

	t, e := Type(idx, m.Type)
	if e != nil {
		return "", e
	}
	data.Result = vartype{
		Var:  variable(t),
		Type: t,
	}

	async := isAsync(m.Type)

	if t == "void" {
		if async {
			return generate("method/"+m.Name, data, `
				func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) {
					js.Rewrite("await $<.{{ .Name }}({{ len .Params | countup }})", {{ joinv .Params }})
				}
			`)
		}

		return generate("method/"+m.Name, data, `
			func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) {
				js.Rewrite("$<.{{ .Name }}({{ len .Params | countup }})", {{ joinv .Params }})
			}
		`)
	}

	if async {
		return generate("method/"+m.Name, data, `
			func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) ({{ .Result.Var }} {{ .Result.Type }}) {
				js.Rewrite("await $<.{{ .Name }}({{ len .Params | countup }})", {{ joinv .Params }})
				return {{ .Result.Var }}
			}
		`)
	}
	return generate("method/"+m.Name, data, `
		func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) ({{ .Result.Var }} {{ .Result.Type }}) {
			js.Rewrite("$<.{{ .Name }}({{ len .Params | countup }})", {{ joinv .Params }})
			return {{ .Result.Var }}
		}
	`)
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

type propertyData struct {
	Recv   string
	Name   string
	Result vartype
}

// Generate fn
func (p *Property) Generate(idx *index, recv *Interface) (string, error) {
	var result []string

	data := propertyData{}
	data.Recv = pointer(recv.Name)
	data.Name = p.Name

	if p.Type == "EventHandler" {
		n := "e"
		event := idx.Interfaces["Event"]

		ev := findEvent(idx, recv, p.EventHandler)
		if ev != nil {
			n = ev.Name
			event = idx.Interfaces[ev.Type]
		}
		if event == nil {
			return "", fmt.Errorf("%s.%s: couldn't find event name: %s", recv.Name, p.Name, p.EventHandler)
		}

		t, e := Type(idx, event.Name)
		if e != nil {
			return "", e
		}

		data.Result = vartype{
			Var:  name(n),
			Type: t,
		}
	} else if idx.CallbackFunctions[p.Type] != nil {
		fn := idx.CallbackFunctions[p.Type]
		t, e := fn.Generate(idx)
		if e != nil {
			return "", e
		}

		data.Result = vartype{
			Var:  name(fn.Name),
			Type: t,
		}
	} else {
		t, e := Type(idx, p.Type)
		if e != nil {
			return "", e
		}
		data.Result = vartype{
			Var:  name(p.Name),
			Type: t,
		}
	}

	// only one known instance of this (property that returns a Promise that returns undefined)
	// so we'll just special case this for now
	async := isAsync(p.Type)
	if async && data.Result.Type == "void" {
		getter, e := generate("property_getter/"+p.Name, data, `
			func ({{ .Recv }}) Get{{ capitalize .Name }}() {
				js.Rewrite("await $<.{{ .Name }}")
			}
		`)
		if e != nil {
			return "", e
		}
		result = append(result, getter)
	} else {
		getter, e := generate("property_getter/"+p.Name, data, `
		func ({{ .Recv }}) Get{{ capitalize .Name }}() ({{ .Result.Var }} {{ .Result.Type }}) {
			js.Rewrite("$<.{{ .Name }}")
			return {{ .Result.Var }}
		}
		`)
		if e != nil {
			return "", e
		}
		result = append(result, getter)
	}

	if !p.ReadOnly {
		setter, e := generate("property_setter/"+p.Name, data, `
		func ({{ .Recv }}) Set{{ capitalize .Name }} ({{ .Result.Var }} {{ .Result.Type }}) {
			js.Rewrite("$<.{{ .Name }} = $1", {{ .Result.Var }})
		}
		`)
		if e != nil {
			return "", e
		}
		result = append(result, setter)
	}

	return strings.Join(result, "\n"), nil
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

type callbackInterfaceData struct {
	Name    string
	Extends string
	Params  []string
	Result  string
}

// Generate the type
func (i *CallbackInterface) Generate(idx *index) (string, error) {
	data := callbackInterfaceData{}
	data.Name = i.Name

	if i.Extends != "" && i.Extends != "Object" {
		data.Extends = pointer(i.Extends)
	}

	if len(i.Methods) != 1 {
		return "", fmt.Errorf("callback_interface: expected %s to only have 1 method, but it has %d methods", i.Name, len(i.Methods))
	}

	method := i.Methods[0]
	for _, param := range method.Params {
		p, e := param.Generate(idx)
		if e != nil {
			return "", e
		}
		data.Params = append(data.Params, p)
	}

	t, e := Type(idx, method.Type)
	if e != nil {
		return "", e
	}
	data.Result = t

	if t == "void" {
		return generate("callback_interface/"+i.Name, data, `func ({{ join .Params }})`)
	}

	return generate("callback_interface/"+i.Name, data, `func ({{ join .Params }}) ({{ .Result }})`)
}

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

type interfaceData struct {
	Name       string
	Extends    string
	Implements []string
	Methods    []string
	Properties []string
}

// Generate the type
func (i *Interface) Generate(idx *index, implementor *Interface) (string, error) {
	data := interfaceData{}
	data.Name = i.Name

	recv := i
	if implementor != nil {
		recv = implementor
	}

	if i.Extends != "" && i.Extends != "Object" {
		data.Extends = pointer(i.Extends)
	}

	for _, imp := range i.Implements {
		data.Implements = append(data.Implements, pointer(imp))
	}

	for _, method := range i.Methods {
		if method == nil {
			continue
		}

		m, e := method.Generate(idx, recv)
		if e != nil {
			return "", e
		}
		data.Methods = append(data.Methods, m)
	}

	for _, property := range i.Properties {
		if property == nil {
			continue
		}

		m, e := property.Generate(idx, recv)
		if e != nil {
			return "", e
		}
		data.Properties = append(data.Properties, m)
	}

	return generate("interface/"+i.Name, data, `
type {{ .Name }} struct {
	{{ .Extends }}

	{{- range .Implements }}
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

// Member struct
type Member struct {
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr,omitempty"`
	Required bool   `xml:"required,attr,omitempty"`
	Default  string `xml:"default,attr,omitempty"`
	Nullable bool   `xml:"nullable,attr,omitempty"`
}

type memberData struct {
	Name string
	Type string
}

// Generate fn
func (m *Member) Generate(idx *index) (string, error) {
	data := memberData{}
	data.Name = name(m.Name)

	t, e := Type(idx, m.Type)
	if e != nil {
		return "", e
	}

	// make the optional fields pointers
	if m.Nullable || !m.Required {
		zero, isset := builtin.Types[t]
		if isset && zero != nil {
			t = "*" + t
		}
	}
	data.Type = t

	return generate("member/"+m.Name, data, `{{ .Name }} {{ .Type }}`)
}

// Dictionary struct
type Dictionary struct {
	Name    string   `xml:"name,attr"`
	Extends string   `xml:"extends,attr"`
	Members []Member `xml:"members>member"`
}

type dictionaryData struct {
	Name    string
	Extends string
	Members []string
}

// Generate fn
func (d *Dictionary) Generate(idx *index) (string, error) {
	data := &dictionaryData{}
	data.Name = d.Name

	if d.Extends != "" && d.Extends != "Object" {
		data.Extends = pointer(d.Extends)
	}

	for _, member := range d.Members {
		m, e := member.Generate(idx)
		if e != nil {
			return "", e
		}
		data.Members = append(data.Members, m)
	}

	return generate("dictionary/"+d.Name, data, `
		type {{ .Name }} struct {
			{{ .Extends }}

			{{ range .Members }}
			{{ . }}
			{{- end }}
		}
	`)

	// return "", nil
}

// Enum struct
type Enum struct {
	Name   string   `xml:"name,attr"`
	Values []string `xml:"value"`
}

// Generate fn
func (e *Enum) Generate(idx *index) (string, error) {
	return generate("enum/"+e.Name, e, `type {{ .Name }} string`)
}

// Typedef struct
type Typedef struct {
	NewType string `xml:"new-type,attr"`
	Type    string `xml:"type,attr"`
}

// DOM struct
type DOM struct {
	WebIDL             xml.Name             `xml:"webidl-xml"`
	CallbackFunctions  []*CallbackFunction  `xml:"callback-functions>callback-function"`
	CallbackInterfaces []*CallbackInterface `xml:"callback-interfaces>interface"`
	Dictionaries       []*Dictionary        `xml:"dictionaries>dictionary"`
	Enums              []*Enum              `xml:"enums>enum"`
	Interfaces         []*Interface         `xml:"interfaces>interface"`
	MixinInterfaces    []*Interface         `xml:"mixin-interfaces>interface"`
	TypeDefs           []*Typedef           `xml:"typedefs>typedef"`
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

type index struct {
	CallbackFunctions  map[string]*CallbackFunction
	CallbackInterfaces map[string]*CallbackInterface
	Dictionaries       map[string]*Dictionary
	Enums              map[string]*Enum
	Interfaces         map[string]*Interface
	MixinInterfaces    map[string]*Interface
	TypeDefs           map[string]*Typedef
}

func main() {
	log.SetHandler(text.New(os.Stderr))

	src, err := ioutil.ReadFile("./inputs/browser.webidl.xml")
	if err != nil {
		log.WithError(err).Fatalf("error reading file")
		return
	}

	var dom DOM
	xmlDecoder := xml.NewDecoder(bytes.NewBuffer(src))
	if e := xmlDecoder.Decode(&dom); e != nil {
		log.WithError(e).Fatalf("error decoding")
		return
	}

	// append over override with added types
	buf, err := ioutil.ReadFile("./inputs/addedTypes.json")
	if err != nil {
		log.WithError(err).Fatalf("error reading file")
		return
	}

	var added []addedType
	jsonDecoder := json.NewDecoder(bytes.NewBuffer(buf))
	if e := jsonDecoder.Decode(&added); e != nil {
		log.WithError(e).Fatalf("error decoding")
		return
	}

	// 1. index
	idx := &index{
		CallbackFunctions:  map[string]*CallbackFunction{},
		CallbackInterfaces: map[string]*CallbackInterface{},
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

	idx = adjust(idx)

	// 2. generate
	var stmts []string

	// generate our enums
	// TODO: actually generate enums
	for _, enum := range dom.Enums {
		str, e := enum.Generate(idx)
		if e != nil {
			log.WithError(e).Fatalf("error generating")
		}
		stmts = append(stmts, str)
	}

	// generate our dictionaries
	for _, dict := range dom.Dictionaries {
		str, e := dict.Generate(idx)
		if e != nil {
			log.WithError(e).Fatalf("error generating")
		}
		stmts = append(stmts, str)
	}

	// generate our interfaces
	for _, iface := range dom.Interfaces {
		str, e := iface.Generate(idx, nil)
		if e != nil {
			log.WithError(e).Fatalf("error generating")
		}
		stmts = append(stmts, str)
	}

	// generate our mixin interfaces
	for _, iface := range dom.MixinInterfaces {
		str, e := iface.Generate(idx, nil)
		if e != nil {
			log.WithError(e).Fatalf("error generating")
		}
		stmts = append(stmts, str)
	}

	result := "package dom\n\n" + strings.Join(stmts, "\n")

	output, e := formatter.Format(result)
	if e != nil {
		if e := ioutil.WriteFile("generated.go", []byte(result), os.ModePerm); e != nil {
			log.WithError(e).Fatalf("unable to write")
			return
		}
		log.WithError(e).Fatalf("error formatting")
		return
	}
	if e := ioutil.WriteFile("generated.go", []byte(output), os.ModePerm); e != nil {
		log.WithError(e).Fatalf("unable to write")
		return
	}
}

// type data struct {
// 	Packages []*Package
// }

// // Package struct
// type Package struct {
// 	Name      string
// 	Variables []*Variable
// }

// // Variable struct
// type Variable struct {
// 	Name        string
// 	Rewrite     string
// 	Type        string
// 	BuiltinType bool
// }

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

func capitalize(s string) string {
	return snaker.SnakeToCamelIdentifier(snaker.CamelToSnake(s))
}

func lowercase(s string) string {
	if builtin.Identifiers[s] != "" {
		return builtin.Identifiers[s]
	}
	return strings.ToLower(snaker.SnakeToCamelIdentifier(s))
}

func name(s string) string {
	if alias, isset := builtin.Identifiers[s]; isset {
		return alias
	}
	return s
}

// Result transforms the XML type to a Go type
func variable(s string) string {
	s = strings.TrimLeft(s, "*[]")
	if len(s) == 0 {
		return s
	}
	letter := s[:1]
	return name(lowercase(letter))
}

func pointer(s string) string {
	if _, isset := builtin.Types[s]; isset {
		return s
	}
	return "*" + s
}

func joinvt(vts []vartype) string {
	var a []string
	for _, vt := range vts {
		a = append(a, vt.Var+" "+vt.Type)
	}
	return strings.Join(a, ", ")
}

func joinv(vts []vartype) string {
	var a []string
	for _, vt := range vts {
		a = append(a, vt.Var)
	}
	return strings.Join(a, ", ")
}

func join(a []string) string {
	return strings.Join(a, ", ")
}

func countup(until int) string {
	var n []string
	for i := 0; i < until; i++ {
		n = append(n, "$"+strconv.Itoa(i+1))
	}
	return strings.Join(n, ", ")
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
		"joinvt":      joinvt,
		"joinv":       joinv,
		"join":        join,
		"countup":     countup,
	}
}

func findEvent(idx *index, i *Interface, name string) *Event {

	for _, event := range i.Events {
		if event.Name == name {
			return event
		}
	}

	if i.Extends != "" && i.Extends != "Object" && idx.Interfaces[i.Extends] != nil {
		return findEvent(idx, idx.Interfaces[i.Extends], name)
	}

	return nil
}

// manual adjustments
func adjust(idx *index) *index {
	var iface *Interface

	// Remove some Pointer event properties that
	// conflict with methods of the same name
	// and aren't in the spec
	iface = idx.Interfaces["MSPointerEvent"]
	for i, prop := range iface.Properties {
		if prop.Name == "currentPoint" {
			iface.Properties[i] = nil
		}
		if prop.Name == "intermediatePoints" {
			iface.Properties[i] = nil
		}
	}
	iface = idx.Interfaces["PointerEvent"]
	for i, prop := range iface.Properties {
		if prop.Name == "currentPoint" {
			iface.Properties[i] = nil
		}
		if prop.Name == "intermediatePoints" {
			iface.Properties[i] = nil
		}
	}

	// AudioBufferSourceNode has wrong event-handler ref
	// iface = idx.Interfaces["AudioBufferSourceNode"]
	// for i, prop := range iface.Properties {
	// 	if prop.Name == "onended" {
	// 		iface.Properties[i].EventHandler = "end"
	// 	}
	// }

	// remove state change since i'm not sure where that's coming from
	// iface = idx.Interfaces["AudioContext"]
	// for i, prop := range iface.Properties {
	// 	if prop.Name == "onstatechange" {
	// 		iface.Properties[i] = nil
	// 		// iface.Properties[i].EventHandler = "end"
	// 	}
	// }

	// remove state change since i'm not sure where that's coming from
	// iface = idx.Interfaces["Document"]
	// for i, prop := range iface.Properties {
	// 	if prop.Name == "onabort" {
	// 		iface.Properties[i] = nil
	// 	}
	// 	if prop.Name == "onactivate" {
	// 		iface.Properties[i] = nil
	// 	}
	// 	if prop.Name == "onbeforeactivate" {
	// 		iface.Properties[i] = nil
	// 	}
	// 	if prop.Name == "onbeforedeactivate" {
	// 		iface.Properties[i] = nil
	// 	}
	// 	if prop.Name == "onblur" {
	// 		iface.Properties[i] = nil
	// 	}
	// 	if prop.Name == "oncanplay" {
	// 		iface.Properties[i] = nil
	// 	}
	// 	if prop.Name == "oncanplaythrough" {
	// 		iface.Properties[i] = nil
	// 	}
	// 	if prop.Name == "onchange" {
	// 		iface.Properties[i] = nil
	// 	}
	// 	if prop.Name == "onclick" {
	// 		iface.Properties[i] = nil
	// 	}
	// 	if prop.Name == "oncontextmenu" {
	// 		iface.Properties[i] = nil
	// 	}
	// }

	return idx
}
