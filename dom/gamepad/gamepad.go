package gamepad

import (
	"github.com/matthewmueller/joy/dom/gamepadbutton"
	"github.com/matthewmueller/joy/macro"
)

// Gamepad struct
// js:"Gamepad,omit"
type Gamepad struct {
}

// Axes prop
// js:"axes"
func (*Gamepad) Axes() (axes []float32) {
	macro.Rewrite("$_.axes")
	return axes
}

// Buttons prop
// js:"buttons"
func (*Gamepad) Buttons() (buttons []*gamepadbutton.GamepadButton) {
	macro.Rewrite("$_.buttons")
	return buttons
}

// Connected prop
// js:"connected"
func (*Gamepad) Connected() (connected bool) {
	macro.Rewrite("$_.connected")
	return connected
}

// ID prop
// js:"id"
func (*Gamepad) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// Index prop
// js:"index"
func (*Gamepad) Index() (index int) {
	macro.Rewrite("$_.index")
	return index
}

// Mapping prop
// js:"mapping"
func (*Gamepad) Mapping() (mapping string) {
	macro.Rewrite("$_.mapping")
	return mapping
}

// Timestamp prop
// js:"timestamp"
func (*Gamepad) Timestamp() (timestamp int) {
	macro.Rewrite("$_.timestamp")
	return timestamp
}
