package xmlserializer

import "github.com/matthewmueller/golly/js"

// js:"XMLSerializer,omit"
type XMLSerializer struct {
}

// SerializeToString
func (*XMLSerializer) SerializeToString(target window.Node) (s string) {
	js.Rewrite("$<.SerializeToString($1)", target)
	return s
}
