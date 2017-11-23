package htmltableelement

import (
	"github.com/matthewmueller/golly/dom2/htmltablecaptionelement"
	"github.com/matthewmueller/golly/dom2/htmltablesectionelement"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTableElement struct
// js:"HTMLTableElement,omit"
type HTMLTableElement struct {
	window.HTMLElement
}

// CreateCaption fn Creates an empty caption element in the table.
func (*HTMLTableElement) CreateCaption() (w window.HTMLElement) {
	js.Rewrite("$<.createCaption()")
	return w
}

// CreateTBody fn Creates an empty tBody element in the table.
func (*HTMLTableElement) CreateTBody() (w window.HTMLElement) {
	js.Rewrite("$<.createTBody()")
	return w
}

// CreateTFoot fn Creates an empty tFoot element in the table.
func (*HTMLTableElement) CreateTFoot() (w window.HTMLElement) {
	js.Rewrite("$<.createTFoot()")
	return w
}

// CreateTHead fn Returns the tHead element object if successful, or null otherwise.
func (*HTMLTableElement) CreateTHead() (w window.HTMLElement) {
	js.Rewrite("$<.createTHead()")
	return w
}

// DeleteCaption fn Deletes the caption element and its contents from the table.
func (*HTMLTableElement) DeleteCaption() {
	js.Rewrite("$<.deleteCaption()")
}

// DeleteRow fn Removes the specified row (tr) from the element and from the rows collection.
//     * @param index Number that specifies the zero-based position in the rows collection of the row to remove.
func (*HTMLTableElement) DeleteRow(index *int) {
	js.Rewrite("$<.deleteRow($1)", index)
}

// DeleteTFoot fn Deletes the tFoot element and its contents from the table.
func (*HTMLTableElement) DeleteTFoot() {
	js.Rewrite("$<.deleteTFoot()")
}

// DeleteTHead fn Deletes the tHead element and its contents from the table.
func (*HTMLTableElement) DeleteTHead() {
	js.Rewrite("$<.deleteTHead()")
}

// InsertRow fn Creates a new row (tr) in the table, and adds the row to the rows collection.
//     * @param index Number that specifies where to insert the row in the rows collection. The default value is -1, which appends the new row to the end of the rows collection.
func (*HTMLTableElement) InsertRow(index *int) (w window.HTMLElement) {
	js.Rewrite("$<.insertRow($1)", index)
	return w
}

// Align prop Sets or retrieves a value that indicates the table alignment.
func (*HTMLTableElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves a value that indicates the table alignment.
func (*HTMLTableElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// BgColor prop
func (*HTMLTableElement) BgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

// BgColor prop
func (*HTMLTableElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

// Border prop Sets or retrieves the width of the border to draw around the object.
func (*HTMLTableElement) Border() (border string) {
	js.Rewrite("$<.border")
	return border
}

// Border prop Sets or retrieves the width of the border to draw around the object.
func (*HTMLTableElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

// BorderColor prop Sets or retrieves the border color of the object.
func (*HTMLTableElement) BorderColor() (borderColor interface{}) {
	js.Rewrite("$<.borderColor")
	return borderColor
}

// BorderColor prop Sets or retrieves the border color of the object.
func (*HTMLTableElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$<.borderColor = $1", borderColor)
}

// Caption prop Retrieves the caption object of a table.
func (*HTMLTableElement) Caption() (caption *htmltablecaptionelement.HTMLTableCaptionElement) {
	js.Rewrite("$<.caption")
	return caption
}

// Caption prop Retrieves the caption object of a table.
func (*HTMLTableElement) SetCaption(caption *htmltablecaptionelement.HTMLTableCaptionElement) {
	js.Rewrite("$<.caption = $1", caption)
}

// CellPadding prop Sets or retrieves the amount of space between the border of the cell and the content of the cell.
func (*HTMLTableElement) CellPadding() (cellPadding string) {
	js.Rewrite("$<.cellPadding")
	return cellPadding
}

// CellPadding prop Sets or retrieves the amount of space between the border of the cell and the content of the cell.
func (*HTMLTableElement) SetCellPadding(cellPadding string) {
	js.Rewrite("$<.cellPadding = $1", cellPadding)
}

// CellSpacing prop Sets or retrieves the amount of space between cells in a table.
func (*HTMLTableElement) CellSpacing() (cellSpacing string) {
	js.Rewrite("$<.cellSpacing")
	return cellSpacing
}

// CellSpacing prop Sets or retrieves the amount of space between cells in a table.
func (*HTMLTableElement) SetCellSpacing(cellSpacing string) {
	js.Rewrite("$<.cellSpacing = $1", cellSpacing)
}

// Cols prop Sets or retrieves the number of columns in the table.
func (*HTMLTableElement) Cols() (cols int) {
	js.Rewrite("$<.cols")
	return cols
}

// Cols prop Sets or retrieves the number of columns in the table.
func (*HTMLTableElement) SetCols(cols int) {
	js.Rewrite("$<.cols = $1", cols)
}

// Frame prop Sets or retrieves the way the border frame around the table is displayed.
func (*HTMLTableElement) Frame() (frame string) {
	js.Rewrite("$<.frame")
	return frame
}

// Frame prop Sets or retrieves the way the border frame around the table is displayed.
func (*HTMLTableElement) SetFrame(frame string) {
	js.Rewrite("$<.frame = $1", frame)
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLTableElement) Height() (height interface{}) {
	js.Rewrite("$<.height")
	return height
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLTableElement) SetHeight(height interface{}) {
	js.Rewrite("$<.height = $1", height)
}

// Rows prop Sets or retrieves the number of horizontal rows contained in the object.
func (*HTMLTableElement) Rows() (rows window.HTMLCollection) {
	js.Rewrite("$<.rows")
	return rows
}

// Rules prop Sets or retrieves which dividing lines (inner borders) are displayed.
func (*HTMLTableElement) Rules() (rules string) {
	js.Rewrite("$<.rules")
	return rules
}

// Rules prop Sets or retrieves which dividing lines (inner borders) are displayed.
func (*HTMLTableElement) SetRules(rules string) {
	js.Rewrite("$<.rules = $1", rules)
}

// Summary prop Sets or retrieves a description and/or structure of the object.
func (*HTMLTableElement) Summary() (summary string) {
	js.Rewrite("$<.summary")
	return summary
}

// Summary prop Sets or retrieves a description and/or structure of the object.
func (*HTMLTableElement) SetSummary(summary string) {
	js.Rewrite("$<.summary = $1", summary)
}

// TBodies prop Retrieves a collection of all tBody objects in the table. Objects in this collection are in source order.
func (*HTMLTableElement) TBodies() (tBodies window.HTMLCollection) {
	js.Rewrite("$<.tBodies")
	return tBodies
}

// TFoot prop Retrieves the tFoot object of the table.
func (*HTMLTableElement) TFoot() (tFoot *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$<.tFoot")
	return tFoot
}

// TFoot prop Retrieves the tFoot object of the table.
func (*HTMLTableElement) SetTFoot(tFoot *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$<.tFoot = $1", tFoot)
}

// THead prop Retrieves the tHead object of the table.
func (*HTMLTableElement) THead() (tHead *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$<.tHead")
	return tHead
}

// THead prop Retrieves the tHead object of the table.
func (*HTMLTableElement) SetTHead(tHead *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$<.tHead = $1", tHead)
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLTableElement) Width() (width string) {
	js.Rewrite("$<.width")
	return width
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLTableElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}
