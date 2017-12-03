package xmlserializer

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

// New fn
func New() *XMLSerializer {
	js.Rewrite("XMLSerializer")
	return &XMLSerializer{}
}

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
