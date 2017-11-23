package htmlprogresselement

import (
	"github.com/matthewmueller/golly/dom/htmlformelement"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLProgressElement struct
// js:"HTMLProgressElement,omit"
type HTMLProgressElement struct {
	window.HTMLElement
}

// Form prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLProgressElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

// Max prop Defines the maximum, or "done" value for a progress element.
func (*HTMLProgressElement) Max() (max float32) {
	js.Rewrite("$<.max")
	return max
}

// Max prop Defines the maximum, or "done" value for a progress element.
func (*HTMLProgressElement) SetMax(max float32) {
	js.Rewrite("$<.max = $1", max)
}

// Position prop Returns the quotient of value/max when the value attribute is set (determinate progress bar), or -1 when the value attribute is missing (indeterminate progress bar).
func (*HTMLProgressElement) Position() (position float32) {
	js.Rewrite("$<.position")
	return position
}

// Value prop Sets or gets the current value of a progress element. The value must be a non-negative number between 0 and the max value.
func (*HTMLProgressElement) Value() (value float32) {
	js.Rewrite("$<.value")
	return value
}

// Value prop Sets or gets the current value of a progress element. The value must be a non-negative number between 0 and the max value.
func (*HTMLProgressElement) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}
