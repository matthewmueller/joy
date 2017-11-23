package htmltablerowelement

import (
	"github.com/matthewmueller/golly/dom2/htmltablealignment"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLTableRowElement,omit"
type HTMLTableRowElement struct {
	window.HTMLElement
	htmltablealignment.HTMLTableAlignment
}

// DeleteCell Removes the specified cell from the table row, as well as from the cells collection.
//     * @param index Number that specifies the zero-based position of the cell to remove from the table row. If no value is provided, the last cell in the cells collection is deleted.
func (*HTMLTableRowElement) DeleteCell(index *int) {
	js.Rewrite("$<.DeleteCell($1)", index)
}

// InsertCell Creates a new cell in the table row, and adds the cell to the cells collection.
//     * @param index Number that specifies where to insert the cell in the tr. The default value is -1, which appends the new cell to the end of the cells collection.
func (*HTMLTableRowElement) InsertCell(index *int) (w window.HTMLElement) {
	js.Rewrite("$<.InsertCell($1)", index)
	return w
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLTableRowElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLTableRowElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// BgColor
func (*HTMLTableRowElement) BgColor() (bgColor interface{}) {
	js.Rewrite("$<.BgColor")
	return bgColor
}

// BgColor
func (*HTMLTableRowElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.BgColor = $1", bgColor)
}

// Cells Retrieves a collection of all cells in the table row.
func (*HTMLTableRowElement) Cells() (cells window.HTMLCollection) {
	js.Rewrite("$<.Cells")
	return cells
}

// Height Sets or retrieves the height of the object.
func (*HTMLTableRowElement) Height() (height interface{}) {
	js.Rewrite("$<.Height")
	return height
}

// Height Sets or retrieves the height of the object.
func (*HTMLTableRowElement) SetHeight(height interface{}) {
	js.Rewrite("$<.Height = $1", height)
}

// RowIndex Retrieves the position of the object in the rows collection for the table.
func (*HTMLTableRowElement) RowIndex() (rowIndex int) {
	js.Rewrite("$<.RowIndex")
	return rowIndex
}

// SectionRowIndex Retrieves the position of the object in the collection.
func (*HTMLTableRowElement) SectionRowIndex() (sectionRowIndex int) {
	js.Rewrite("$<.SectionRowIndex")
	return sectionRowIndex
}
