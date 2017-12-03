package gamepadbutton

import "github.com/matthewmueller/joy/macro"

// GamepadButton struct
// js:"GamepadButton,omit"
type GamepadButton struct {
}

// Pressed prop
// js:"pressed"
func (*GamepadButton) Pressed() (pressed bool) {
	macro.Rewrite("$_.pressed")
	return pressed
}

// Value prop
// js:"value"
func (*GamepadButton) Value() (value float32) {
	macro.Rewrite("$_.value")
	return value
}
