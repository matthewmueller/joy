package htmlmenuelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLMenuElement struct
// js:"HTMLMenuElement,omit"
type HTMLMenuElement struct {
	window.HTMLElement
}

// Compact prop
func (*HTMLMenuElement) Compact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

// Compact prop
func (*HTMLMenuElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}

// Type prop
func (*HTMLMenuElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop
func (*HTMLMenuElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
