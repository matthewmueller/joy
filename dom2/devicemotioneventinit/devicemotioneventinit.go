package devicemotioneventinit

import (
	"github.com/matthewmueller/golly/dom2/deviceaccelerationdict"
	"github.com/matthewmueller/golly/dom2/devicerotationratedict"
	"github.com/matthewmueller/golly/dom2/eventinit"
)

type DeviceMotionEventInit struct {
	*eventinit.EventInit

	acceleration                 *deviceaccelerationdict.DeviceAccelerationDict
	accelerationIncludingGravity *deviceaccelerationdict.DeviceAccelerationDict
	interval                     *float32
	rotationRate                 *devicerotationratedict.DeviceRotationRateDict
}
