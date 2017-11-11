package preact

import (
	"github.com/matthewmueller/golly/dom/document"
	"github.com/matthewmueller/golly/js"
	"github.com/matthewmueller/golly/jsx"
)

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
	js.Rewrite("$1.render($2, $3)", js.RawFile("./preact.js"), component, el)
}
