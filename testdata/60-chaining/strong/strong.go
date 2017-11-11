package strong

import "strings"
import "github.com/matthewmueller/golly/js"
import "github.com/matthewmueller/golly/jsx"

// Strong struct
// js:"strong,omit"
type Strong struct {
	jsx.Node
	jsx.JSX

	attrs    map[string]interface{}
	children []jsx.Node
}

// Props struct
// js:"props,omit"
type Props struct {
	attrs map[string]interface{}
}

// New fn
func New(props *Props, children ...jsx.Node) *Strong {
	js.Rewrite("$1.h('strong', $2.JSON(), $3)", js.RawFile("../preact/preact.js"), props, children)
	return nil
}

// Render fn
func (s *Strong) Render() jsx.JSX {
	return s
}

func (s *Strong) String() string {
	var children []string
	for _, child := range s.children {
		children = append(children, child.Render().String())
	}
	return "<strong>"+strings.Join(children, "")+"</strong>"
}

// Class fn
func Class(cls string) *Props {
	js.Rewrite("$1().Set('class', $2)", js.Runtime("Map", "Set", "JSON"), cls)
	p := &Props{}
	return p.Class(cls)
}

// Class fn
func (p *Props) Class(cls string) *Props {
	js.Rewrite("$<.Set('class', $1)", cls)
	p.attrs["class"] = cls
	return p
}

// ID fn
func (p *Props) ID(id string) *Props {
	js.Rewrite("$<.Set('id', $1)",id)
	p.attrs["id"] = id
	return p
}

// Key fn
func (p *Props) Key(key string) *Props {
	js.Rewrite("$<.Set('key', $1)",key)
	p.attrs["key"] = key
	return p
}