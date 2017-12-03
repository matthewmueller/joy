package preact

import (
	"github.com/matthewmueller/joy/js"
	"github.com/matthewmueller/joy/testdata/49-jsx/window"
	"github.com/matthewmueller/joy/vdom"
)

// Render the component
func Render(component vdom.Child, el window.Node) {
	js.Rewrite("$1.render($2, $3)", vdom.File(), component, el)
}
