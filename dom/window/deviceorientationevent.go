package window

import (
	"github.com/matthewmueller/golly/dom/deviceorientationeventinit"
	"github.com/matthewmueller/golly/js"
)

// NewDeviceOrientationEvent fn
func NewDeviceOrientationEvent(typearg string, eventinitdict *deviceorientationeventinit.DeviceOrientationEventInit) *DeviceOrientationEvent {
	js.Rewrite("DeviceOrientationEvent")
	return &DeviceOrientationEvent{}
}

// DeviceOrientationEvent struct
// js:"DeviceOrientationEvent,omit"
type DeviceOrientationEvent struct {
	Event
}

// InitDeviceOrientationEvent fn
func (*DeviceOrientationEvent) InitDeviceOrientationEvent(kind string, bubbles bool, cancelable bool, alpha float32, beta float32, gamma float32, absolute bool) {
	js.Rewrite("$<.initDeviceOrientationEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, alpha, beta, gamma, absolute)
}

// Absolute prop
func (*DeviceOrientationEvent) Absolute() (absolute bool) {
	js.Rewrite("$<.absolute")
	return absolute
}

// Alpha prop
func (*DeviceOrientationEvent) Alpha() (alpha float32) {
	js.Rewrite("$<.alpha")
	return alpha
}

// Beta prop
func (*DeviceOrientationEvent) Beta() (beta float32) {
	js.Rewrite("$<.beta")
	return beta
}

// Gamma prop
func (*DeviceOrientationEvent) Gamma() (gamma float32) {
	js.Rewrite("$<.gamma")
	return gamma
}
