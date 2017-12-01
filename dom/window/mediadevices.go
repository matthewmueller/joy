package window

import (
	"github.com/matthewmueller/golly/dom/mediadeviceinfo"
	"github.com/matthewmueller/golly/dom/mediastreamconstraints"
	"github.com/matthewmueller/golly/dom/mediatracksupportedconstraints"
	"github.com/matthewmueller/golly/js"
)

var _ EventTarget = (*MediaDevices)(nil)

// MediaDevices struct
// js:"MediaDevices,omit"
type MediaDevices struct {
}

// EnumerateDevices fn
// js:"enumerateDevices"
func (*MediaDevices) EnumerateDevices() (m []*mediadeviceinfo.MediaDeviceInfo) {
	js.Rewrite("await $_.enumerateDevices()")
	return m
}

// GetSupportedConstraints fn
// js:"getSupportedConstraints"
func (*MediaDevices) GetSupportedConstraints() (m *mediatracksupportedconstraints.MediaTrackSupportedConstraints) {
	js.Rewrite("$_.getSupportedConstraints()")
	return m
}

// GetUserMedia fn
// js:"getUserMedia"
func (*MediaDevices) GetUserMedia(constraints *mediastreamconstraints.MediaStreamConstraints) (m *MediaStream) {
	js.Rewrite("await $_.getUserMedia($1)", constraints)
	return m
}

// AddEventListener fn
// js:"addEventListener"
func (*MediaDevices) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MediaDevices) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MediaDevices) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Ondevicechange prop
// js:"ondevicechange"
func (*MediaDevices) Ondevicechange() (ondevicechange func(Event)) {
	js.Rewrite("$_.ondevicechange")
	return ondevicechange
}

// SetOndevicechange prop
// js:"ondevicechange"
func (*MediaDevices) SetOndevicechange(ondevicechange func(Event)) {
	js.Rewrite("$_.ondevicechange = $1", ondevicechange)
}
