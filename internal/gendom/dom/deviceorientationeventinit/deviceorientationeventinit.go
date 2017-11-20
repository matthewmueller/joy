package deviceorientationeventinit

type DeviceOrientationEventInit struct {
	*EventInit

	absolute *bool
	alpha    *float32
	beta     *float32
	gamma    *float32
}
