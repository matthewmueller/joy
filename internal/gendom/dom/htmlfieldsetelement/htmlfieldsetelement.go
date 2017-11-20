package htmlfieldsetelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/validitystate"
	"github.com/matthewmueller/golly/js"
)

type HTMLFieldSetElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLFieldSetElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLFieldSetElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLFieldSetElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLFieldSetElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLFieldSetElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLFieldSetElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLFieldSetElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLFieldSetElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLFieldSetElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLFieldSetElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLFieldSetElement) GetValidity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLFieldSetElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}
