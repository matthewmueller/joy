package htmlolistelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLOListElement struct
// js:"HTMLOListElement,omit"
type HTMLOListElement struct {
	window.HTMLElement
}

// Compact
func (*HTMLOListElement) Compact() (compact bool) {
	js.Rewrite("$<.Compact")
	return compact
}

// Compact
func (*HTMLOListElement) SetCompact(compact bool) {
	js.Rewrite("$<.Compact = $1", compact)
}

// Start The starting number.
func (*HTMLOListElement) Start() (start int) {
	js.Rewrite("$<.Start")
	return start
}

// Start The starting number.
func (*HTMLOListElement) SetStart(start int) {
	js.Rewrite("$<.Start = $1", start)
}

// Type
func (*HTMLOListElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type
func (*HTMLOListElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}
