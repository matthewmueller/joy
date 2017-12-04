package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	"github.com/apex/log"
	"github.com/knq/snaker"
)

func main() {
	if e := generate(); e != nil {
		log.WithError(e).Fatal("error generating")
	}
}

// DB struct
type DB struct {
	Tags  []string `json:"tags"`
	Index [][]int  `json:"index"`
	All   []string `json:"all"`
	Attrs []string `json:"attributes"`
}

func generate() error {
	buf, err := ioutil.ReadFile("./elements.json")
	if err != nil {
		return err
	}

	var db DB
	if e := json.Unmarshal(buf, &db); e != nil {
		return e
	}

	tags := map[string][]string{}
	for i, attrs := range db.Index {
		tag := db.Tags[i]

		tags[tag] = append(tags[tag], db.All...)

		if len(attrs) > 0 {
			for _, attr := range attrs {
				tags[tag] = append(tags[tag], db.Attrs[attr])
			}
		}
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	out := map[string]string{}
	for tag, attrs := range tags {
		code, err := render(&data{
			Tag:   tag,
			Attrs: attrs,
		})
		if err != nil {
			return err
		}

		formatted, err := format(code)
		if err != nil {
			return err
		}

		out[tag] = formatted
	}

	for tag, code := range out {
		dir := path.Join(pwd, "h", tag)
		if e := os.MkdirAll(dir, os.ModePerm); e != nil {
			return e
		}

		if e := ioutil.WriteFile(path.Join(dir, tag+".go"), []byte(code), os.ModePerm); e != nil {
			return e
		}
	}

	return nil
}

type data struct {
	Tag   string
	Attrs []string
}

func render(data *data) (string, error) {
	// tpl, err := template.New(name).Funcs(templateMap).Parse(string(raw))
	tpl, err := template.New(data.Tag).Funcs(fns()).Parse(pkg())
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	if e := tpl.Execute(&b, data); e != nil {
		return "", e
	}

	return string(b.Bytes()), nil
}

func fns() template.FuncMap {
	return template.FuncMap{
		"capitalize": func(s string) string {
			return snaker.SnakeToCamelIdentifier(snaker.CamelToSnake(s))
		},
		"lowercase": func(s string) string {
			if builtins[s] != "" {
				return builtins[s]
			}
			return strings.ToLower(snaker.SnakeToCamelIdentifier(s))
		},
	}
}

// format the output using goimports
func format(input string) (output string, err error) {
	cmd := exec.Command("goimports")
	stdin, err := cmd.StdinPipe()

	if err != nil {
		return output, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return output, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return output, err
	}

	reader := bytes.NewBufferString(input)

	if e := cmd.Start(); e != nil {
		return output, e
	}

	io.Copy(stdin, reader)
	stdin.Close()

	formatted, err := ioutil.ReadAll(stdout)
	if err != nil {
		return output, err
	}

	formattingError, err := ioutil.ReadAll(stderr)
	if err != nil {
		return output, err
	}

	stderr.Close()
	stdout.Close()

	if e := cmd.Wait(); e != nil {
		return output, errors.New(string(formattingError))
	}

	return string(formatted), nil
}

func pkg() string {
	return `
{{ $c := capitalize .Tag }}
{{ $l := lowercase .Tag }}
package {{ $l }}

import (
	"encoding/json"
	"strings"

	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/vdom"
)

// {{ $c }} struct
// js:"{{ $l }},omit"
type {{ $c }} struct {
	vdom.Child
	vdom.Node

	attrs    map[string]interface{}
	children []vdom.Child
}

// Props struct
// js:"props,omit"
type Props struct {
	attrs map[string]interface{}
}

// New fn
func New(props *Props, children ...vdom.Child) *{{ $c }} {
	macro.Rewrite("$1('{{ .Tag }}', $2 ? $2.JSON() : {}, $3)", vdom.Pragma(), props, children)
	if props == nil {
		props = &Props{attrs: map[string]interface{}{}}
	}
	return &{{ $c }}{
		attrs:    props.attrs,
		children: children,
	}
}

// Render fn
func (s *{{ $c }}) Render() vdom.Node {
	return s
}

func (s *{{ $c }}) String() string {
	// props
	var props []string
	for key, val := range s.attrs {
		bytes, e := json.Marshal(val)
		// TODO: skip over errors?
		if e != nil {
			continue
		}
		props = append(props, key+"="+string(bytes))
	}

	// children
	var children []string
	for _, child := range s.children {
		children = append(children, child.Render().String())
	}

	if len(props) > 0 {
		return "<{{ $l }} " + strings.Join(props, " ") + ">" + strings.Join(children, "") + "</{{ $l }}>"
	}

	return "<{{ $l }}>" + strings.Join(children, "") + "</{{ $l }}>"
}

{{ range $i, $attr := .Attrs }}
{{- $ac := capitalize $attr -}}
{{- $al := lowercase $attr -}}
// {{ $ac }} fn
func {{ $ac }}({{ $al }} string) *Props {
	macro.Rewrite("$1().Set('{{ $al }}', $2)", macro.Runtime("Map", "Set", "JSON"), {{ $al }})
	p := &Props{attrs: map[string]interface{}{}}
	return p.{{ $ac }}({{ $al }})
}

// {{ $ac }} fn
func (p *Props) {{ $ac }}({{ $al }} string) *Props {
	macro.Rewrite("$_.Set('{{ $al }}', $1)", {{ $al }})
	p.attrs["{{ $attr }}"] = {{ $al }}
	return p
}
{{ end }}

// Attr fn
func Attr(key string, value interface{}) *Props {
	macro.Rewrite("$1().Set($2, $3)", macro.Runtime("Map", "Set", "JSON"), key, value)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Attr(key, value)
}

// Attr fn
func (p *Props) Attr(key string, value interface{}) *Props {
	macro.Rewrite("$_.Set($1, $2)", key, value)
	p.attrs[key] = value
	return p
}
`
}

var builtins = map[string]string{
	"bool":        "boolean",
	"byte":        "b",
	"complex64":   "c64",
	"complex128":  "c128",
	"error":       "err",
	"float32":     "f32",
	"float64":     "f64",
	"int":         "integer",
	"int8":        "integer8",
	"int16":       "integer16",
	"int32":       "integer32",
	"int64":       "integer64",
	"rune":        "ru",
	"string":      "str",
	"uint":        "uinteger",
	"uint8":       "uinteger8",
	"uint16":      "uinteger16",
	"uint32":      "uinteger32",
	"uint64":      "uinteger64",
	"uintptr":     "uintpointer",
	"true":        "yes",
	"false":       "no",
	"iota":        "ita",
	"nil":         "null",
	"append":      "app",
	"cap":         "capacity",
	"close":       "cls",
	"complex":     "cpx",
	"copy":        "cpy",
	"delete":      "del",
	"imag":        "img",
	"len":         "l",
	"make":        "mk",
	"new":         "nw",
	"panic":       "pnc",
	"print":       "prt",
	"println":     "prtln",
	"real":        "rl",
	"recover":     "rec",
	"break":       "brk",
	"default":     "def",
	"func":        "fn",
	"interface":   "iface",
	"select":      "sel",
	"case":        "cs",
	"defer":       "dfr",
	"go":          "g",
	"map":         "mp",
	"struct":      "structure",
	"chan":        "chn",
	"else":        "els",
	"goto":        "gto",
	"package":     "pkg",
	"switch":      "swch",
	"const":       "cst",
	"fallthrough": "flth",
	"if":          "ifs",
	"range":       "rng",
	"type":        "kind",
	"continue":    "cont",
	"for":         "fors",
	"import":      "imp",
	"return":      "ret",
	"var":         "v",
}
