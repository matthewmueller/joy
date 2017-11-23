package gamepadbutton

import "github.com/matthewmueller/golly/js"

// js:"GamepadButton,omit"
type GamepadButton struct {
}

// Pressed
func (*GamepadButton) Pressed() (pressed bool) {
	js.Rewrite("$<.Pressed")
	return pressed
}

// Value
func (*GamepadButton) Value() (value float32) {
	js.Rewrite("$<.Value")
	return value
}
