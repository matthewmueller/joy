package htmlbrelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLBRElement struct
// js:"HTMLBRElement,omit"
type HTMLBRElement struct {
	window.HTMLElement
}

// Clear prop Sets or retrieves the side on which floating objects are not to be positioned when any IHTMLBlockElement is inserted into the document.
func (*HTMLBRElement) Clear() (clear string) {
	js.Rewrite("$<.clear")
	return clear
}

// Clear prop Sets or retrieves the side on which floating objects are not to be positioned when any IHTMLBlockElement is inserted into the document.
func (*HTMLBRElement) SetClear(clear string) {
	js.Rewrite("$<.clear = $1", clear)
}
