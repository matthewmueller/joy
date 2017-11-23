package idb

type WheelEventInit struct {
	*MouseEventInit

	deltaMode *uint
	deltaX    *float32
	deltaY    *float32
	deltaZ    *float32
}
