package strong

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
	js.Rewrite("preact.h('strong', $1.json(), $2)", props, children)
	return nil
}

// Render fn
func (s *Strong) Render() jsx.JSX {
	return nil
}

func (s *Strong) String() string {
	return "strong..."
}

// Class fn
func Class(cls string) *Props {
	js.Rewrite("map.set('class', $1)", cls)
	p := &Props{}
	return p.Class(cls)
}

// Class fn
func (p *Props) Class(cls string) *Props {
	js.Rewrite("$<.set('class', $1)", cls)
	p.attrs["class"] = cls
	return p
}

// ID fn
func (p *Props) ID(id string) *Props {
	js.Rewrite("$<.set('id', $1)",id)
	p.attrs["id"] = id
	return p
}

// Key fn
func (p *Props) Key(key string) *Props {
	js.Rewrite("$<.set('key', $1)",key)
	p.attrs["key"] = key
	return p
}