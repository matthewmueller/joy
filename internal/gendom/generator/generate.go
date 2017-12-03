package generator

import (
	"fmt"
	"path"
	"strings"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/gen"
	"github.com/tj/go/semaphore"
)

// Generate the DOM
func Generate(src string) (files []gen.File, err error) {
	dom, e := parse(src)
	if e != nil {
		return files, e
	}

	idx := newIndex(dom)

	// generate our enums
	// TODO: actually generate enums
	for _, enum := range dom.Enums {
		src, e := enum.Generate(idx)
		if e != nil {
			return files, e
		}

		name := gen.Lowercase(enum.Name)
		files = append(files, gen.File{
			Name:   path.Join(name, name+".go"),
			Source: "package " + name + "\n\n" + src,
		})
	}

	// generate our dictionaries
	for _, dict := range dom.Dictionaries {
		src, e := dict.Generate(idx)
		if e != nil {
			return files, e
		}

		name := gen.Lowercase(dict.Name)
		files = append(files, gen.File{
			Name:   path.Join(name, name+".go"),
			Source: "package " + name + "\n\n" + src,
		})
	}

	// generate our mixin interfaces
	for _, iface := range dom.MixinInterfaces {
		src, e := iface.Generate(idx, nil)
		if e != nil {
			return files, e
		}

		name := gen.Lowercase(iface.Name)
		files = append(files, gen.File{
			Name:   path.Join(name, name+".go"),
			Source: "package " + name + "\n\n" + src,
		})
	}

	// generate our interfaces
	for _, iface := range dom.Interfaces {
		src, e := iface.Generate(idx, nil)
		if e != nil {
			return files, e
		}

		name := gen.Lowercase(iface.Name)

		files = append(files, gen.File{
			Name:   path.Join(name, name+".go"),
			Source: "package " + name + "\n\n" + src,
		})
	}

	filemap := map[string]bool{}

	sem := semaphore.New(5)
	for index, file := range files {
		if filemap[file.Name] {
			panic(file.Name + ": already exists")
		}
		filemap[file.Name] = true

		i := index
		f := file
		sem.Run(func() {
			output, e := gen.Format(f.Source)
			if e != nil {
				log.WithError(e).WithField("name", f.Name).Errorf("format error: %s\n", f.Source)
				err = e
			}
			f.Source = output
			files[i] = f
		})
	}
	sem.Wait()

	return files, err
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

	pkgname := gen.Lowercase(i.Name)

	recv := i
	if implementor != nil {
		recv = implementor
	}

	if i.Extends != "" && i.Extends != "Object" {
		data.Extends = gen.Pointer(findPackage(idx, pkgname, i.Extends))
	}

	for _, imp := range i.Implements {
		data.Implements = append(data.Implements, gen.Pointer(findPackage(idx, pkgname, imp)))
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

	pkgname := gen.Lowercase(recv.Name)

	for _, param := range m.Params {
		// handle callback interfaces
		if idx.CallbackInterfaces[param.Type] != nil {
			fn := idx.CallbackInterfaces[param.Type]
			t, e := fn.Generate(idx, recv)
			if e != nil {
				return "", e
			}

			data.Params = append(data.Params, gen.Vartype{
				Var:  gen.Identifier(param.Name),
				Type: t,
			})
			continue
		}

		// handle callback functions
		if idx.CallbackFunctions[param.Type] != nil {
			fn := idx.CallbackFunctions[param.Type]
			t, e := fn.Generate(idx, recv)
			if e != nil {
				return "", e
			}

			data.Params = append(data.Params, gen.Vartype{
				Var:  gen.Identifier(param.Name),
				Type: t,
			})
			continue
		}

		t, e := coerce(idx, pkgname, param.Type)
		if e != nil {
			return "", e
		}

		data.Params = append(data.Params, gen.Vartype{
			Var:  gen.Identifier(param.Name),
			Type: t,
		})
	}

	t, e := coerce(idx, pkgname, m.Type)
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
					js.Rewrite("await $_.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
				}
			`)
		}

		return gen.Generate("method/"+m.Name, data, `
			func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) {
				js.Rewrite("$_.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
			}
		`)
	}

	if async {
		return gen.Generate("method/"+m.Name, data, `
			func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) ({{ .Result.Var }} {{ .Result.Type }}) {
				js.Rewrite("await $_.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
				return {{ .Result.Var }}
			}
		`)
	}
	return gen.Generate("method/"+m.Name, data, `
		func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) ({{ .Result.Var }} {{ .Result.Type }}) {
			js.Rewrite("$_.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
			return {{ .Result.Var }}
		}
	`)
}

type paramData struct {
	Name string
	Type string
}

// Generate fn
func (p *Param) Generate(idx *index, recv *Interface) (string, error) {
	data := paramData{}
	data.Name = gen.Identifier(p.Name)
	pkgname := gen.Lowercase(recv.Name)

	t, e := coerce(idx, pkgname, p.Type)
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

	pkgname := gen.Lowercase(recv.Name)

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

		t, e := coerce(idx, pkgname, event.Name)
		if e != nil {
			return "", e
		}

		data.Result = gen.Vartype{
			Var:  gen.Identifier(n),
			Type: t,
		}
	} else if idx.CallbackFunctions[p.Type] != nil {
		fn := idx.CallbackFunctions[p.Type]
		t, e := fn.Generate(idx, recv)
		if e != nil {
			return "", e
		}

		data.Result = gen.Vartype{
			Var:  gen.Identifier(fn.Name),
			Type: t,
		}
	} else {
		t, e := coerce(idx, pkgname, p.Type)
		if e != nil {
			return "", e
		}
		data.Result = gen.Vartype{
			Var:  gen.Identifier(p.Name),
			Type: t,
		}
	}

	// only one known instance of this (property that returns a Promise that returns undefined)
	// so we'll just special case this for now
	async := isAsync(p.Type)
	if async && data.Result.Type == "void" {
		getter, e := gen.Generate("property_getter/"+p.Name, data, `
			func ({{ .Recv }}) Get{{ capitalize .Name }}() {
				js.Rewrite("await $_.{{ .Name }}")
			}
		`)
		if e != nil {
			return "", e
		}
		result = append(result, getter)
	} else {
		getter, e := gen.Generate("property_getter/"+p.Name, data, `
		func ({{ .Recv }}) Get{{ capitalize .Name }}() ({{ .Result.Var }} {{ .Result.Type }}) {
			js.Rewrite("$_.{{ .Name }}")
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
			js.Rewrite("$_.{{ .Name }} = $1", {{ .Result.Var }})
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
func (i *CallbackInterface) Generate(idx *index, recv *Interface) (string, error) {
	data := callbackInterfaceData{}
	data.Name = i.Name

	pkgname := gen.Lowercase(recv.Name)

	if i.Extends != "" && i.Extends != "Object" {
		data.Extends = gen.Pointer(findPackage(idx, pkgname, i.Extends))
	}

	if len(i.Methods) != 1 {
		return "", fmt.Errorf("callback_interface: expected %s to only have 1 method, but it has %d methods", i.Name, len(i.Methods))
	}

	method := i.Methods[0]
	for _, param := range method.Params {
		p, e := param.Generate(idx, recv)
		if e != nil {
			return "", e
		}
		data.Params = append(data.Params, p)
	}

	t, e := coerce(idx, pkgname, method.Type)
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
	data.Name = gen.Identifier(m.Name)

	t, e := coerce(idx, "dom", m.Type)
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
func (c *CallbackFunction) Generate(idx *index, recv *Interface) (string, error) {
	data := callbackFunctionData{}

	pkgname := gen.Lowercase(recv.Name)

	for _, param := range c.Params {
		p, e := param.Generate(idx, recv)
		if e != nil {
			return "", e
		}
		data.Params = append(data.Params, p)
	}

	t, e := coerce(idx, pkgname, c.Type)
	if e != nil {
		return "", e
	}
	data.Result = t

	if t == "void" {
		return gen.Generate("callback_fn/"+c.Name, data, `func ({{ join .Params }})`)
	}

	return gen.Generate("callback_fn/"+c.Name, data, `func ({{ join .Params }}) ({{ .Result }})`)
}
