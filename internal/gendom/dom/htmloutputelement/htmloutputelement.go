package htmloutputelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domsettabletokenlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/validitystate"
	"github.com/matthewmueller/golly/js"
)

type HTMLOutputElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLOutputElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLOutputElement) ReportValidity() (b bool) {
	js.Rewrite("$<.reportValidity()")
	return b
}

func (*HTMLOutputElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLOutputElement) GetDefaultValue() (defaultValue string) {
	js.Rewrite("$<.defaultValue")
	return defaultValue
}

func (*HTMLOutputElement) SetDefaultValue(defaultValue string) {
	js.Rewrite("$<.defaultValue = $1", defaultValue)
}

func (*HTMLOutputElement) GetForm() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLOutputElement) GetHTMLFor() (htmlFor *domsettabletokenlist.DOMSettableTokenList) {
	js.Rewrite("$<.htmlFor")
	return htmlFor
}

func (*HTMLOutputElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLOutputElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLOutputElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLOutputElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLOutputElement) GetValidity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLOutputElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLOutputElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

func (*HTMLOutputElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}
