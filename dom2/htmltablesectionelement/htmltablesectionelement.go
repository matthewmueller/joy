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

// DeleteRow Removes the specified row (tr) from the element and from the rows collection.
//     * @param index Number that specifies the zero-based position in the rows collection of the row to remove.
func (*HTMLTableSectionElement) DeleteRow(index *int) {
	js.Rewrite("$<.DeleteRow($1)", index)
}

// InsertRow Creates a new row (tr) in the table, and adds the row to the rows collection.
//     * @param index Number that specifies where to insert the row in the rows collection. The default value is -1, which appends the new row to the end of the rows collection.
func (*HTMLTableSectionElement) InsertRow(index *int) (w window.HTMLElement) {
	js.Rewrite("$<.InsertRow($1)", index)
	return w
}

// Ch
func (*HTMLTableSectionElement) Ch() (ch string) {
	js.Rewrite("$<.Ch")
	return ch
}

// Ch
func (*HTMLTableSectionElement) SetCh(ch string) {
	js.Rewrite("$<.Ch = $1", ch)
}

// ChOff
func (*HTMLTableSectionElement) ChOff() (chOff string) {
	js.Rewrite("$<.ChOff")
	return chOff
}

// ChOff
func (*HTMLTableSectionElement) SetChOff(chOff string) {
	js.Rewrite("$<.ChOff = $1", chOff)
}

// VAlign
func (*HTMLTableSectionElement) VAlign() (vAlign string) {
	js.Rewrite("$<.VAlign")
	return vAlign
}

// VAlign
func (*HTMLTableSectionElement) SetVAlign(vAlign string) {
	js.Rewrite("$<.VAlign = $1", vAlign)
}

// Align Sets or retrieves a value that indicates the table alignment.
func (*HTMLTableSectionElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves a value that indicates the table alignment.
func (*HTMLTableSectionElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// Rows Sets or retrieves the number of horizontal rows contained in the object.
func (*HTMLTableSectionElement) Rows() (rows window.HTMLCollection) {
	js.Rewrite("$<.Rows")
	return rows
}
