package h2

import (
	"github.com/matthewmueller/golly/testdata/49-jsx/jsx"
)

// MouseEvent struct
type MouseEvent struct {
	Type string
}

// H2 struct
type H2 struct {
	Class    string
	OnClick  func(MouseEvent)
	Children []jsx.Component
}

// Render fn
func (d *H2) Render() jsx.Node {
	return &jsx.Element{
		NodeName: "h2",
		Attributes: map[string]string{
			"class": d.Class,
			// "onClick": d.OnClick,
		},
		Children: d.Children,
	}
}

func (d *H2) String() string {
	return d.Render().String()
}
