package htmltablecolelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTableColElement struct
// js:"HTMLTableColElement,omit"
type HTMLTableColElement struct {
	window.HTMLElement
}

// Ch prop
func (*HTMLTableColElement) Ch() (ch string) {
	js.Rewrite("$<.ch")
	return ch
}

// Ch prop
func (*HTMLTableColElement) SetCh(ch string) {
	js.Rewrite("$<.ch = $1", ch)
}

// ChOff prop
func (*HTMLTableColElement) ChOff() (chOff string) {
	js.Rewrite("$<.chOff")
	return chOff
}

// ChOff prop
func (*HTMLTableColElement) SetChOff(chOff string) {
	js.Rewrite("$<.chOff = $1", chOff)
}

// VAlign prop
func (*HTMLTableColElement) VAlign() (vAlign string) {
	js.Rewrite("$<.vAlign")
	return vAlign
}

// VAlign prop
func (*HTMLTableColElement) SetVAlign(vAlign string) {
	js.Rewrite("$<.vAlign = $1", vAlign)
}

// Align prop Sets or retrieves the alignment of the object relative to the display or table.
func (*HTMLTableColElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves the alignment of the object relative to the display or table.
func (*HTMLTableColElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// Span prop Sets or retrieves the number of columns in the group.
func (*HTMLTableColElement) Span() (span uint) {
	js.Rewrite("$<.span")
	return span
}

// Span prop Sets or retrieves the number of columns in the group.
func (*HTMLTableColElement) SetSpan(span uint) {
	js.Rewrite("$<.span = $1", span)
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLTableColElement) Width() (width interface{}) {
	js.Rewrite("$<.width")
	return width
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLTableColElement) SetWidth(width interface{}) {
	js.Rewrite("$<.width = $1", width)
}
