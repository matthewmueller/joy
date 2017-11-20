package htmltablecellelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmltablealignment"
	"github.com/matthewmueller/golly/js"
)

type HTMLTableCellElement struct {
	*htmlelement.HTMLElement
	*htmltablealignment.HTMLTableAlignment
}

func (*HTMLTableCellElement) GetAbbr() (abbr string) {
	js.Rewrite("$<.abbr")
	return abbr
}

func (*HTMLTableCellElement) SetAbbr(abbr string) {
	js.Rewrite("$<.abbr = $1", abbr)
}

func (*HTMLTableCellElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLTableCellElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLTableCellElement) GetAxis() (axis string) {
	js.Rewrite("$<.axis")
	return axis
}

func (*HTMLTableCellElement) SetAxis(axis string) {
	js.Rewrite("$<.axis = $1", axis)
}

func (*HTMLTableCellElement) GetBgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

func (*HTMLTableCellElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

func (*HTMLTableCellElement) GetCellIndex() (cellIndex int) {
	js.Rewrite("$<.cellIndex")
	return cellIndex
}

func (*HTMLTableCellElement) GetColSpan() (colSpan uint) {
	js.Rewrite("$<.colSpan")
	return colSpan
}

func (*HTMLTableCellElement) SetColSpan(colSpan uint) {
	js.Rewrite("$<.colSpan = $1", colSpan)
}

func (*HTMLTableCellElement) GetHeaders() (headers string) {
	js.Rewrite("$<.headers")
	return headers
}

func (*HTMLTableCellElement) SetHeaders(headers string) {
	js.Rewrite("$<.headers = $1", headers)
}

func (*HTMLTableCellElement) GetHeight() (height interface{}) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLTableCellElement) SetHeight(height interface{}) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLTableCellElement) GetNoWrap() (noWrap bool) {
	js.Rewrite("$<.noWrap")
	return noWrap
}

func (*HTMLTableCellElement) SetNoWrap(noWrap bool) {
	js.Rewrite("$<.noWrap = $1", noWrap)
}

func (*HTMLTableCellElement) GetRowSpan() (rowSpan uint) {
	js.Rewrite("$<.rowSpan")
	return rowSpan
}

func (*HTMLTableCellElement) SetRowSpan(rowSpan uint) {
	js.Rewrite("$<.rowSpan = $1", rowSpan)
}

func (*HTMLTableCellElement) GetScope() (scope string) {
	js.Rewrite("$<.scope")
	return scope
}

func (*HTMLTableCellElement) SetScope(scope string) {
	js.Rewrite("$<.scope = $1", scope)
}

func (*HTMLTableCellElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLTableCellElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}
