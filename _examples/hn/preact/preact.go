package preact

import (
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/jsx"
)

var preact = macro.File("./preact.js")

var _ jsx.Node = (*Component)(nil)

// Component interface
// js:"component,omit"
type Component struct{}

// Render fn
func (c *Component) Render() jsx.JSX {
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
func Render(component jsx.Node, el *document.Node) {
	macro.Rewrite("$1.render($2, $3)", preact, component, el)
}
