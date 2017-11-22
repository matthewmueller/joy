package defs

import (
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/index"
	"github.com/matthewmueller/golly/internal/dom/raw"
)

var _ Method = (*method)(nil)

// NewMethod creates a new method
func NewMethod(index index.Index, data *raw.Method, receiver def.Definition, comment string) Method {
	return &method{
		index:   index,
		data:    data,
		comment: comment,
		recv:    receiver,
	}
}

// Method interface
type Method interface {
	def.Definition
}

// method struct
type method struct {
	data *raw.Method

	index   index.Index
	comment string
	recv    def.Definition
}

func (d *method) ID() string {
	return ""
}

func (d *method) Name() string {
	return d.data.Name
}

func (d *method) Kind() string {
	return "METHOD"
}

func (d *method) Generate() (string, error) {
	return "", nil
}

// Children fn
func (d *method) Children() (defs []def.Definition, e error) {
	return defs, nil
}

// // Generate fn
// func (m *method) Generate(idx *index, recv *Interface) (string, error) {
// 	data := methodData{}
// 	data.Recv = gen.Pointer(recv.Name)
// 	data.Name = m.Name

// 	pkgname := gen.Lowercase(recv.Name)

// 	for _, param := range m.Params {
// 		// handle callback interfaces
// 		if idx.CallbackInterfaces[param.Type] != nil {
// 			fn := idx.CallbackInterfaces[param.Type]
// 			t, e := fn.Generate(idx, recv)
// 			if e != nil {
// 				return "", e
// 			}

// 			data.Params = append(data.Params, gen.Vartype{
// 				Var:  gen.Name(param.Name),
// 				Type: t,
// 			})
// 			continue
// 		}

// 		// handle callback functions
// 		if idx.CallbackFunctions[param.Type] != nil {
// 			fn := idx.CallbackFunctions[param.Type]
// 			t, e := fn.Generate(idx, recv)
// 			if e != nil {
// 				return "", e
// 			}

// 			data.Params = append(data.Params, gen.Vartype{
// 				Var:  gen.Name(param.Name),
// 				Type: t,
// 			})
// 			continue
// 		}

// 		t, e := coerce(idx, pkgname, param.Type)
// 		if e != nil {
// 			return "", e
// 		}

// 		data.Params = append(data.Params, gen.Vartype{
// 			Var:  gen.Name(param.Name),
// 			Type: t,
// 		})
// 	}

// 	t, e := coerce(idx, pkgname, m.Type)
// 	if e != nil {
// 		return "", e
// 	}
// 	data.Result = gen.Vartype{
// 		Var:  gen.Variable(t),
// 		Type: t,
// 	}

// 	async := isAsync(m.Type)

// 	if t == "void" {
// 		if async {
// 			return gen.Generate("method/"+m.Name, data, `
// 				func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) {
// 					js.Rewrite("await $<.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
// 				}
// 			`)
// 		}

// 		return gen.Generate("method/"+m.Name, data, `
// 			func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) {
// 				js.Rewrite("$<.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
// 			}
// 		`)
// 	}

// 	if async {
// 		return gen.Generate("method/"+m.Name, data, `
// 			func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) ({{ .Result.Var }} {{ .Result.Type }}) {
// 				js.Rewrite("await $<.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
// 				return {{ .Result.Var }}
// 			}
// 		`)
// 	}
// 	return gen.Generate("method/"+m.Name, data, `
// 		func ({{ .Recv }}) {{ capitalize .Name }}({{ joinvt .Params }}) ({{ .Result.Var }} {{ .Result.Type }}) {
// 			js.Rewrite("$<.{{ .Name }}({{ len .Params | sequence | join }})", {{ joinv .Params }})
// 			return {{ .Result.Var }}
// 		}
// 	`)
// }
