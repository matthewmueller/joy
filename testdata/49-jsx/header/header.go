package header

import (
	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/49-jsx/h2"
)

// Header struct
type Header struct {
	Title    string
	Children []jsx.Component
}

// Render header
func (d *Header) Render() jsx.JSX {
	return &h2.H2{Class: d.Title, Children: d.Children}
}
