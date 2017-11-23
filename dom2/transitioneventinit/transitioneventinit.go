package transitioneventinit

import "github.com/matthewmueller/golly/dom2/eventinit"

type TransitionEventInit struct {
	*eventinit.EventInit

	elapsedTime  *float32
	propertyName *string
}
