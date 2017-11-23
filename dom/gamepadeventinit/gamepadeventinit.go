package gamepadeventinit

import (
	"github.com/matthewmueller/golly/dom/eventinit"
	"github.com/matthewmueller/golly/dom/gamepad"
)

type GamepadEventInit struct {
	*eventinit.EventInit

	gamepad *gamepad.Gamepad
}
