package transitioneventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type TransitionEventInit struct {
	*eventinit.EventInit

	elapsedTime  *float32
	propertyName *string
}
