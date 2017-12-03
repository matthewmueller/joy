package devicemotioneventinit

import (
	"github.com/matthewmueller/joy/dom/deviceaccelerationdict"
	"github.com/matthewmueller/joy/dom/devicerotationratedict"
	"github.com/matthewmueller/joy/dom/eventinit"
)

type DeviceMotionEventInit struct {
	*eventinit.EventInit

	acceleration                 *deviceaccelerationdict.DeviceAccelerationDict
	accelerationIncludingGravity *deviceaccelerationdict.DeviceAccelerationDict
	interval                     *float32
	rotationRate                 *devicerotationratedict.DeviceRotationRateDict
}
