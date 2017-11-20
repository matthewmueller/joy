package deviceorientationevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type DeviceOrientationEvent struct {
	*event.Event
}

func (*DeviceOrientationEvent) InitDeviceOrientationEvent(kind string, bubbles bool, cancelable bool, alpha float32, beta float32, gamma float32, absolute bool) {
	js.Rewrite("$<.initDeviceOrientationEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, alpha, beta, gamma, absolute)
}

func (*DeviceOrientationEvent) GetAbsolute() (absolute bool) {
	js.Rewrite("$<.absolute")
	return absolute
}

func (*DeviceOrientationEvent) GetAlpha() (alpha float32) {
	js.Rewrite("$<.alpha")
	return alpha
}

func (*DeviceOrientationEvent) GetBeta() (beta float32) {
	js.Rewrite("$<.beta")
	return beta
}

func (*DeviceOrientationEvent) GetGamma() (gamma float32) {
	js.Rewrite("$<.gamma")
	return gamma
}
