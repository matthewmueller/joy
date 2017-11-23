package gamepadevent

import (
	"github.com/matthewmueller/golly/dom2/gamepad"
	"github.com/matthewmueller/golly/dom2/gamepadeventinit"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(typearg string, eventinitdict *gamepadeventinit.GamepadEventInit) *GamepadEvent {
	js.Rewrite("GamepadEvent")
	return &GamepadEvent{}
}

// GamepadEvent struct
// js:"GamepadEvent,omit"
type GamepadEvent struct {
	window.Event
}

// Gamepad prop
func (*GamepadEvent) Gamepad() (gamepad *gamepad.Gamepad) {
	js.Rewrite("$<.gamepad")
	return gamepad
}
