package htmltablecaptionelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTableCaptionElement struct
// js:"HTMLTableCaptionElement,omit"
type HTMLTableCaptionElement struct {
	window.HTMLElement
}

// Align prop Sets or retrieves the alignment of the caption or legend.
func (*HTMLTableCaptionElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves the alignment of the caption or legend.
func (*HTMLTableCaptionElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// VAlign prop Sets or retrieves whether the caption appears at the top or bottom of the table.
func (*HTMLTableCaptionElement) VAlign() (vAlign string) {
	js.Rewrite("$<.vAlign")
	return vAlign
}

// VAlign prop Sets or retrieves whether the caption appears at the top or bottom of the table.
func (*HTMLTableCaptionElement) SetVAlign(vAlign string) {
	js.Rewrite("$<.vAlign = $1", vAlign)
}
