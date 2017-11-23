package htmlolistelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLOListElement struct
// js:"HTMLOListElement,omit"
type HTMLOListElement struct {
	window.HTMLElement
}

// Compact prop
func (*HTMLOListElement) Compact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

// Compact prop
func (*HTMLOListElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}

// Start prop The starting number.
func (*HTMLOListElement) Start() (start int) {
	js.Rewrite("$<.start")
	return start
}

// Start prop The starting number.
func (*HTMLOListElement) SetStart(start int) {
	js.Rewrite("$<.start = $1", start)
}

// Type prop
func (*HTMLOListElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop
func (*HTMLOListElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
