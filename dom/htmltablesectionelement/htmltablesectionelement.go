package htmltablesectionelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTableSectionElement struct
// js:"HTMLTableSectionElement,omit"
type HTMLTableSectionElement struct {
	window.HTMLElement
}

// DeleteRow fn Removes the specified row (tr) from the element and from the rows collection.
//     * @param index Number that specifies the zero-based position in the rows collection of the row to remove.
func (*HTMLTableSectionElement) DeleteRow(index *int) {
	js.Rewrite("$<.deleteRow($1)", index)
}

// InsertRow fn Creates a new row (tr) in the table, and adds the row to the rows collection.
//     * @param index Number that specifies where to insert the row in the rows collection. The default value is -1, which appends the new row to the end of the rows collection.
func (*HTMLTableSectionElement) InsertRow(index *int) (w window.HTMLElement) {
	js.Rewrite("$<.insertRow($1)", index)
	return w
}

// Ch prop
func (*HTMLTableSectionElement) Ch() (ch string) {
	js.Rewrite("$<.ch")
	return ch
}

// Ch prop
func (*HTMLTableSectionElement) SetCh(ch string) {
	js.Rewrite("$<.ch = $1", ch)
}

// ChOff prop
func (*HTMLTableSectionElement) ChOff() (chOff string) {
	js.Rewrite("$<.chOff")
	return chOff
}

// ChOff prop
func (*HTMLTableSectionElement) SetChOff(chOff string) {
	js.Rewrite("$<.chOff = $1", chOff)
}

// VAlign prop
func (*HTMLTableSectionElement) VAlign() (vAlign string) {
	js.Rewrite("$<.vAlign")
	return vAlign
}

// VAlign prop
func (*HTMLTableSectionElement) SetVAlign(vAlign string) {
	js.Rewrite("$<.vAlign = $1", vAlign)
}

// Align prop Sets or retrieves a value that indicates the table alignment.
func (*HTMLTableSectionElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves a value that indicates the table alignment.
func (*HTMLTableSectionElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// Rows prop Sets or retrieves the number of horizontal rows contained in the object.
func (*HTMLTableSectionElement) Rows() (rows window.HTMLCollection) {
	js.Rewrite("$<.rows")
	return rows
}
