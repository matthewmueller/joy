package deviceacceleration

import "github.com/matthewmueller/joy/macro"

// DeviceAcceleration struct
// js:"DeviceAcceleration,omit"
type DeviceAcceleration struct {
}

// X prop
// js:"x"
func (*DeviceAcceleration) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// Y prop
// js:"y"
func (*DeviceAcceleration) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// Z prop
// js:"z"
func (*DeviceAcceleration) Z() (z float32) {
	macro.Rewrite("$_.z")
	return z
}
