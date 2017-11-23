package mediadeviceinfo

import (
	"github.com/matthewmueller/golly/dom2/mediadevicekind"
	"github.com/matthewmueller/golly/js"
)

// MediaDeviceInfo struct
// js:"MediaDeviceInfo,omit"
type MediaDeviceInfo struct {
}

// DeviceID prop
func (*MediaDeviceInfo) DeviceID() (deviceId string) {
	js.Rewrite("$<.deviceId")
	return deviceId
}

// GroupID prop
func (*MediaDeviceInfo) GroupID() (groupId string) {
	js.Rewrite("$<.groupId")
	return groupId
}

// Kind prop
func (*MediaDeviceInfo) Kind() (kind *mediadevicekind.MediaDeviceKind) {
	js.Rewrite("$<.kind")
	return kind
}

// Label prop
func (*MediaDeviceInfo) Label() (label string) {
	js.Rewrite("$<.label")
	return label
}
