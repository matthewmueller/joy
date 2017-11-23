package idb

import (
	"github.com/matthewmueller/golly/dom2/deviceacceleration"
	"github.com/matthewmueller/golly/dom2/deviceaccelerationdict"
	"github.com/matthewmueller/golly/dom2/devicemotioneventinit"
	"github.com/matthewmueller/golly/dom2/devicerotationrate"
	"github.com/matthewmueller/golly/dom2/devicerotationratedict"
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

// InitDeviceMotionEvent
func (*DeviceMotionEvent) InitDeviceMotionEvent(kind string, bubbles bool, cancelable bool, acceleration *deviceaccelerationdict.DeviceAccelerationDict, accelerationIncludingGravity *deviceaccelerationdict.DeviceAccelerationDict, rotationRate *devicerotationratedict.DeviceRotationRateDict, interval float32) {
	js.Rewrite("$<.InitDeviceMotionEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, acceleration, accelerationIncludingGravity, rotationRate, interval)
}

// Acceleration
func (*DeviceMotionEvent) Acceleration() (acceleration *deviceacceleration.DeviceAcceleration) {
	js.Rewrite("$<.Acceleration")
	return acceleration
}

// AccelerationIncludingGravity
func (*DeviceMotionEvent) AccelerationIncludingGravity() (accelerationIncludingGravity *deviceacceleration.DeviceAcceleration) {
	js.Rewrite("$<.AccelerationIncludingGravity")
	return accelerationIncludingGravity
}

// Interval
func (*DeviceMotionEvent) Interval() (interval float32) {
	js.Rewrite("$<.Interval")
	return interval
}

// RotationRate
func (*DeviceMotionEvent) RotationRate() (rotationRate *devicerotationrate.DeviceRotationRate) {
	js.Rewrite("$<.RotationRate")
	return rotationRate
}
