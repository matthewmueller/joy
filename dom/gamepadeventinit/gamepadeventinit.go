package gamepadeventinit

import (
	"github.com/matthewmueller/golly/dom2/eventinit"
	"github.com/matthewmueller/golly/dom2/gamepad"
)

type GamepadEventInit struct {
	*eventinit.EventInit

	gamepad *gamepad.Gamepad
}
