package mediadeviceinfo

import (
	"github.com/matthewmueller/golly/dom2/mediadevicekind"
	"github.com/matthewmueller/golly/js"
)

// js:"MediaDeviceInfo,omit"
type MediaDeviceInfo struct {
}

// DeviceID
func (*MediaDeviceInfo) DeviceID() (deviceId string) {
	js.Rewrite("$<.DeviceID")
	return deviceId
}

// GroupID
func (*MediaDeviceInfo) GroupID() (groupId string) {
	js.Rewrite("$<.GroupID")
	return groupId
}

// Kind
func (*MediaDeviceInfo) Kind() (kind *mediadevicekind.MediaDeviceKind) {
	js.Rewrite("$<.Kind")
	return kind
}

// Label
func (*MediaDeviceInfo) Label() (label string) {
	js.Rewrite("$<.Label")
	return label
}
