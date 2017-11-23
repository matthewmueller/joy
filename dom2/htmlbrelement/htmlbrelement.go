package htmlbrelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLBRElement,omit"
type HTMLBRElement struct {
	window.HTMLElement
}

// Clear Sets or retrieves the side on which floating objects are not to be positioned when any IHTMLBlockElement is inserted into the document.
func (*HTMLBRElement) Clear() (clear string) {
	js.Rewrite("$<.Clear")
	return clear
}

// Clear Sets or retrieves the side on which floating objects are not to be positioned when any IHTMLBlockElement is inserted into the document.
func (*HTMLBRElement) SetClear(clear string) {
	js.Rewrite("$<.Clear = $1", clear)
}
