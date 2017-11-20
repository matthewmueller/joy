package gamepadeventinit

import "github.com/matthewmueller/golly/internal/gendom/dom/gamepad"

type GamepadEventInit struct {
	*EventInit

	gamepad *gamepad.Gamepad
}
