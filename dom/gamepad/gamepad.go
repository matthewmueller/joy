package gamepad

import (
	"github.com/matthewmueller/joy/dom/gamepadbutton"
	"github.com/matthewmueller/joy/js"
)

// Gamepad struct
// js:"Gamepad,omit"
type Gamepad struct {
}

// Axes prop
// js:"axes"
func (*Gamepad) Axes() (axes []float32) {
	js.Rewrite("$_.axes")
	return axes
}

// Buttons prop
// js:"buttons"
func (*Gamepad) Buttons() (buttons []*gamepadbutton.GamepadButton) {
	js.Rewrite("$_.buttons")
	return buttons
}

// Connected prop
// js:"connected"
func (*Gamepad) Connected() (connected bool) {
	js.Rewrite("$_.connected")
	return connected
}

// ID prop
// js:"id"
func (*Gamepad) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// Index prop
// js:"index"
func (*Gamepad) Index() (index int) {
	js.Rewrite("$_.index")
	return index
}

// Mapping prop
// js:"mapping"
func (*Gamepad) Mapping() (mapping string) {
	js.Rewrite("$_.mapping")
	return mapping
}

// Timestamp prop
// js:"timestamp"
func (*Gamepad) Timestamp() (timestamp int) {
	js.Rewrite("$_.timestamp")
	return timestamp
}
