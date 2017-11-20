package htmldivelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLDivElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLDivElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLDivElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLDivElement) GetNoWrap() (noWrap bool) {
	js.Rewrite("$<.noWrap")
	return noWrap
}

func (*HTMLDivElement) SetNoWrap(noWrap bool) {
	js.Rewrite("$<.noWrap = $1", noWrap)
}
