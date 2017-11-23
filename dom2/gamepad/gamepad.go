package gamepad

import (
	"github.com/matthewmueller/golly/dom2/gamepadbutton"
	"github.com/matthewmueller/golly/js"
)

// Gamepad struct
// js:"Gamepad,omit"
type Gamepad struct {
}

// Axes
func (*Gamepad) Axes() (axes []float32) {
	js.Rewrite("$<.Axes")
	return axes
}

// Buttons
func (*Gamepad) Buttons() (buttons []*gamepadbutton.GamepadButton) {
	js.Rewrite("$<.Buttons")
	return buttons
}

// Connected
func (*Gamepad) Connected() (connected bool) {
	js.Rewrite("$<.Connected")
	return connected
}

// ID
func (*Gamepad) ID() (id string) {
	js.Rewrite("$<.ID")
	return id
}

// Index
func (*Gamepad) Index() (index int) {
	js.Rewrite("$<.Index")
	return index
}

// Mapping
func (*Gamepad) Mapping() (mapping string) {
	js.Rewrite("$<.Mapping")
	return mapping
}

// Timestamp
func (*Gamepad) Timestamp() (timestamp int) {
	js.Rewrite("$<.Timestamp")
	return timestamp
}
