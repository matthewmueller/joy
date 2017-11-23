package htmllabelelement

import (
	"github.com/matthewmueller/golly/dom/htmlformelement"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLLabelElement struct
// js:"HTMLLabelElement,omit"
type HTMLLabelElement struct {
	window.HTMLElement
}

// Form prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLLabelElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

// HTMLFor prop Sets or retrieves the object to which the given label object is assigned.
func (*HTMLLabelElement) HTMLFor() (htmlFor string) {
	js.Rewrite("$<.htmlFor")
	return htmlFor
}

// HTMLFor prop Sets or retrieves the object to which the given label object is assigned.
func (*HTMLLabelElement) SetHTMLFor(htmlFor string) {
	js.Rewrite("$<.htmlFor = $1", htmlFor)
}
