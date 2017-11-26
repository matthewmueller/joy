package preact

import (
	"github.com/matthewmueller/golly/js"
	"github.com/matthewmueller/golly/testdata/49-jsx/window"
	"github.com/matthewmueller/golly/vdom"
)

// Render the component
func Render(component vdom.Child, el window.Node) {
	js.Rewrite("$1.render($2, $3)", vdom.File(), component, el)
}
