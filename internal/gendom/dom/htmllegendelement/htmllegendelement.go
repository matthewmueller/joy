package htmllegendelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/js"
)

type HTMLLegendElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLLegendElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLLegendElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLLegendElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}
