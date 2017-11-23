package gamepadevent

import (
	"github.com/matthewmueller/golly/dom2/gamepad"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"GamepadEvent,omit"
type GamepadEvent struct {
	window.Event
}

// Gamepad
func (*GamepadEvent) Gamepad() (gamepad *gamepad.Gamepad) {
	js.Rewrite("$<.Gamepad")
	return gamepad
}
