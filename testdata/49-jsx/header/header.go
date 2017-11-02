package header

import (
	"github.com/matthewmueller/golly/jsx"
)

// Header struct
type header struct {
	title    string
	children []jsx.Component
	bats     string
}

// New Header
func New(title string, children ...jsx.Component) jsx.Component {
	return &header{
		title:    title,
		children: children,
		bats:     "are crazy",
	}
}

// Render header
// js:"render"
func (d *header) Render() jsx.JSX {
	return jsx.H("h3", map[string]interface{}{"class": d.title}, d.children[0])
}
