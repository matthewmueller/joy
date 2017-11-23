package gamepadevent

import (
	"github.com/matthewmueller/golly/dom/gamepad"
	"github.com/matthewmueller/golly/dom/gamepadeventinit"
	"github.com/matthewmueller/golly/dom/window"
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
