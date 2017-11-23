package htmlfieldsetelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/validitystate"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLFieldSetElement struct
// js:"HTMLFieldSetElement,omit"
type HTMLFieldSetElement struct {
	window.HTMLElement
}

// CheckValidity fn Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLFieldSetElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

// SetCustomValidity fn Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
func (*HTMLFieldSetElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLFieldSetElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLFieldSetElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// Disabled prop
func (*HTMLFieldSetElement) Disabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

// Disabled prop
func (*HTMLFieldSetElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

// Form prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLFieldSetElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

// Name prop
func (*HTMLFieldSetElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop
func (*HTMLFieldSetElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// ValidationMessage prop Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
func (*HTMLFieldSetElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

// Validity prop Returns a  ValidityState object that represents the validity states of an element.
func (*HTMLFieldSetElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

// WillValidate prop Returns whether an element will successfully validate based on forms validation rules and constraints.
func (*HTMLFieldSetElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}
