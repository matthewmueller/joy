package htmltablecaptionelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLTableCaptionElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLTableCaptionElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLTableCaptionElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLTableCaptionElement) GetVAlign() (vAlign string) {
	js.Rewrite("$<.vAlign")
	return vAlign
}

func (*HTMLTableCaptionElement) SetVAlign(vAlign string) {
	js.Rewrite("$<.vAlign = $1", vAlign)
}
