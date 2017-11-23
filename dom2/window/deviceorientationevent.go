package window

import (
	"github.com/matthewmueller/golly/dom2/deviceorientationeventinit"
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

// InitDeviceOrientationEvent
func (*DeviceOrientationEvent) InitDeviceOrientationEvent(kind string, bubbles bool, cancelable bool, alpha float32, beta float32, gamma float32, absolute bool) {
	js.Rewrite("$<.InitDeviceOrientationEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, alpha, beta, gamma, absolute)
}

// Absolute
func (*DeviceOrientationEvent) Absolute() (absolute bool) {
	js.Rewrite("$<.Absolute")
	return absolute
}

// Alpha
func (*DeviceOrientationEvent) Alpha() (alpha float32) {
	js.Rewrite("$<.Alpha")
	return alpha
}

// Beta
func (*DeviceOrientationEvent) Beta() (beta float32) {
	js.Rewrite("$<.Beta")
	return beta
}

// Gamma
func (*DeviceOrientationEvent) Gamma() (gamma float32) {
	js.Rewrite("$<.Gamma")
	return gamma
}
