package window

import (
	"github.com/matthewmueller/golly/dom/mediastreamtrackstate"
	"github.com/matthewmueller/golly/dom/mediatrackcapabilities"
	"github.com/matthewmueller/golly/dom/mediatrackconstraints"
	"github.com/matthewmueller/golly/dom/mediatracksettings"
	"github.com/matthewmueller/golly/js"
)

// MediaStreamTrack struct
// js:"MediaStreamTrack,omit"
type MediaStreamTrack struct {
	EventTarget
}

// ApplyConstraints fn
func (*MediaStreamTrack) ApplyConstraints(constraints *mediatrackconstraints.MediaTrackConstraints) {
	js.Rewrite("await $<.applyConstraints($1)", constraints)
}

// Clone fn
func (*MediaStreamTrack) Clone() (m *MediaStreamTrack) {
	js.Rewrite("$<.clone()")
	return m
}

// GetCapabilities fn
func (*MediaStreamTrack) GetCapabilities() (m *mediatrackcapabilities.MediaTrackCapabilities) {
	js.Rewrite("$<.getCapabilities()")
	return m
}

// GetConstraints fn
func (*MediaStreamTrack) GetConstraints() (m *mediatrackconstraints.MediaTrackConstraints) {
	js.Rewrite("$<.getConstraints()")
	return m
}

// GetSettings fn
func (*MediaStreamTrack) GetSettings() (m *mediatracksettings.MediaTrackSettings) {
	js.Rewrite("$<.getSettings()")
	return m
}

// Stop fn
func (*MediaStreamTrack) Stop() {
	js.Rewrite("$<.stop()")
}

// Enabled prop
func (*MediaStreamTrack) Enabled() (enabled bool) {
	js.Rewrite("$<.enabled")
	return enabled
}

// Enabled prop
func (*MediaStreamTrack) SetEnabled(enabled bool) {
	js.Rewrite("$<.enabled = $1", enabled)
}

// ID prop
func (*MediaStreamTrack) ID() (id string) {
	js.Rewrite("$<.id")
	return id
}

// Kind prop
func (*MediaStreamTrack) Kind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

// Label prop
func (*MediaStreamTrack) Label() (label string) {
	js.Rewrite("$<.label")
	return label
}

// Muted prop
func (*MediaStreamTrack) Muted() (muted bool) {
	js.Rewrite("$<.muted")
	return muted
}

// Onended prop
func (*MediaStreamTrack) Onended() (onended func(*MediaStreamErrorEvent)) {
	js.Rewrite("$<.onended")
	return onended
}

// Onended prop
func (*MediaStreamTrack) SetOnended(onended func(*MediaStreamErrorEvent)) {
	js.Rewrite("$<.onended = $1", onended)
}

// Onmute prop
func (*MediaStreamTrack) Onmute() (onmute func(Event)) {
	js.Rewrite("$<.onmute")
	return onmute
}

// Onmute prop
func (*MediaStreamTrack) SetOnmute(onmute func(Event)) {
	js.Rewrite("$<.onmute = $1", onmute)
}

// Onoverconstrained prop
func (*MediaStreamTrack) Onoverconstrained() (onoverconstrained func(*MediaStreamErrorEvent)) {
	js.Rewrite("$<.onoverconstrained")
	return onoverconstrained
}

// Onoverconstrained prop
func (*MediaStreamTrack) SetOnoverconstrained(onoverconstrained func(*MediaStreamErrorEvent)) {
	js.Rewrite("$<.onoverconstrained = $1", onoverconstrained)
}

// Onunmute prop
func (*MediaStreamTrack) Onunmute() (onunmute func(Event)) {
	js.Rewrite("$<.onunmute")
	return onunmute
}

// Onunmute prop
func (*MediaStreamTrack) SetOnunmute(onunmute func(Event)) {
	js.Rewrite("$<.onunmute = $1", onunmute)
}

// Readonly prop
func (*MediaStreamTrack) Readonly() (readonly bool) {
	js.Rewrite("$<.readonly")
	return readonly
}

// ReadyState prop
func (*MediaStreamTrack) ReadyState() (readyState *mediastreamtrackstate.MediaStreamTrackState) {
	js.Rewrite("$<.readyState")
	return readyState
}

// Remote prop
func (*MediaStreamTrack) Remote() (remote bool) {
	js.Rewrite("$<.remote")
	return remote
}
