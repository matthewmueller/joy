package animationeventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type AnimationEventInit struct {
	*eventinit.EventInit

	animationName *string
	elapsedTime   *float32
}
