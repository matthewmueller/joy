package htmloutputelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/validitystate"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLOutputElement,omit"
type HTMLOutputElement struct {
	window.HTMLElement
}

// CheckValidity
func (*HTMLOutputElement) CheckValidity() (b bool) {
	js.Rewrite("$<.CheckValidity()")
	return b
}

// ReportValidity
func (*HTMLOutputElement) ReportValidity() (b bool) {
	js.Rewrite("$<.ReportValidity()")
	return b
}

// SetCustomValidity
func (*HTMLOutputElement) SetCustomValidity(err string) {
	js.Rewrite("$<.SetCustomValidity($1)", err)
}

// DefaultValue
func (*HTMLOutputElement) DefaultValue() (defaultValue string) {
	js.Rewrite("$<.DefaultValue")
	return defaultValue
}

// DefaultValue
func (*HTMLOutputElement) SetDefaultValue(defaultValue string) {
	js.Rewrite("$<.DefaultValue = $1", defaultValue)
}

// Form
func (*HTMLOutputElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}

// HTMLFor
func (*HTMLOutputElement) HTMLFor() (htmlFor *domsettabletokenlist.DOMSettableTokenList) {
	js.Rewrite("$<.HTMLFor")
	return htmlFor
}

// Name
func (*HTMLOutputElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name
func (*HTMLOutputElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// Type
func (*HTMLOutputElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// ValidationMessage
func (*HTMLOutputElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.ValidationMessage")
	return validationMessage
}

// Validity
func (*HTMLOutputElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.Validity")
	return validity
}

// Value
func (*HTMLOutputElement) Value() (value string) {
	js.Rewrite("$<.Value")
	return value
}

// Value
func (*HTMLOutputElement) SetValue(value string) {
	js.Rewrite("$<.Value = $1", value)
}

// WillValidate
func (*HTMLOutputElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.WillValidate")
	return willValidate
}
