package htmllabelelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLLabelElement struct
// js:"HTMLLabelElement,omit"
type HTMLLabelElement struct {
	window.HTMLElement
}

// Form Retrieves a reference to the form that the object is embedded in.
func (*HTMLLabelElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}

// HTMLFor Sets or retrieves the object to which the given label object is assigned.
func (*HTMLLabelElement) HTMLFor() (htmlFor string) {
	js.Rewrite("$<.HTMLFor")
	return htmlFor
}

// HTMLFor Sets or retrieves the object to which the given label object is assigned.
func (*HTMLLabelElement) SetHTMLFor(htmlFor string) {
	js.Rewrite("$<.HTMLFor = $1", htmlFor)
}
