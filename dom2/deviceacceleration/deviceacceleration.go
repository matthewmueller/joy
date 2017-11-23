package deviceacceleration

import "github.com/matthewmueller/golly/js"

// js:"DeviceAcceleration,omit"
type DeviceAcceleration struct {
}

// X
func (*DeviceAcceleration) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*DeviceAcceleration) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Z
func (*DeviceAcceleration) Z() (z float32) {
	js.Rewrite("$<.Z")
	return z
}
