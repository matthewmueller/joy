package htmlulistelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLUListElement struct
// js:"HTMLUListElement,omit"
type HTMLUListElement struct {
	window.HTMLElement
}

// Compact prop
func (*HTMLUListElement) Compact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

// Compact prop
func (*HTMLUListElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}

// Type prop
func (*HTMLUListElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop
func (*HTMLUListElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}
