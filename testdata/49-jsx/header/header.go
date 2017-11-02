package header

import (
	"github.com/matthewmueller/golly/testdata/49-jsx/h2"
	"github.com/matthewmueller/golly/testdata/49-jsx/jsx"
)

// Header struct
type Header struct {
	Title    string
	Children []jsx.Component
}

// Render header
func (d *Header) Render() jsx.Node {
	return &h2.H2{Class: d.Title, Children: d.Children}
}
