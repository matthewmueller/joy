package htmltablerowelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTableRowElement struct
// js:"HTMLTableRowElement,omit"
type HTMLTableRowElement struct {
	window.HTMLElement
}

// DeleteCell fn Removes the specified cell from the table row, as well as from the cells collection.
//     * @param index Number that specifies the zero-based position of the cell to remove from the table row. If no value is provided, the last cell in the cells collection is deleted.
func (*HTMLTableRowElement) DeleteCell(index *int) {
	js.Rewrite("$<.deleteCell($1)", index)
}

// InsertCell fn Creates a new cell in the table row, and adds the cell to the cells collection.
//     * @param index Number that specifies where to insert the cell in the tr. The default value is -1, which appends the new cell to the end of the cells collection.
func (*HTMLTableRowElement) InsertCell(index *int) (w window.HTMLElement) {
	js.Rewrite("$<.insertCell($1)", index)
	return w
}

// Ch prop
func (*HTMLTableRowElement) Ch() (ch string) {
	js.Rewrite("$<.ch")
	return ch
}

// Ch prop
func (*HTMLTableRowElement) SetCh(ch string) {
	js.Rewrite("$<.ch = $1", ch)
}

// ChOff prop
func (*HTMLTableRowElement) ChOff() (chOff string) {
	js.Rewrite("$<.chOff")
	return chOff
}

// ChOff prop
func (*HTMLTableRowElement) SetChOff(chOff string) {
	js.Rewrite("$<.chOff = $1", chOff)
}

// VAlign prop
func (*HTMLTableRowElement) VAlign() (vAlign string) {
	js.Rewrite("$<.vAlign")
	return vAlign
}

// VAlign prop
func (*HTMLTableRowElement) SetVAlign(vAlign string) {
	js.Rewrite("$<.vAlign = $1", vAlign)
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLTableRowElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLTableRowElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// BgColor prop
func (*HTMLTableRowElement) BgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

// BgColor prop
func (*HTMLTableRowElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

// Cells prop Retrieves a collection of all cells in the table row.
func (*HTMLTableRowElement) Cells() (cells window.HTMLCollection) {
	js.Rewrite("$<.cells")
	return cells
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLTableRowElement) Height() (height interface{}) {
	js.Rewrite("$<.height")
	return height
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLTableRowElement) SetHeight(height interface{}) {
	js.Rewrite("$<.height = $1", height)
}

// RowIndex prop Retrieves the position of the object in the rows collection for the table.
func (*HTMLTableRowElement) RowIndex() (rowIndex int) {
	js.Rewrite("$<.rowIndex")
	return rowIndex
}

// SectionRowIndex prop Retrieves the position of the object in the collection.
func (*HTMLTableRowElement) SectionRowIndex() (sectionRowIndex int) {
	js.Rewrite("$<.sectionRowIndex")
	return sectionRowIndex
}
