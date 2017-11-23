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

// CheckValidity Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLFieldSetElement) CheckValidity() (b bool) {
	js.Rewrite("$<.CheckValidity()")
	return b
}

// SetCustomValidity Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
func (*HTMLFieldSetElement) SetCustomValidity(err string) {
	js.Rewrite("$<.SetCustomValidity($1)", err)
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLFieldSetElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align Sets or retrieves how the object is aligned with adjacent text.
func (*HTMLFieldSetElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// Disabled
func (*HTMLFieldSetElement) Disabled() (disabled bool) {
	js.Rewrite("$<.Disabled")
	return disabled
}

// Disabled
func (*HTMLFieldSetElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.Disabled = $1", disabled)
}

// Form Retrieves a reference to the form that the object is embedded in.
func (*HTMLFieldSetElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}

// Name
func (*HTMLFieldSetElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name
func (*HTMLFieldSetElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// ValidationMessage Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
func (*HTMLFieldSetElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.ValidationMessage")
	return validationMessage
}

// Validity Returns a  ValidityState object that represents the validity states of an element.
func (*HTMLFieldSetElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.Validity")
	return validity
}

// WillValidate Returns whether an element will successfully validate based on forms validation rules and constraints.
func (*HTMLFieldSetElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.WillValidate")
	return willValidate
}
