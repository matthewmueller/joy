package transitioneventinit

type TransitionEventInit struct {
	*eventinit.EventInit

	elapsedTime  *float32
	propertyName *string
}
