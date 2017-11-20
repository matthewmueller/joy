package htmlbuttonelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/validitystate"
	"github.com/matthewmueller/golly/js"
)

type HTMLButtonElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLButtonElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLButtonElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLButtonElement) GetAutofocus() (autofocus bool) {
	js.Rewrite("$<.autofocus")
	return autofocus
}

func (*HTMLButtonElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.autofocus = $1", autofocus)
}

func (*HTMLButtonElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLButtonElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLButtonElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLButtonElement) GetFormAction() (formAction string) {
	js.Rewrite("$<.formAction")
	return formAction
}

func (*HTMLButtonElement) SetFormAction(formAction string) {
	js.Rewrite("$<.formAction = $1", formAction)
}

func (*HTMLButtonElement) GetFormEnctype() (formEnctype string) {
	js.Rewrite("$<.formEnctype")
	return formEnctype
}

func (*HTMLButtonElement) SetFormEnctype(formEnctype string) {
	js.Rewrite("$<.formEnctype = $1", formEnctype)
}

func (*HTMLButtonElement) GetFormMethod() (formMethod string) {
	js.Rewrite("$<.formMethod")
	return formMethod
}

func (*HTMLButtonElement) SetFormMethod(formMethod string) {
	js.Rewrite("$<.formMethod = $1", formMethod)
}

func (*HTMLButtonElement) GetFormNoValidate() (formNoValidate string) {
	js.Rewrite("$<.formNoValidate")
	return formNoValidate
}

func (*HTMLButtonElement) SetFormNoValidate(formNoValidate string) {
	js.Rewrite("$<.formNoValidate = $1", formNoValidate)
}

func (*HTMLButtonElement) GetFormTarget() (formTarget string) {
	js.Rewrite("$<.formTarget")
	return formTarget
}

func (*HTMLButtonElement) SetFormTarget(formTarget string) {
	js.Rewrite("$<.formTarget = $1", formTarget)
}

func (*HTMLButtonElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLButtonElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLButtonElement) GetStatus() (status interface{}) {
	js.Rewrite("$<.status")
	return status
}

func (*HTMLButtonElement) SetStatus(status interface{}) {
	js.Rewrite("$<.status = $1", status)
}

func (*HTMLButtonElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLButtonElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLButtonElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLButtonElement) GetValidity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLButtonElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLButtonElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

func (*HTMLButtonElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}
