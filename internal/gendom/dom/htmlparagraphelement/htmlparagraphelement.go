package htmlparagraphelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLParagraphElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLParagraphElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLParagraphElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLParagraphElement) GetClear() (clear string) {
	js.Rewrite("$<.clear")
	return clear
}

func (*HTMLParagraphElement) SetClear(clear string) {
	js.Rewrite("$<.clear = $1", clear)
}
