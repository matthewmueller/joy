package htmlquoteelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLQuoteElement struct
// js:"HTMLQuoteElement,omit"
type HTMLQuoteElement struct {
	window.HTMLElement
}

// Cite Sets or retrieves reference information about the object.
func (*HTMLQuoteElement) Cite() (cite string) {
	js.Rewrite("$<.Cite")
	return cite
}

// Cite Sets or retrieves reference information about the object.
func (*HTMLQuoteElement) SetCite(cite string) {
	js.Rewrite("$<.Cite = $1", cite)
}
