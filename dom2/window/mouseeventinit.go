package window

type MouseEventInit struct {
	*EventModifierInit

	button        *int8
	buttons       *uint8
	clientX       *int
	clientY       *int
	relatedTarget *EventTarget
	screenX       *int
	screenY       *int
}
