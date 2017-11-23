package domparser

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New() *DOMParser {
	js.Rewrite("DOMParser")
	return &DOMParser{}
}

// DOMParser struct
// js:"DOMParser,omit"
type DOMParser struct {
}

// ParseFromString
func (*DOMParser) ParseFromString(source string, mimeType string) (w window.Document) {
	js.Rewrite("$<.ParseFromString($1, $2)", source, mimeType)
	return w
}
