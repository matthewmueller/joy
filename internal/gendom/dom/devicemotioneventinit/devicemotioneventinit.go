package devicemotioneventinit

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/deviceaccelerationdict"
	"github.com/matthewmueller/golly/internal/gendom/dom/devicerotationratedict"
)

type DeviceMotionEventInit struct {
	*EventInit

	acceleration                 *deviceaccelerationdict.DeviceAccelerationDict
	accelerationIncludingGravity *deviceaccelerationdict.DeviceAccelerationDict
	interval                     *float32
	rotationRate                 *devicerotationratedict.DeviceRotationRateDict
}
