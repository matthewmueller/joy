package gamepadevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/gamepad"
	"github.com/matthewmueller/golly/js"
)

type GamepadEvent struct {
	*event.Event
}

func (*GamepadEvent) GetGamepad() (gamepad *gamepad.Gamepad) {
	js.Rewrite("$<.gamepad")
	return gamepad
}
