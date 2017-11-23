package xmlserializer

import (
	"github.com/matthewmueller/golly/dom2/window"
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

// SerializeToString
func (*XMLSerializer) SerializeToString(target window.Node) (s string) {
	js.Rewrite("$<.SerializeToString($1)", target)
	return s
}
