package htmloutputelement

import (
	"github.com/matthewmueller/golly/dom/domsettabletokenlist"
	"github.com/matthewmueller/golly/dom/htmlformelement"
	"github.com/matthewmueller/golly/dom/validitystate"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLOutputElement struct
// js:"HTMLOutputElement,omit"
type HTMLOutputElement struct {
	window.HTMLElement
}

// CheckValidity fn
func (*HTMLOutputElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

// ReportValidity fn
func (*HTMLOutputElement) ReportValidity() (b bool) {
	js.Rewrite("$<.reportValidity()")
	return b
}

// SetCustomValidity fn
func (*HTMLOutputElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

// DefaultValue prop
func (*HTMLOutputElement) DefaultValue() (defaultValue string) {
	js.Rewrite("$<.defaultValue")
	return defaultValue
}

// DefaultValue prop
func (*HTMLOutputElement) SetDefaultValue(defaultValue string) {
	js.Rewrite("$<.defaultValue = $1", defaultValue)
}

// Form prop
func (*HTMLOutputElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

// HTMLFor prop
func (*HTMLOutputElement) HTMLFor() (htmlFor *domsettabletokenlist.DOMSettableTokenList) {
	js.Rewrite("$<.htmlFor")
	return htmlFor
}

// Name prop
func (*HTMLOutputElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop
func (*HTMLOutputElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// Type prop
func (*HTMLOutputElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// ValidationMessage prop
func (*HTMLOutputElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

// Validity prop
func (*HTMLOutputElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

// Value prop
func (*HTMLOutputElement) Value() (value string) {
	js.Rewrite("$<.value")
	return value
}

// Value prop
func (*HTMLOutputElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

// WillValidate prop
func (*HTMLOutputElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}
