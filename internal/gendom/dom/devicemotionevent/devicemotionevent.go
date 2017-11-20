package devicemotionevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/deviceacceleration"
	"github.com/matthewmueller/golly/internal/gendom/dom/deviceaccelerationdict"
	"github.com/matthewmueller/golly/internal/gendom/dom/devicerotationrate"
	"github.com/matthewmueller/golly/internal/gendom/dom/devicerotationratedict"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type DeviceMotionEvent struct {
	*event.Event
}

func (*DeviceMotionEvent) InitDeviceMotionEvent(kind string, bubbles bool, cancelable bool, acceleration *deviceaccelerationdict.DeviceAccelerationDict, accelerationIncludingGravity *deviceaccelerationdict.DeviceAccelerationDict, rotationRate *devicerotationratedict.DeviceRotationRateDict, interval float32) {
	js.Rewrite("$<.initDeviceMotionEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, acceleration, accelerationIncludingGravity, rotationRate, interval)
}

func (*DeviceMotionEvent) GetAcceleration() (acceleration *deviceacceleration.DeviceAcceleration) {
	js.Rewrite("$<.acceleration")
	return acceleration
}

func (*DeviceMotionEvent) GetAccelerationIncludingGravity() (accelerationIncludingGravity *deviceacceleration.DeviceAcceleration) {
	js.Rewrite("$<.accelerationIncludingGravity")
	return accelerationIncludingGravity
}

func (*DeviceMotionEvent) GetInterval() (interval float32) {
	js.Rewrite("$<.interval")
	return interval
}

func (*DeviceMotionEvent) GetRotationRate() (rotationRate *devicerotationrate.DeviceRotationRate) {
	js.Rewrite("$<.rotationRate")
	return rotationRate
}
