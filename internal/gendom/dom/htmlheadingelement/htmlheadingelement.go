package htmlheadingelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLHeadingElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLHeadingElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLHeadingElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}
