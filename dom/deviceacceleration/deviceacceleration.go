package deviceacceleration

import "github.com/matthewmueller/joy/js"

// DeviceAcceleration struct
// js:"DeviceAcceleration,omit"
type DeviceAcceleration struct {
}

// X prop
// js:"x"
func (*DeviceAcceleration) X() (x float32) {
	js.Rewrite("$_.x")
	return x
}

// Y prop
// js:"y"
func (*DeviceAcceleration) Y() (y float32) {
	js.Rewrite("$_.y")
	return y
}

// Z prop
// js:"z"
func (*DeviceAcceleration) Z() (z float32) {
	js.Rewrite("$_.z")
	return z
}
