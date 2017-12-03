package mediadeviceinfo

import (
	"github.com/matthewmueller/joy/dom/mediadevicekind"
	"github.com/matthewmueller/joy/macro"
)

// MediaDeviceInfo struct
// js:"MediaDeviceInfo,omit"
type MediaDeviceInfo struct {
}

// DeviceID prop
// js:"deviceId"
func (*MediaDeviceInfo) DeviceID() (deviceId string) {
	macro.Rewrite("$_.deviceId")
	return deviceId
}

// GroupID prop
// js:"groupId"
func (*MediaDeviceInfo) GroupID() (groupId string) {
	macro.Rewrite("$_.groupId")
	return groupId
}

// Kind prop
// js:"kind"
func (*MediaDeviceInfo) Kind() (kind *mediadevicekind.MediaDeviceKind) {
	macro.Rewrite("$_.kind")
	return kind
}

// Label prop
// js:"label"
func (*MediaDeviceInfo) Label() (label string) {
	macro.Rewrite("$_.label")
	return label
}
