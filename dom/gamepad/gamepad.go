package gamepad

import (
	"github.com/matthewmueller/golly/dom/gamepadbutton"
	"github.com/matthewmueller/golly/js"
)

// Gamepad struct
// js:"Gamepad,omit"
type Gamepad struct {
}

// Axes prop
func (*Gamepad) Axes() (axes []float32) {
	js.Rewrite("$<.axes")
	return axes
}

// Buttons prop
func (*Gamepad) Buttons() (buttons []*gamepadbutton.GamepadButton) {
	js.Rewrite("$<.buttons")
	return buttons
}

// Connected prop
func (*Gamepad) Connected() (connected bool) {
	js.Rewrite("$<.connected")
	return connected
}

// ID prop
func (*Gamepad) ID() (id string) {
	js.Rewrite("$<.id")
	return id
}

// Index prop
func (*Gamepad) Index() (index int) {
	js.Rewrite("$<.index")
	return index
}

// Mapping prop
func (*Gamepad) Mapping() (mapping string) {
	js.Rewrite("$<.mapping")
	return mapping
}

// Timestamp prop
func (*Gamepad) Timestamp() (timestamp int) {
	js.Rewrite("$<.timestamp")
	return timestamp
}
