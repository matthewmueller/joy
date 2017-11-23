package htmlulistelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLUListElement,omit"
type HTMLUListElement struct {
	window.HTMLElement
}

// Compact
func (*HTMLUListElement) Compact() (compact bool) {
	js.Rewrite("$<.Compact")
	return compact
}

// Compact
func (*HTMLUListElement) SetCompact(compact bool) {
	js.Rewrite("$<.Compact = $1", compact)
}

// Type
func (*HTMLUListElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type
func (*HTMLUListElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}
