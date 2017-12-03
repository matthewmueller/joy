package animationeventinit

import "github.com/matthewmueller/joy/dom/eventinit"

type AnimationEventInit struct {
	*eventinit.EventInit

	animationName *string
	elapsedTime   *float32
}
