package main

import "strings"

// VNode interface
type VNode interface {
	String() string
}

// NewElement node
func NewElement(tagName string, attributes map[string]string, children ...VNode) *Element {
	return &Element{
		name:       tagName,
		kind:       1,
		attributes: attributes,
		children:   children,
	}
}

// Element node
type Element struct {
	name       string
	kind       int
	attributes map[string]string
	children   []VNode
}

// String fn
func (e *Element) String() string {
	var children []string
	for _, child := range e.children {
		children = append(children, child.String())
	}
	return "<" + e.name + ">" + strings.Join(children, "") + "</" + e.name + ">"
}

// // Component interface
// type Component interface {
// 	Render() VNode
// }

// // NewComponent fn
// func NewComponent(com Component) VNode {
// 	return &component{
// 		component: com,
// 	}
// }

// // component struct
// type component struct {
// 	component Component
// }

// // Render fn
// func (c *component) Render() VNode {
// 	return c.component.Render()
// }

// // String fn
// func (c *component) String() string {
// 	return c.Render().String()
// }

// // Props struct
// type Props struct {
// 	Location string
// }

// // Page struct
// type Page struct {
// 	props *Props
// }

// // New fn
// func New(p *Props) VNode {
// 	return NewComponent(&Page{p})
// }

// // Render the index page
// func (p *Page) Render() VNode {
// 	return NewElement("div", nil)
// }

func main() {
	e := NewElement("div", nil)
	println(e.String())
	// c := New(&Props{
	// 	Location: "mars",
	// })
	// println(c.String())
}
