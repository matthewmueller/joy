package htmldlistelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLDListElement struct
// js:"HTMLDListElement,omit"
type HTMLDListElement struct {
	window.HTMLElement
}

// Compact prop
func (*HTMLDListElement) Compact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

// Compact prop
func (*HTMLDListElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}
