package htmltablecaptionelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLTableCaptionElement,omit"
type HTMLTableCaptionElement struct {
	window.HTMLElement
}

// Align Sets or retrieves the alignment of the caption or legend.
func (*HTMLTableCaptionElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves the alignment of the caption or legend.
func (*HTMLTableCaptionElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// VAlign Sets or retrieves whether the caption appears at the top or bottom of the table.
func (*HTMLTableCaptionElement) VAlign() (vAlign string) {
	js.Rewrite("$<.VAlign")
	return vAlign
}

// VAlign Sets or retrieves whether the caption appears at the top or bottom of the table.
func (*HTMLTableCaptionElement) SetVAlign(vAlign string) {
	js.Rewrite("$<.VAlign = $1", vAlign)
}
