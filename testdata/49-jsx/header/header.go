package header

import (
	"github.com/matthewmueller/golly/jsx"
)

// Header struct
type Header struct {
	*Props
}

// Props struct
type Props struct {
	Title    string
	Children []jsx.Component
}

// Render header
// js:"render"
func (d *Header) Render() jsx.JSX {
	return &jsx.Element{
		NodeName: "h3",
		Attributes: map[string]string{
			"class": d.Props.Title,
		},
		Children: d.Props.Children,
	}
}
