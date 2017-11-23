package gamepadbutton

import "github.com/matthewmueller/golly/js"

// GamepadButton struct
// js:"GamepadButton,omit"
type GamepadButton struct {
}

// Pressed prop
func (*GamepadButton) Pressed() (pressed bool) {
	js.Rewrite("$<.pressed")
	return pressed
}

// Value prop
func (*GamepadButton) Value() (value float32) {
	js.Rewrite("$<.value")
	return value
}
