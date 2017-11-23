package htmltablesectionelement

import (
	"github.com/matthewmueller/golly/dom2/htmltablealignment"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLTableSectionElement,omit"
type HTMLTableSectionElement struct {
	window.HTMLElement
	htmltablealignment.HTMLTableAlignment
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
