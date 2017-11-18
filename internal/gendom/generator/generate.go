package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/matthewmueller/golly/internal/gen"
)

// Generate the DOM
func Generate(src string) (files []gen.File, err error) {
	dom, err := parse(src)
	if err != nil {
		return files, err
	}

	idx := newIndex(dom)

	// 2. generate
	var stmts []string

	// generate our enums
	// TODO: actually generate enums
	for _, enum := range dom.Enums {
		str, e := enum.Generate(idx)
		if e != nil {
			return files, e
		}
		stmts = append(stmts, str)
	}

	// generate our dictionaries
	for _, dict := range dom.Dictionaries {
		str, e := dict.Generate(idx)
		if e != nil {
			return files, e
		}
		stmts = append(stmts, str)
	}

	// generate our interfaces
	for _, iface := range dom.Interfaces {
		str, e := iface.Generate(idx, nil)
		if e != nil {
			return files, e
		}
		stmts = append(stmts, str)
	}

	// generate our mixin interfaces
	for _, iface := range dom.MixinInterfaces {
		str, e := iface.Generate(idx, nil)
		if e != nil {
			return files, e
		}
		stmts = append(stmts, str)
	}

	result := "package dom\n\n" + strings.Join(stmts, "\n")

	output, e := gen.Format(result)
	if e != nil {
		if e := ioutil.WriteFile("generated.go", []byte(result), os.ModePerm); e != nil {
			return files, e
		}
		return files, e
	}
	_ = output

	return files, nil
}

// Generate enums
func (e *Enum) Generate(idx *index) (string, error) {
	return gen.Generate("enum/"+e.Name, e, `type {{ .Name }} string`)
}

type dictionaryData struct {
	Name    string
	Extends string
	Members []string
}

