package htmlquoteelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLQuoteElement struct
// js:"HTMLQuoteElement,omit"
type HTMLQuoteElement struct {
	window.HTMLElement
}

// Cite prop Sets or retrieves reference information about the object.
func (*HTMLQuoteElement) Cite() (cite string) {
	js.Rewrite("$<.cite")
	return cite
}

// Cite prop Sets or retrieves reference information about the object.
func (*HTMLQuoteElement) SetCite(cite string) {
	js.Rewrite("$<.cite = $1", cite)
}
