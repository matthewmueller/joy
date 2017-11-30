package domparser

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// DOMParser struct
// js:"DOMParser,omit"
type DOMParser struct {
}

// ParseFromString fn
// js:"parseFromString"
func (*DOMParser) ParseFromString(source string, mimeType string) (w window.Document) {
	js.Rewrite("$_.parseFromString($1, $2)", source, mimeType)
	return w
}
