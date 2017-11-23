package deviceacceleration

import "github.com/matthewmueller/golly/js"

// DeviceAcceleration struct
// js:"DeviceAcceleration,omit"
type DeviceAcceleration struct {
}

// X prop
func (*DeviceAcceleration) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*DeviceAcceleration) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Z prop
func (*DeviceAcceleration) Z() (z float32) {
	js.Rewrite("$<.z")
	return z
}
