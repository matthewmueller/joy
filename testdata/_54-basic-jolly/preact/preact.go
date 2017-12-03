package preact

import (
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/js"
	"github.com/matthewmueller/joy/vdom"
)

// var preact = js.RawFile("./preact.js")

// var _ jsx.Node = (*Component)(nil)

// Component interface
// js:"component,omit"
// type Component struct{}

// // Render fn
// func (c *Component) Render() jsx.JSX {
// 	return nil
// }

// // SetState fn
// // js:"setState"
// func (c *Component) SetState(interface{}) {

// }

// // ForceUpdate fn
// func (c *Component) ForceUpdate() {

// }

// Render the component
func Render(component vdom.Child, el *document.Node) {
	js.Rewrite("$1.render($2, $3)", vdom.File(), component, el)
}
