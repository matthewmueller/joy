package deviceacceleration

import "github.com/matthewmueller/golly/js"

type DeviceAcceleration struct {
}

func (*DeviceAcceleration) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*DeviceAcceleration) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*DeviceAcceleration) GetZ() (z float32) {
	js.Rewrite("$<.z")
	return z
}
