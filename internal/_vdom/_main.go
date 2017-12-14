package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/dom/raw"
	"github.com/matthewmueller/joy/internal/gen"
	"github.com/pkg/errors"
)

func main() {
	if err := generate(); err != nil {
		log.WithError(err).Fatal("error generating")
	}
	log.Infof("done")
}

func generate() error {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("couldn't get file from runtime")
	}
	dirname := path.Dir(file)

	vdomPath := path.Join(dirname, "..", "..", "vdom")

	buf, err := ioutil.ReadFile(path.Join(dirname, "inputs", "tags.json"))
	if err != nil {
		return errors.Wrapf(err, "error reading tags.json")
	}

	var data struct {
		Tags map[string]struct {
			Comment string
			Attrs   []string
		}
		Globals []string
		Events  []string
		Types   map[string]struct {
			Type   string
			Values []string
		}
	}

	if err := json.Unmarshal(buf, &data); err != nil {
		return errors.Wrapf(err, "error unmarshalling tags")
	}

	eventmap := map[string]string{}
	for _, event := range data.Events {
		eventmap[event] = ""
	}

	// match events with event types
	// TODO: clean this up
	buf, err = ioutil.ReadFile(path.Join(dirname, "..", "dom", "inputs", "browser.webidl.xml"))
	if err != nil {
		return err
	}
	var api raw.API
	if e := xml.Unmarshal([]byte(buf), &api); e != nil {
		return errors.Wrap(e, "error parsing xml")
	}
	for _, iface := range api.Interfaces {
		for _, event := range iface.Events {
			if _, isset := eventmap["on"+event.Name]; isset {
				eventmap["on"+event.Name] = "func (e window." + event.Type + ")"
			}
		}
	}
	for name, kind := range eventmap {
		if kind == "" {
			return fmt.Errorf("missed an event in browser apis %s", name)
		}
	}

	for name, tag := range data.Tags {
		type Attr struct {
			Key   string
			Value gen.Vartype
		}

		d := struct {
			Name    string
			Comment string
			Attrs   []Attr
		}{
			Name:    name,
			Comment: tag.Comment,
		}

		attrs := append(tag.Attrs, data.Globals...)
		for _, attr := range attrs {
			parts := strings.Split(attr, ":")
			key := parts[0]
			kind := "string"

			if len(parts) > 2 {
				key = strings.Join(parts[0:len(parts)-2], ":")
				// kind = parts[len(parts)-1]
			} else if len(parts) == 2 {
				// kind = parts[1]
			}

			// if _, isset := data.Types[kind]; isset {
			// 	log.Infof("kind isset %s", kind)
			// }

			d.Attrs = append(d.Attrs, Attr{
				Key: key,
				Value: gen.Vartype{
					Var:  key,
					Type: kind,
				},
			})
		}

		for _, event := range data.Events {
			d.Attrs = append(d.Attrs, Attr{
				Key: event,
				Value: gen.Vartype{
					Var:  event,
					Type: eventmap[event],
				},
			})
		}

		code, err := gen.Generate(name, d, `
			package {{ lowercase .Name }}

			// {{ capitalize .Name }} struct
			// js:"{{ lowercase .Name }},omit"
			type {{ capitalize .Name }} struct {
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

			// New {{ lowercase .Name }} element
			// 
			// {{ .Comment }}
			func New(props *Props, children ...vdom.Child) *{{ capitalize .Name }} {
				macro.Rewrite("$1('{{ .Name }}', $2 ? $2.JSON() : {}, $3)", vdom.Pragma(), props, children)
				if props == nil {
					props = &Props{attrs: map[string]interface{}{}}
				}
				return &{{ capitalize .Name }}{
					attrs:    props.attrs,
					children: children,
				}
			}

			// Render fn
			func (s *{{ capitalize .Name }}) Render() vdom.Node {
				return s
			}

			func (s *{{ capitalize .Name }}) String() string {
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
					return "<{{ lowercase .Name }} " + strings.Join(props, " ") + ">" + strings.Join(children, "") + "</{{ lowercase .Name }}>"
				}
			
				return "<{{ lowercase .Name }}>" + strings.Join(children, "") + "</{{ lowercase .Name }}>"
			}

			{{ range .Attrs }}
			// {{ capitalize .Key }} fn
			func {{ capitalize .Key }}({{ vt .Value }}) *Props {
				macro.Rewrite("$1().Set('{{ .Key }}', $2)", macro.Runtime("Map", "Set", "JSON"), {{ identifier .Key }})
				p := &Props{attrs: map[string]interface{}{}}
				return p.{{ capitalize .Key }}({{ identifier .Key }})
			}

			// {{ capitalize .Key }} fn
			func (p *Props) {{ capitalize .Key }}({{ vt .Value }}) *Props {
				macro.Rewrite("$_.Set('{{ .Key }}', $1)", {{ identifier .Key }})
				p.attrs["{{ .Key }}"] = {{ identifier .Key }}
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
		`)
		if err != nil {
			return errors.Wrapf(err, "error generating code")
		}

		dirpath := path.Join(vdomPath, d.Name)
		if err := os.MkdirAll(dirpath, 0755); err != nil {
			return errors.Wrapf(err, "error mkdir")
		}

		if err := ioutil.WriteFile(path.Join(dirpath, d.Name+".go"), []byte(code), 0644); err != nil {
			return errors.Wrapf(err, "error writefile")
		}
	}

	// finally, copy over some supporting files
	files := []string{"text.go", "vdom.go"}
	for _, file := range files {
		buf, err := ioutil.ReadFile(path.Join(dirname, "inputs", file))
		if err != nil {
			errors.Wrapf(err, "error reading %s", file)
		}

		if err := ioutil.WriteFile(path.Join(vdomPath, file), buf, 0644); err != nil {
			return errors.Wrapf(err, "error writing file %s", file)
		}
	}

	// form	at all the files now that they're written
	if err := gen.FormatAll(vdomPath); err != nil {
		return errors.Wrapf(err, "error formatting vdom/")
	}

	return nil
}
