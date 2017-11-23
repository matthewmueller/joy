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

// CreateCaption Creates an empty caption element in the table.
func (*HTMLTableElement) CreateCaption() (w window.HTMLElement) {
	js.Rewrite("$<.CreateCaption()")
	return w
}

// CreateTBody Creates an empty tBody element in the table.
func (*HTMLTableElement) CreateTBody() (w window.HTMLElement) {
	js.Rewrite("$<.CreateTBody()")
	return w
}

// CreateTFoot Creates an empty tFoot element in the table.
func (*HTMLTableElement) CreateTFoot() (w window.HTMLElement) {
	js.Rewrite("$<.CreateTFoot()")
	return w
}

// CreateTHead Returns the tHead element object if successful, or null otherwise.
func (*HTMLTableElement) CreateTHead() (w window.HTMLElement) {
	js.Rewrite("$<.CreateTHead()")
	return w
}

// DeleteCaption Deletes the caption element and its contents from the table.
func (*HTMLTableElement) DeleteCaption() {
	js.Rewrite("$<.DeleteCaption()")
}

// DeleteRow Removes the specified row (tr) from the element and from the rows collection.
//     * @param index Number that specifies the zero-based position in the rows collection of the row to remove.
func (*HTMLTableElement) DeleteRow(index *int) {
	js.Rewrite("$<.DeleteRow($1)", index)
}

// DeleteTFoot Deletes the tFoot element and its contents from the table.
func (*HTMLTableElement) DeleteTFoot() {
	js.Rewrite("$<.DeleteTFoot()")
}

// DeleteTHead Deletes the tHead element and its contents from the table.
func (*HTMLTableElement) DeleteTHead() {
	js.Rewrite("$<.DeleteTHead()")
}

// InsertRow Creates a new row (tr) in the table, and adds the row to the rows collection.
//     * @param index Number that specifies where to insert the row in the rows collection. The default value is -1, which appends the new row to the end of the rows collection.
func (*HTMLTableElement) InsertRow(index *int) (w window.HTMLElement) {
	js.Rewrite("$<.InsertRow($1)", index)
	return w
}

// Align Sets or retrieves a value that indicates the table alignment.
func (*HTMLTableElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves a value that indicates the table alignment.
func (*HTMLTableElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// BgColor
func (*HTMLTableElement) BgColor() (bgColor interface{}) {
	js.Rewrite("$<.BgColor")
	return bgColor
}

// BgColor
func (*HTMLTableElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.BgColor = $1", bgColor)
}

// Border Sets or retrieves the width of the border to draw around the object.
func (*HTMLTableElement) Border() (border string) {
	js.Rewrite("$<.Border")
	return border
}

// Border Sets or retrieves the width of the border to draw around the object.
func (*HTMLTableElement) SetBorder(border string) {
	js.Rewrite("$<.Border = $1", border)
}

// BorderColor Sets or retrieves the border color of the object.
func (*HTMLTableElement) BorderColor() (borderColor interface{}) {
	js.Rewrite("$<.BorderColor")
	return borderColor
}

// BorderColor Sets or retrieves the border color of the object.
func (*HTMLTableElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$<.BorderColor = $1", borderColor)
}

// Caption Retrieves the caption object of a table.
func (*HTMLTableElement) Caption() (caption *htmltablecaptionelement.HTMLTableCaptionElement) {
	js.Rewrite("$<.Caption")
	return caption
}

// Caption Retrieves the caption object of a table.
func (*HTMLTableElement) SetCaption(caption *htmltablecaptionelement.HTMLTableCaptionElement) {
	js.Rewrite("$<.Caption = $1", caption)
}

// CellPadding Sets or retrieves the amount of space between the border of the cell and the content of the cell.
func (*HTMLTableElement) CellPadding() (cellPadding string) {
	js.Rewrite("$<.CellPadding")
	return cellPadding
}

// CellPadding Sets or retrieves the amount of space between the border of the cell and the content of the cell.
func (*HTMLTableElement) SetCellPadding(cellPadding string) {
	js.Rewrite("$<.CellPadding = $1", cellPadding)
}

// CellSpacing Sets or retrieves the amount of space between cells in a table.
func (*HTMLTableElement) CellSpacing() (cellSpacing string) {
	js.Rewrite("$<.CellSpacing")
	return cellSpacing
}

// CellSpacing Sets or retrieves the amount of space between cells in a table.
func (*HTMLTableElement) SetCellSpacing(cellSpacing string) {
	js.Rewrite("$<.CellSpacing = $1", cellSpacing)
}

// Cols Sets or retrieves the number of columns in the table.
func (*HTMLTableElement) Cols() (cols int) {
	js.Rewrite("$<.Cols")
	return cols
}

// Cols Sets or retrieves the number of columns in the table.
func (*HTMLTableElement) SetCols(cols int) {
	js.Rewrite("$<.Cols = $1", cols)
}

// Frame Sets or retrieves the way the border frame around the table is displayed.
func (*HTMLTableElement) Frame() (frame string) {
	js.Rewrite("$<.Frame")
	return frame
}

// Frame Sets or retrieves the way the border frame around the table is displayed.
func (*HTMLTableElement) SetFrame(frame string) {
	js.Rewrite("$<.Frame = $1", frame)
}

// Height Sets or retrieves the height of the object.
func (*HTMLTableElement) Height() (height interface{}) {
	js.Rewrite("$<.Height")
	return height
}

// Height Sets or retrieves the height of the object.
func (*HTMLTableElement) SetHeight(height interface{}) {
	js.Rewrite("$<.Height = $1", height)
}

// Rows Sets or retrieves the number of horizontal rows contained in the object.
func (*HTMLTableElement) Rows() (rows window.HTMLCollection) {
	js.Rewrite("$<.Rows")
	return rows
}

// Rules Sets or retrieves which dividing lines (inner borders) are displayed.
func (*HTMLTableElement) Rules() (rules string) {
	js.Rewrite("$<.Rules")
	return rules
}

// Rules Sets or retrieves which dividing lines (inner borders) are displayed.
func (*HTMLTableElement) SetRules(rules string) {
	js.Rewrite("$<.Rules = $1", rules)
}

// Summary Sets or retrieves a description and/or structure of the object.
func (*HTMLTableElement) Summary() (summary string) {
	js.Rewrite("$<.Summary")
	return summary
}

// Summary Sets or retrieves a description and/or structure of the object.
func (*HTMLTableElement) SetSummary(summary string) {
	js.Rewrite("$<.Summary = $1", summary)
}

// TBodies Retrieves a collection of all tBody objects in the table. Objects in this collection are in source order.
func (*HTMLTableElement) TBodies() (tBodies window.HTMLCollection) {
	js.Rewrite("$<.TBodies")
	return tBodies
}

// TFoot Retrieves the tFoot object of the table.
func (*HTMLTableElement) TFoot() (tFoot *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$<.TFoot")
	return tFoot
}

// TFoot Retrieves the tFoot object of the table.
func (*HTMLTableElement) SetTFoot(tFoot *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$<.TFoot = $1", tFoot)
}

// THead Retrieves the tHead object of the table.
func (*HTMLTableElement) THead() (tHead *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$<.THead")
	return tHead
}

// THead Retrieves the tHead object of the table.
func (*HTMLTableElement) SetTHead(tHead *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$<.THead = $1", tHead)
}

// Width Sets or retrieves the width of the object.
func (*HTMLTableElement) Width() (width string) {
	js.Rewrite("$<.Width")
	return width
}

// Width Sets or retrieves the width of the object.
func (*HTMLTableElement) SetWidth(width string) {
	js.Rewrite("$<.Width = $1", width)
}
