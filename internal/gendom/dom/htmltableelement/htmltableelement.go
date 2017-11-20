package htmltableelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlcollection"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmltablecaptionelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmltablesectionelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLTableElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLTableElement) CreateCaption() (h *htmlelement.HTMLElement) {
	js.Rewrite("$<.createCaption()")
	return h
}

func (*HTMLTableElement) CreateTBody() (h *htmlelement.HTMLElement) {
	js.Rewrite("$<.createTBody()")
	return h
}

func (*HTMLTableElement) CreateTFoot() (h *htmlelement.HTMLElement) {
	js.Rewrite("$<.createTFoot()")
	return h
}

func (*HTMLTableElement) CreateTHead() (h *htmlelement.HTMLElement) {
	js.Rewrite("$<.createTHead()")
	return h
}

func (*HTMLTableElement) DeleteCaption() {
	js.Rewrite("$<.deleteCaption()")
}

func (*HTMLTableElement) DeleteRow(index int) {
	js.Rewrite("$<.deleteRow($1)", index)
}

func (*HTMLTableElement) DeleteTFoot() {
	js.Rewrite("$<.deleteTFoot()")
}

func (*HTMLTableElement) DeleteTHead() {
	js.Rewrite("$<.deleteTHead()")
}

func (*HTMLTableElement) InsertRow(index int) (h *htmlelement.HTMLElement) {
	js.Rewrite("$<.insertRow($1)", index)
	return h
}

func (*HTMLTableElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLTableElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLTableElement) GetBgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

func (*HTMLTableElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

func (*HTMLTableElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLTableElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLTableElement) GetBorderColor() (borderColor interface{}) {
	js.Rewrite("$<.borderColor")
	return borderColor
}

func (*HTMLTableElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$<.borderColor = $1", borderColor)
}

func (*HTMLTableElement) GetCaption() (caption *htmltablecaptionelement.HTMLTableCaptionElement) {
	js.Rewrite("$<.caption")
	return caption
}

func (*HTMLTableElement) SetCaption(caption *htmltablecaptionelement.HTMLTableCaptionElement) {
	js.Rewrite("$<.caption = $1", caption)
}

func (*HTMLTableElement) GetCellPadding() (cellPadding string) {
	js.Rewrite("$<.cellPadding")
	return cellPadding
}

func (*HTMLTableElement) SetCellPadding(cellPadding string) {
	js.Rewrite("$<.cellPadding = $1", cellPadding)
}

func (*HTMLTableElement) GetCellSpacing() (cellSpacing string) {
	js.Rewrite("$<.cellSpacing")
	return cellSpacing
}

func (*HTMLTableElement) SetCellSpacing(cellSpacing string) {
	js.Rewrite("$<.cellSpacing = $1", cellSpacing)
}

func (*HTMLTableElement) GetCols() (cols int) {
	js.Rewrite("$<.cols")
	return cols
}

func (*HTMLTableElement) SetCols(cols int) {
	js.Rewrite("$<.cols = $1", cols)
}

func (*HTMLTableElement) GetFrame() (frame string) {
	js.Rewrite("$<.frame")
	return frame
}

func (*HTMLTableElement) SetFrame(frame string) {
	js.Rewrite("$<.frame = $1", frame)
}

func (*HTMLTableElement) GetHeight() (height interface{}) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLTableElement) SetHeight(height interface{}) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLTableElement) GetRows() (rows *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.rows")
	return rows
}

func (*HTMLTableElement) GetRules() (rules string) {
	js.Rewrite("$<.rules")
	return rules
}

func (*HTMLTableElement) SetRules(rules string) {
	js.Rewrite("$<.rules = $1", rules)
}

func (*HTMLTableElement) GetSummary() (summary string) {
	js.Rewrite("$<.summary")
	return summary
}

func (*HTMLTableElement) SetSummary(summary string) {
	js.Rewrite("$<.summary = $1", summary)
}

func (*HTMLTableElement) GetTBodies() (tBodies *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.tBodies")
	return tBodies
}

func (*HTMLTableElement) GetTFoot() (tFoot *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$<.tFoot")
	return tFoot
}

func (*HTMLTableElement) SetTFoot(tFoot *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$<.tFoot = $1", tFoot)
}

func (*HTMLTableElement) GetTHead() (tHead *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$<.tHead")
	return tHead
}

func (*HTMLTableElement) SetTHead(tHead *htmltablesectionelement.HTMLTableSectionElement) {
	js.Rewrite("$<.tHead = $1", tHead)
}

func (*HTMLTableElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLTableElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}
