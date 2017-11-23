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

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLParagraphElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLParagraphElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// Clear
func (*HTMLParagraphElement) Clear() (clear string) {
	js.Rewrite("$<.Clear")
	return clear
}

// Clear
func (*HTMLParagraphElement) SetClear(clear string) {
	js.Rewrite("$<.Clear = $1", clear)
}
