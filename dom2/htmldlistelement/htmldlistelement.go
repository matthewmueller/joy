package htmldlistelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLDListElement,omit"
type HTMLDListElement struct {
	window.HTMLElement
}

// Compact
func (*HTMLDListElement) Compact() (compact bool) {
	js.Rewrite("$<.Compact")
	return compact
}

// Compact
func (*HTMLDListElement) SetCompact(compact bool) {
	js.Rewrite("$<.Compact = $1", compact)
}
