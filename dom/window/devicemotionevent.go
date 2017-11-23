package window

import (
	"github.com/matthewmueller/golly/dom/deviceacceleration"
	"github.com/matthewmueller/golly/dom/deviceaccelerationdict"
	"github.com/matthewmueller/golly/dom/devicemotioneventinit"
	"github.com/matthewmueller/golly/dom/devicerotationrate"
	"github.com/matthewmueller/golly/dom/devicerotationratedict"
	"github.com/matthewmueller/golly/js"
)

// NewDeviceMotionEvent fn
func NewDeviceMotionEvent(typearg string, eventinitdict *devicemotioneventinit.DeviceMotionEventInit) *DeviceMotionEvent {
	js.Rewrite("DeviceMotionEvent")
	return &DeviceMotionEvent{}
}

// DeviceMotionEvent struct
// js:"DeviceMotionEvent,omit"
type DeviceMotionEvent struct {
	Event
}

// InitDeviceMotionEvent fn
func (*DeviceMotionEvent) InitDeviceMotionEvent(kind string, bubbles bool, cancelable bool, acceleration *deviceaccelerationdict.DeviceAccelerationDict, accelerationIncludingGravity *deviceaccelerationdict.DeviceAccelerationDict, rotationRate *devicerotationratedict.DeviceRotationRateDict, interval float32) {
	js.Rewrite("$<.initDeviceMotionEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, acceleration, accelerationIncludingGravity, rotationRate, interval)
}

// Acceleration prop
func (*DeviceMotionEvent) Acceleration() (acceleration *deviceacceleration.DeviceAcceleration) {
	js.Rewrite("$<.acceleration")
	return acceleration
}

// AccelerationIncludingGravity prop
func (*DeviceMotionEvent) AccelerationIncludingGravity() (accelerationIncludingGravity *deviceacceleration.DeviceAcceleration) {
	js.Rewrite("$<.accelerationIncludingGravity")
	return accelerationIncludingGravity
}

// Interval prop
func (*DeviceMotionEvent) Interval() (interval float32) {
	js.Rewrite("$<.interval")
	return interval
}

// RotationRate prop
func (*DeviceMotionEvent) RotationRate() (rotationRate *devicerotationrate.DeviceRotationRate) {
	js.Rewrite("$<.rotationRate")
	return rotationRate
}
