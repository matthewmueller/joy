package transitioneventinit

type TransitionEventInit struct {
	*EventInit

	elapsedTime  *float32
	propertyName *string
}
