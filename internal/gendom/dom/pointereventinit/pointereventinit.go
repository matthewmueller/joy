package pointereventinit

type PointerEventInit struct {
	*MouseEventInit

	height      *int
	isPrimary   *bool
	pointerId   *int
	pointerType *string
	pressure    *float32
	tiltX       *int
	tiltY       *int
	width       *int
}
