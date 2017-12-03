package gamepadeventinit

import (
	"github.com/matthewmueller/joy/dom/eventinit"
	"github.com/matthewmueller/joy/dom/gamepad"
)

type GamepadEventInit struct {
	*eventinit.EventInit

	gamepad *gamepad.Gamepad
}
