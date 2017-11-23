package devicerotationrate

import "github.com/matthewmueller/golly/js"

// DeviceRotationRate struct
// js:"DeviceRotationRate,omit"
type DeviceRotationRate struct {
}

// Alpha
func (*DeviceRotationRate) Alpha() (alpha float32) {
	js.Rewrite("$<.Alpha")
	return alpha
}

// Beta
func (*DeviceRotationRate) Beta() (beta float32) {
	js.Rewrite("$<.Beta")
	return beta
}

// Gamma
func (*DeviceRotationRate) Gamma() (gamma float32) {
	js.Rewrite("$<.Gamma")
	return gamma
}
