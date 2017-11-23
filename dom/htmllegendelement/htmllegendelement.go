package htmllegendelement

import (
	"github.com/matthewmueller/golly/dom/htmlformelement"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLLegendElement struct
// js:"HTMLLegendElement,omit"
type HTMLLegendElement struct {
	window.HTMLElement
}

// Align prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLLegendElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLLegendElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// Form prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLLegendElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}
