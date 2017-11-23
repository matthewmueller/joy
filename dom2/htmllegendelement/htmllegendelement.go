package htmllegendelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLLegendElement struct
// js:"HTMLLegendElement,omit"
type HTMLLegendElement struct {
	window.HTMLElement
}

// Align Retrieves a reference to the form that the object is embedded in.
func (*HTMLLegendElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Retrieves a reference to the form that the object is embedded in.
func (*HTMLLegendElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// Form Retrieves a reference to the form that the object is embedded in.
func (*HTMLLegendElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}
