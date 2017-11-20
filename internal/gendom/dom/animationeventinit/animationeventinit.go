package animationeventinit

type AnimationEventInit struct {
	*EventInit

	animationName *string
	elapsedTime   *float32
}
