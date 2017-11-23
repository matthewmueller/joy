package htmlmenuelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLMenuElement struct
// js:"HTMLMenuElement,omit"
type HTMLMenuElement struct {
	window.HTMLElement
}

// Compact
func (*HTMLMenuElement) Compact() (compact bool) {
	js.Rewrite("$<.Compact")
	return compact
}

// Compact
func (*HTMLMenuElement) SetCompact(compact bool) {
	js.Rewrite("$<.Compact = $1", compact)
}

// Type
func (*HTMLMenuElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type
func (*HTMLMenuElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}
