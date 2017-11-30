package gamepadbutton

import "github.com/matthewmueller/golly/js"

// GamepadButton struct
// js:"GamepadButton,omit"
type GamepadButton struct {
}

// Pressed prop
// js:"pressed"
func (*GamepadButton) Pressed() (pressed bool) {
	js.Rewrite("$_.pressed")
	return pressed
}

// Value prop
// js:"value"
func (*GamepadButton) Value() (value float32) {
	js.Rewrite("$_.value")
	return value
}
