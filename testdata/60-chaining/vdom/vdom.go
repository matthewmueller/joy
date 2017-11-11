package vdom

import "github.com/matthewmueller/golly/jsx"
import "github.com/matthewmueller/golly/js"

// Text struct
// js:"text,omit"
type Text struct {
	jsx.Node
	jsx.JSX

	value string
}

// T fn
func T(text string) *Text {
	js.Rewrite("$1", text)
	return &Text {
		value: text,
	}
}

// Render fn
// js:"render"
func (t *Text) Render() jsx.JSX {
	return t
}

// String fn
func (t *Text) String() string {
	return t.value
}