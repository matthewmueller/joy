package xmlserializer

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// XMLSerializer struct
// js:"XMLSerializer,omit"
type XMLSerializer struct {
}

// SerializeToString fn
// js:"serializeToString"
func (*XMLSerializer) SerializeToString(target window.Node) (s string) {
	js.Rewrite("$_.serializeToString($1)", target)
	return s
}
