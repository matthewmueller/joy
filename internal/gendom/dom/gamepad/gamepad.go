package gamepad

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/gamepadbutton"
	"github.com/matthewmueller/golly/js"
)

type Gamepad struct {
}

func (*Gamepad) GetAxes() (axes []float32) {
	js.Rewrite("$<.axes")
	return axes
}

func (*Gamepad) GetButtons() (buttons []*gamepadbutton.GamepadButton) {
	js.Rewrite("$<.buttons")
	return buttons
}

func (*Gamepad) GetConnected() (connected bool) {
	js.Rewrite("$<.connected")
	return connected
}

func (*Gamepad) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*Gamepad) GetIndex() (index int) {
	js.Rewrite("$<.index")
	return index
}

func (*Gamepad) GetMapping() (mapping string) {
	js.Rewrite("$<.mapping")
	return mapping
}

func (*Gamepad) GetTimestamp() (timestamp int) {
	js.Rewrite("$<.timestamp")
	return timestamp
}
