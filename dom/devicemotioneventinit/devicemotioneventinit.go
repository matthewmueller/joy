package devicemotioneventinit

import (
	"github.com/matthewmueller/golly/dom/deviceaccelerationdict"
	"github.com/matthewmueller/golly/dom/devicerotationratedict"
	"github.com/matthewmueller/golly/dom/eventinit"
)

type DeviceMotionEventInit struct {
	*eventinit.EventInit

	acceleration                 *deviceaccelerationdict.DeviceAccelerationDict
	accelerationIncludingGravity *deviceaccelerationdict.DeviceAccelerationDict
	interval                     *float32
	rotationRate                 *devicerotationratedict.DeviceRotationRateDict
}
