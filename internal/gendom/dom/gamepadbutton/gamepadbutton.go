package gamepadbutton

import "github.com/matthewmueller/golly/js"

type GamepadButton struct {
}

func (*GamepadButton) GetPressed() (pressed bool) {
	js.Rewrite("$<.pressed")
	return pressed
}

func (*GamepadButton) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}
