package devicerotationrate

import "github.com/matthewmueller/golly/js"

// DeviceRotationRate struct
// js:"DeviceRotationRate,omit"
type DeviceRotationRate struct {
}

// Alpha prop
// js:"alpha"
func (*DeviceRotationRate) Alpha() (alpha float32) {
	js.Rewrite("$_.alpha")
	return alpha
}

// Beta prop
// js:"beta"
func (*DeviceRotationRate) Beta() (beta float32) {
	js.Rewrite("$_.beta")
	return beta
}

// Gamma prop
// js:"gamma"
func (*DeviceRotationRate) Gamma() (gamma float32) {
	js.Rewrite("$_.gamma")
	return gamma
}
