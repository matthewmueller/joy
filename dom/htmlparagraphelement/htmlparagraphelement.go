package htmlparagraphelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLParagraphElement struct
// js:"HTMLParagraphElement,omit"
type HTMLParagraphElement struct {
	window.HTMLElement
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLParagraphElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLParagraphElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// Clear prop
func (*HTMLParagraphElement) Clear() (clear string) {
	js.Rewrite("$<.clear")
	return clear
}

// Clear prop
func (*HTMLParagraphElement) SetClear(clear string) {
	js.Rewrite("$<.clear = $1", clear)
}
