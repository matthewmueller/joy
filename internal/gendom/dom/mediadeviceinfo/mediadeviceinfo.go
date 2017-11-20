package mediadeviceinfo

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/mediadevicekind"
	"github.com/matthewmueller/golly/js"
)

type MediaDeviceInfo struct {
}

func (*MediaDeviceInfo) GetDeviceID() (deviceId string) {
	js.Rewrite("$<.deviceId")
	return deviceId
}

func (*MediaDeviceInfo) GetGroupID() (groupId string) {
	js.Rewrite("$<.groupId")
	return groupId
}

func (*MediaDeviceInfo) GetKind() (kind *mediadevicekind.MediaDeviceKind) {
	js.Rewrite("$<.kind")
	return kind
}

func (*MediaDeviceInfo) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}
