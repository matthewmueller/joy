package devicerotationrate

import "github.com/matthewmueller/golly/js"

// DeviceRotationRate struct
// js:"DeviceRotationRate,omit"
type DeviceRotationRate struct {
}

// Alpha prop
func (*DeviceRotationRate) Alpha() (alpha float32) {
	js.Rewrite("$<.alpha")
	return alpha
}

// Beta prop
func (*DeviceRotationRate) Beta() (beta float32) {
	js.Rewrite("$<.beta")
	return beta
}

// Gamma prop
func (*DeviceRotationRate) Gamma() (gamma float32) {
	js.Rewrite("$<.gamma")
	return gamma
}
