package deviceorientationeventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type DeviceOrientationEventInit struct {
	*eventinit.EventInit

	absolute *bool
	alpha    *float32
	beta     *float32
	gamma    *float32
}