// Generate the Dictionary
func (d *Dictionary) Generate(idx *index) (string, error) {
	data := &dictionaryData{}
	data.Name = d.Name

	if d.Extends != "" && d.Extends != "Object" {
		data.Extends = gen.Pointer(d.Extends)
	}

	for _, member := range d.Members {
		m, e := member.Generate(idx)
		if e != nil {
			return "", e
		}
		data.Members = append(data.Members, m)
	}

	return gen.Generate("dictionary/"+d.Name, data, `
		type {{ .Name }} struct {
			{{ .Extends }}

			{{ range .Members }}
			{{ . }}
			{{- end }}
		}
	`)
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
		data.Extends = gen.Pointer(i.Extends)
	}

	for _, imp := range i.Implements {
		data.Implements = append(data.Implements, gen.Pointer(imp))
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

	return gen.Generate("interface/"+i.Name, data, `
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

type methodData struct {
	Recv   string
	Name   string
	Params []gen.Vartype
	Result gen.Vartype
}

// Generate fn
func (m *Method) Generate(idx *index, recv *Interface) (string, error) {
	data := methodData{}
	data.Recv = gen.Pointer(recv.Name)
	data.Name = m.Name

	for _, param := range m.Params {
		// handle callback interfaces
		if idx.CallbackInterfaces[param.Type] != nil {
			fn := idx.CallbackInterfaces[param.Type]
			t, e := fn.Generate(idx)
			if e != nil {
				return "", e
			}

			data.Params = append(data.Params, gen.Vartype{
				Var:  gen.Name(param.Name),
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

			data.Params = append(data.Params, gen.Vartype{
				Var:  gen.Name(param.Name),
				Type: t,
			})
			continue
		}

		t, e := coerce(idx, param.Type)
		if e != nil {
			return "", e
		}

		data.Params = append(data.Params, gen.Vartype{
			Var:  gen.Name(param.Name),
			Type: t,
		})
	}

	t, e := coerce(idx, m.Type)
	if e != nil {
		return "", e
	}
	data.Result = gen.Vartype{
		Var:  gen.Variable(t),
		Type: t,
	}

	async := isAsync(m.Type)

	if t == "void" {
		if async {
			return gen.Generate("method/"+m.Name, data, `
				func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) {
					js.Rewrite("await $<.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
				}
			`)
		}

		return gen.Generate("method/"+m.Name, data, `
			func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) {
				js.Rewrite("$<.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
			}
		`)
	}

	if async {
		return gen.Generate("method/"+m.Name, data, `
			func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) ({{ .Result.Var }} {{ .Result.Type }}) {
				js.Rewrite("await $<.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
				return {{ .Result.Var }}
			}
		`)
	}
	return gen.Generate("method/"+m.Name, data, `
		func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) ({{ .Result.Var }} {{ .Result.Type }}) {
			js.Rewrite("$<.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
			return {{ .Result.Var }}
		}
	`)
}

type paramData struct {
	Name string
	Type string
}

// Generate fn
func (p *Param) Generate(idx *index) (string, error) {
	data := paramData{}
	data.Name = gen.Name(p.Name)

	t, e := coerce(idx, p.Type)
	if e != nil {
		return "", e
	}
	data.Type = t

	return gen.Generate("param/"+p.Name, data, `{{ .Name }} {{ .Type }}`)
}

type propertyData struct {
	Recv   string
	Name   string
	Result gen.Vartype
}

// Generate fn
func (p *Property) Generate(idx *index, recv *Interface) (string, error) {
	var result []string

	data := propertyData{}
	data.Recv = gen.Pointer(recv.Name)
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

		t, e := coerce(idx, event.Name)
		if e != nil {
			return "", e
		}

		data.Result = gen.Vartype{
			Var:  gen.Name(n),
			Type: t,
		}
	} else if idx.CallbackFunctions[p.Type] != nil {
		fn := idx.CallbackFunctions[p.Type]
		t, e := fn.Generate(idx)
		if e != nil {
			return "", e
		}

		data.Result = gen.Vartype{
			Var:  gen.Name(fn.Name),
			Type: t,
		}
	} else {
		t, e := coerce(idx, p.Type)
		if e != nil {
			return "", e
		}
		data.Result = gen.Vartype{
			Var:  gen.Name(p.Name),
			Type: t,
		}
	}

	// only one known instance of this (property that returns a Promise that returns undefined)
	// so we'll just special case this for now
	async := isAsync(p.Type)
	if async && data.Result.Type == "void" {
		getter, e := gen.Generate("property_getter/"+p.Name, data, `
			func ({{ .Recv }}) Get{{ capitalize .Name }}() {
				js.Rewrite("await $<.{{ .Name }}")
			}
		`)
		if e != nil {
			return "", e
		}
		result = append(result, getter)
	} else {
		getter, e := gen.Generate("property_getter/"+p.Name, data, `
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
		setter, e := gen.Generate("property_setter/"+p.Name, data, `
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
		data.Extends = gen.Pointer(i.Extends)
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

	t, e := coerce(idx, method.Type)
	if e != nil {
		return "", e
	}
	data.Result = t

	if t == "void" {
		return gen.Generate("callback_interface/"+i.Name, data, `func ({{ join .Params }})`)
	}

	return gen.Generate("callback_interface/"+i.Name, data, `func ({{ join .Params }}) ({{ .Result }})`)
}

type memberData struct {
	Name string
	Type string
}

// Generate fn
func (m *Member) Generate(idx *index) (string, error) {
	data := memberData{}
	data.Name = gen.Name(m.Name)

	t, e := coerce(idx, m.Type)
	if e != nil {
		return "", e
	}

	// make the optional fields pointers
	if m.Nullable || !m.Required {
		t = gen.Pointer(t)
	}
	data.Type = t

	return gen.Generate("member/"+m.Name, data, `{{ .Name }} {{ .Type }}`)
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

	t, e := coerce(idx, c.Type)
	if e != nil {
		return "", e
	}
	data.Result = t

	if t == "void" {
		return gen.Generate("callback_fn/"+c.Name, data, `func ({{ join .Params }})`)
	}

	return gen.Generate("callback_fn/"+c.Name, data, `func ({{ join .Params }}) ({{ .Result }})`)
}
