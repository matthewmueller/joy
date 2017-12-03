package preact

import (
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/vdom"
)

var _ vdom.Component = (*Component)(nil)

// Component interface
// js:"component,omit"
type Component struct{}

// Render fn
func (c *Component) Render() vdom.Node {
	return nil
}

// SetState fn
// js:"setState"
func (c *Component) SetState(interface{}) {

}

// ForceUpdate fn
func (c *Component) ForceUpdate() {

}

// Render the component
func Render(component vdom.Component, el *document.Node) {
	macro.Rewrite("$1.render($2, $3)", macro.RawFile("./preact.js"), component, el)
}

// String turns the component into a string
func String(component vdom.Component) string {
	return component.Render().String()
}
