package xmlserializer

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
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
func (*XMLSerializer) SerializeToString(target window.Node) (s string) {
	js.Rewrite("$<.serializeToString($1)", target)
	return s
}
