package htmlprogresselement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLProgressElement,omit"
type HTMLProgressElement struct {
	window.HTMLElement
}

// Form Retrieves a reference to the form that the object is embedded in.
func (*HTMLProgressElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}

// Max Defines the maximum, or "done" value for a progress element.
func (*HTMLProgressElement) Max() (max float32) {
	js.Rewrite("$<.Max")
	return max
}

// Max Defines the maximum, or "done" value for a progress element.
func (*HTMLProgressElement) SetMax(max float32) {
	js.Rewrite("$<.Max = $1", max)
}

// Position Returns the quotient of value/max when the value attribute is set (determinate progress bar), or -1 when the value attribute is missing (indeterminate progress bar).
func (*HTMLProgressElement) Position() (position float32) {
	js.Rewrite("$<.Position")
	return position
}

// Value Sets or gets the current value of a progress element. The value must be a non-negative number between 0 and the max value.
func (*HTMLProgressElement) Value() (value float32) {
	js.Rewrite("$<.Value")
	return value
}

// Value Sets or gets the current value of a progress element. The value must be a non-negative number between 0 and the max value.
func (*HTMLProgressElement) SetValue(value float32) {
	js.Rewrite("$<.Value = $1", value)
}
