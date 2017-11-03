package preact

import (
	"github.com/matthewmueller/golly/dom/document"
	"github.com/matthewmueller/golly/js"
	"github.com/matthewmueller/golly/jsx"
)

var preact = js.RawFile("./preact.js")

// Render the component
func Render(component jsx.Component, el *document.Element) {
	js.Rewrite("$1.render($2, $3)", preact, component, el)
}

// // H fn
// func H(name string, attrs map[string]interface{}, children ...jsx.Component) jsx.Element {
// 	js.Rewrite("$1.h($2, $3, $4)", preact, name, attrs, children)
// 	return nil
// }
