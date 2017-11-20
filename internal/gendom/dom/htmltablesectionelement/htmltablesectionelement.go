package htmltablesectionelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlcollection"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmltablealignment"
	"github.com/matthewmueller/golly/js"
)

type HTMLTableSectionElement struct {
	*htmlelement.HTMLElement
	*htmltablealignment.HTMLTableAlignment
}

func (*HTMLTableSectionElement) DeleteRow(index int) {
	js.Rewrite("$<.deleteRow($1)", index)
}

func (*HTMLTableSectionElement) InsertRow(index int) (h *htmlelement.HTMLElement) {
	js.Rewrite("$<.insertRow($1)", index)
	return h
}

func (*HTMLTableSectionElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLTableSectionElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLTableSectionElement) GetRows() (rows *htmlcollection.HTMLCollection) {
	js.Rewrite("$<.rows")
	return rows
}
