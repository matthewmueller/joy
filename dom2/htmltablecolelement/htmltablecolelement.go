package htmltablecolelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTableColElement struct
// js:"HTMLTableColElement,omit"
type HTMLTableColElement struct {
	window.HTMLElement
}

// Ch
func (*HTMLTableColElement) Ch() (ch string) {
	js.Rewrite("$<.Ch")
	return ch
}

// Ch
func (*HTMLTableColElement) SetCh(ch string) {
	js.Rewrite("$<.Ch = $1", ch)
}

// ChOff
func (*HTMLTableColElement) ChOff() (chOff string) {
	js.Rewrite("$<.ChOff")
	return chOff
}

// ChOff
func (*HTMLTableColElement) SetChOff(chOff string) {
	js.Rewrite("$<.ChOff = $1", chOff)
}

// VAlign
func (*HTMLTableColElement) VAlign() (vAlign string) {
	js.Rewrite("$<.VAlign")
	return vAlign
}

// VAlign
func (*HTMLTableColElement) SetVAlign(vAlign string) {
	js.Rewrite("$<.VAlign = $1", vAlign)
}

// Align Sets or retrieves the alignment of the object relative to the display or table.
func (*HTMLTableColElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves the alignment of the object relative to the display or table.
func (*HTMLTableColElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// Span Sets or retrieves the number of columns in the group.
func (*HTMLTableColElement) Span() (span uint) {
	js.Rewrite("$<.Span")
	return span
}

// Span Sets or retrieves the number of columns in the group.
func (*HTMLTableColElement) SetSpan(span uint) {
	js.Rewrite("$<.Span = $1", span)
}

// Width Sets or retrieves the width of the object.
func (*HTMLTableColElement) Width() (width interface{}) {
	js.Rewrite("$<.Width")
	return width
}

// Width Sets or retrieves the width of the object.
func (*HTMLTableColElement) SetWidth(width interface{}) {
	js.Rewrite("$<.Width = $1", width)
}
