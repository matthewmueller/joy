package devicerotationrate

import "github.com/matthewmueller/golly/js"

type DeviceRotationRate struct {
}

func (*DeviceRotationRate) GetAlpha() (alpha float32) {
	js.Rewrite("$<.alpha")
	return alpha
}

func (*DeviceRotationRate) GetBeta() (beta float32) {
	js.Rewrite("$<.beta")
	return beta
}

func (*DeviceRotationRate) GetGamma() (gamma float32) {
	js.Rewrite("$<.gamma")
	return gamma
}
