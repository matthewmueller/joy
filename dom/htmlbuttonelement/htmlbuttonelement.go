package htmlbuttonelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/validitystate"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLButtonElement struct
// js:"HTMLButtonElement,omit"
type HTMLButtonElement struct {
	window.HTMLElement
}

// CheckValidity fn Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLButtonElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

// SetCustomValidity fn Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
func (*HTMLButtonElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

// Autofocus prop Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLButtonElement) Autofocus() (autofocus bool) {
	js.Rewrite("$<.autofocus")
	return autofocus
}

// Autofocus prop Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLButtonElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.autofocus = $1", autofocus)
}

// Disabled prop
func (*HTMLButtonElement) Disabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

// Disabled prop
func (*HTMLButtonElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

// Form prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLButtonElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

// FormAction prop Overrides the action attribute (where the data on a form is sent) on the parent form element.
func (*HTMLButtonElement) FormAction() (formAction string) {
	js.Rewrite("$<.formAction")
	return formAction
}

// FormAction prop Overrides the action attribute (where the data on a form is sent) on the parent form element.
func (*HTMLButtonElement) SetFormAction(formAction string) {
	js.Rewrite("$<.formAction = $1", formAction)
}

// FormEnctype prop Used to override the encoding (formEnctype attribute) specified on the form element.
func (*HTMLButtonElement) FormEnctype() (formEnctype string) {
	js.Rewrite("$<.formEnctype")
	return formEnctype
}

// FormEnctype prop Used to override the encoding (formEnctype attribute) specified on the form element.
func (*HTMLButtonElement) SetFormEnctype(formEnctype string) {
	js.Rewrite("$<.formEnctype = $1", formEnctype)
}

// FormMethod prop Overrides the submit method attribute previously specified on a form element.
func (*HTMLButtonElement) FormMethod() (formMethod string) {
	js.Rewrite("$<.formMethod")
	return formMethod
}

// FormMethod prop Overrides the submit method attribute previously specified on a form element.
func (*HTMLButtonElement) SetFormMethod(formMethod string) {
	js.Rewrite("$<.formMethod = $1", formMethod)
}

// FormNoValidate prop Overrides any validation or required attributes on a form or form elements to allow it to be submitted without validation. This can be used to create a "save draft"-type submit option.
func (*HTMLButtonElement) FormNoValidate() (formNoValidate string) {
	js.Rewrite("$<.formNoValidate")
	return formNoValidate
}

// FormNoValidate prop Overrides any validation or required attributes on a form or form elements to allow it to be submitted without validation. This can be used to create a "save draft"-type submit option.
func (*HTMLButtonElement) SetFormNoValidate(formNoValidate string) {
	js.Rewrite("$<.formNoValidate = $1", formNoValidate)
}

// FormTarget prop Overrides the target attribute on a form element.
func (*HTMLButtonElement) FormTarget() (formTarget string) {
	js.Rewrite("$<.formTarget")
	return formTarget
}

// FormTarget prop Overrides the target attribute on a form element.
func (*HTMLButtonElement) SetFormTarget(formTarget string) {
	js.Rewrite("$<.formTarget = $1", formTarget)
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLButtonElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLButtonElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// Status prop
func (*HTMLButtonElement) Status() (status interface{}) {
	js.Rewrite("$<.status")
	return status
}

// Status prop
func (*HTMLButtonElement) SetStatus(status interface{}) {
	js.Rewrite("$<.status = $1", status)
}

// Type prop Gets the classification and default behavior of the button.
func (*HTMLButtonElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop Gets the classification and default behavior of the button.
func (*HTMLButtonElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

// ValidationMessage prop Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
func (*HTMLButtonElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

// Validity prop Returns a  ValidityState object that represents the validity states of an element.
func (*HTMLButtonElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

// Value prop Sets or retrieves the default or selected value of the control.
func (*HTMLButtonElement) Value() (value string) {
	js.Rewrite("$<.value")
	return value
}

// Value prop Sets or retrieves the default or selected value of the control.
func (*HTMLButtonElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

// WillValidate prop Returns whether an element will successfully validate based on forms validation rules and constraints.
func (*HTMLButtonElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}
