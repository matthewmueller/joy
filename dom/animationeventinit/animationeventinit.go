package animationeventinit

import "github.com/matthewmueller/golly/dom2/eventinit"

type AnimationEventInit struct {
	*eventinit.EventInit

	animationName *string
	elapsedTime   *float32
}
