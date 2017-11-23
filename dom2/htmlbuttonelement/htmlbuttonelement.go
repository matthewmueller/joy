package htmlbuttonelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/validitystate"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLButtonElement,omit"
type HTMLButtonElement struct {
	window.HTMLElement
}

// CheckValidity Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLButtonElement) CheckValidity() (b bool) {
	js.Rewrite("$<.CheckValidity()")
	return b
}

// SetCustomValidity Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
func (*HTMLButtonElement) SetCustomValidity(err string) {
	js.Rewrite("$<.SetCustomValidity($1)", err)
}

// Autofocus Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLButtonElement) Autofocus() (autofocus bool) {
	js.Rewrite("$<.Autofocus")
	return autofocus
}

// Autofocus Provides a way to direct a user to a specific field when a document loads. This can provide both direction and convenience for a user, reducing the need to click or tab to a field when a page opens. This attribute is true when present on an element, and false when missing.
func (*HTMLButtonElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.Autofocus = $1", autofocus)
}

// Disabled
func (*HTMLButtonElement) Disabled() (disabled bool) {
	js.Rewrite("$<.Disabled")
	return disabled
}

// Disabled
func (*HTMLButtonElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.Disabled = $1", disabled)
}

// Form Retrieves a reference to the form that the object is embedded in.
func (*HTMLButtonElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}

// FormAction Overrides the action attribute (where the data on a form is sent) on the parent form element.
func (*HTMLButtonElement) FormAction() (formAction string) {
	js.Rewrite("$<.FormAction")
	return formAction
}

// FormAction Overrides the action attribute (where the data on a form is sent) on the parent form element.
func (*HTMLButtonElement) SetFormAction(formAction string) {
	js.Rewrite("$<.FormAction = $1", formAction)
}

// FormEnctype Used to override the encoding (formEnctype attribute) specified on the form element.
func (*HTMLButtonElement) FormEnctype() (formEnctype string) {
	js.Rewrite("$<.FormEnctype")
	return formEnctype
}

// FormEnctype Used to override the encoding (formEnctype attribute) specified on the form element.
func (*HTMLButtonElement) SetFormEnctype(formEnctype string) {
	js.Rewrite("$<.FormEnctype = $1", formEnctype)
}

// FormMethod Overrides the submit method attribute previously specified on a form element.
func (*HTMLButtonElement) FormMethod() (formMethod string) {
	js.Rewrite("$<.FormMethod")
	return formMethod
}

// FormMethod Overrides the submit method attribute previously specified on a form element.
func (*HTMLButtonElement) SetFormMethod(formMethod string) {
	js.Rewrite("$<.FormMethod = $1", formMethod)
}

// FormNoValidate Overrides any validation or required attributes on a form or form elements to allow it to be submitted without validation. This can be used to create a "save draft"-type submit option.
func (*HTMLButtonElement) FormNoValidate() (formNoValidate string) {
	js.Rewrite("$<.FormNoValidate")
	return formNoValidate
}

// FormNoValidate Overrides any validation or required attributes on a form or form elements to allow it to be submitted without validation. This can be used to create a "save draft"-type submit option.
func (*HTMLButtonElement) SetFormNoValidate(formNoValidate string) {
	js.Rewrite("$<.FormNoValidate = $1", formNoValidate)
}

// FormTarget Overrides the target attribute on a form element.
func (*HTMLButtonElement) FormTarget() (formTarget string) {
	js.Rewrite("$<.FormTarget")
	return formTarget
}

// FormTarget Overrides the target attribute on a form element.
func (*HTMLButtonElement) SetFormTarget(formTarget string) {
	js.Rewrite("$<.FormTarget = $1", formTarget)
}

// Name Sets or retrieves the name of the object.
func (*HTMLButtonElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the name of the object.
func (*HTMLButtonElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// Status
func (*HTMLButtonElement) Status() (status interface{}) {
	js.Rewrite("$<.Status")
	return status
}

// Status
func (*HTMLButtonElement) SetStatus(status interface{}) {
	js.Rewrite("$<.Status = $1", status)
}

// Type Gets the classification and default behavior of the button.
func (*HTMLButtonElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type Gets the classification and default behavior of the button.
func (*HTMLButtonElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}

// ValidationMessage Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
func (*HTMLButtonElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.ValidationMessage")
	return validationMessage
}

// Validity Returns a  ValidityState object that represents the validity states of an element.
func (*HTMLButtonElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.Validity")
	return validity
}

// Value Sets or retrieves the default or selected value of the control.
func (*HTMLButtonElement) Value() (value string) {
	js.Rewrite("$<.Value")
	return value
}

// Value Sets or retrieves the default or selected value of the control.
func (*HTMLButtonElement) SetValue(value string) {
	js.Rewrite("$<.Value = $1", value)
}

// WillValidate Returns whether an element will successfully validate based on forms validation rules and constraints.
func (*HTMLButtonElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.WillValidate")
	return willValidate
}
