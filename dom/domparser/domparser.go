package domparser

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

// New fn
func New() *DOMParser {
	macro.Rewrite("new DOMParser()")
	return &DOMParser{}
}

// DOMParser struct
// js:"DOMParser,omit"
type DOMParser struct {
}

// ParseFromString fn
// js:"parseFromString"
func (*DOMParser) ParseFromString(source string, mimeType string) (w window.Document) {
	macro.Rewrite("$_.parseFromString($1, $2)", source, mimeType)
	return w
}
