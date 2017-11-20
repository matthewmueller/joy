package mediadevices

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediadeviceinfo"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastream"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamconstraints"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediatracksupportedconstraints"
	"github.com/matthewmueller/golly/js"
)

type MediaDevices struct {
	*eventtarget.EventTarget
}

func (*MediaDevices) EnumerateDevices() (m []*mediadeviceinfo.MediaDeviceInfo) {
	js.Rewrite("await $<.enumerateDevices()")
	return m
}

func (*MediaDevices) GetSupportedConstraints() (m *mediatracksupportedconstraints.MediaTrackSupportedConstraints) {
	js.Rewrite("$<.getSupportedConstraints()")
	return m
}

func (*MediaDevices) GetUserMedia(constraints *mediastreamconstraints.MediaStreamConstraints) (m *mediastream.MediaStream) {
	js.Rewrite("await $<.getUserMedia($1)", constraints)
	return m
}

func (*MediaDevices) GetOndevicechange() (devicechange *event.Event) {
	js.Rewrite("$<.ondevicechange")
	return devicechange
}

func (*MediaDevices) SetOndevicechange(devicechange *event.Event) {
	js.Rewrite("$<.ondevicechange = $1", devicechange)
}
