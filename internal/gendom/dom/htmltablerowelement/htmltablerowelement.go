package htmltablerowelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlcollection"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmltablealignment"
	"github.com/matthewmueller/golly/js"
)

type HTMLTableRowElement struct {
	*htmlelement.HTMLElement
	*htmltablealignment.HTMLTableAlignment
}

func (*HTMLTableRowElement) DeleteCell(index int) {
	js.Rewrite("$<.deleteCell($1)", index)
}

func (*HTMLTableRowElement) InsertCell(index int) (h *htmlelement.HTMLElement) {
	js.Rewrite("$<.insertCell($1)", index)
	return h
}

func (*HTMLTableRowElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLTableRowElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLTableRowElement) GetBgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

func (*HTMLTableRowElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

func (*HTMLTableRowElement) GetCells() (cells *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.cells")
	return cells
}

func (*HTMLTableRowElement) GetHeight() (height interface{}) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLTableRowElement) SetHeight(height interface{}) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLTableRowElement) GetRowIndex() (rowIndex int) {
	js.Rewrite("$<.rowIndex")
	return rowIndex
}

func (*HTMLTableRowElement) GetSectionRowIndex() (sectionRowIndex int) {
	js.Rewrite("$<.sectionRowIndex")
	return sectionRowIndex
}
