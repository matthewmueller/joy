package htmltablecolelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmltablealignment"
	"github.com/matthewmueller/golly/js"
)

type HTMLTableColElement struct {
	*htmlelement.HTMLElement
	*htmltablealignment.HTMLTableAlignment
}

func (*HTMLTableColElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLTableColElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLTableColElement) GetSpan() (span uint) {
	js.Rewrite("$<.span")
	return span
}

func (*HTMLTableColElement) SetSpan(span uint) {
	js.Rewrite("$<.span = $1", span)
}

func (*HTMLTableColElement) GetWidth() (width interface{}) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLTableColElement) SetWidth(width interface{}) {
	js.Rewrite("$<.width = $1", width)
}
