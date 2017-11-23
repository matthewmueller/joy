package window

import (
	"github.com/matthewmueller/golly/dom/mediadeviceinfo"
	"github.com/matthewmueller/golly/dom/mediastreamconstraints"
	"github.com/matthewmueller/golly/dom/mediatracksupportedconstraints"
	"github.com/matthewmueller/golly/js"
)

// MediaDevices struct
// js:"MediaDevices,omit"
type MediaDevices struct {
	EventTarget
}

// EnumerateDevices fn
func (*MediaDevices) EnumerateDevices() (m []*mediadeviceinfo.MediaDeviceInfo) {
	js.Rewrite("await $<.enumerateDevices()")
	return m
}

// GetSupportedConstraints fn
func (*MediaDevices) GetSupportedConstraints() (m *mediatracksupportedconstraints.MediaTrackSupportedConstraints) {
	js.Rewrite("$<.getSupportedConstraints()")
	return m
}

// GetUserMedia fn
func (*MediaDevices) GetUserMedia(constraints *mediastreamconstraints.MediaStreamConstraints) (m *MediaStream) {
	js.Rewrite("await $<.getUserMedia($1)", constraints)
	return m
}

// Ondevicechange prop
func (*MediaDevices) Ondevicechange() (ondevicechange func(Event)) {
	js.Rewrite("$<.ondevicechange")
	return ondevicechange
}

// Ondevicechange prop
func (*MediaDevices) SetOndevicechange(ondevicechange func(Event)) {
	js.Rewrite("$<.ondevicechange = $1", ondevicechange)
}
