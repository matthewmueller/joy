package mediadeviceinfo

import (
	"github.com/matthewmueller/golly/dom/mediadevicekind"
	"github.com/matthewmueller/golly/js"
)

// MediaDeviceInfo struct
// js:"MediaDeviceInfo,omit"
type MediaDeviceInfo struct {
}

// DeviceID prop
// js:"deviceId"
func (*MediaDeviceInfo) DeviceID() (deviceId string) {
	js.Rewrite("$_.deviceId")
	return deviceId
}

// GroupID prop
// js:"groupId"
func (*MediaDeviceInfo) GroupID() (groupId string) {
	js.Rewrite("$_.groupId")
	return groupId
}

// Kind prop
// js:"kind"
func (*MediaDeviceInfo) Kind() (kind *mediadevicekind.MediaDeviceKind) {
	js.Rewrite("$_.kind")
	return kind
}

// Label prop
// js:"label"
func (*MediaDeviceInfo) Label() (label string) {
	js.Rewrite("$_.label")
	return label
}
