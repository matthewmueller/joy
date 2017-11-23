package window

import (
	"github.com/matthewmueller/golly/dom2/mediadeviceinfo"
	"github.com/matthewmueller/golly/dom2/mediastreamconstraints"
	"github.com/matthewmueller/golly/dom2/mediatracksupportedconstraints"
	"github.com/matthewmueller/golly/js"
)

// MediaDevices struct
// js:"MediaDevices,omit"
type MediaDevices struct {
	EventTarget
}

// EnumerateDevices
func (*MediaDevices) EnumerateDevices() (m []*mediadeviceinfo.MediaDeviceInfo) {
	js.Rewrite("await $<.EnumerateDevices()")
	return m
}

// GetSupportedConstraints
func (*MediaDevices) GetSupportedConstraints() (m *mediatracksupportedconstraints.MediaTrackSupportedConstraints) {
	js.Rewrite("$<.GetSupportedConstraints()")
	return m
}

// GetUserMedia
func (*MediaDevices) GetUserMedia(constraints *mediastreamconstraints.MediaStreamConstraints) (m *MediaStream) {
	js.Rewrite("await $<.GetUserMedia($1)", constraints)
	return m
}

// Ondevicechange
func (*MediaDevices) Ondevicechange() (ondevicechange func(Event)) {
	js.Rewrite("$<.Ondevicechange")
	return ondevicechange
}

// Ondevicechange
func (*MediaDevices) SetOndevicechange(ondevicechange func(Event)) {
	js.Rewrite("$<.Ondevicechange = $1", ondevicechange)
}
