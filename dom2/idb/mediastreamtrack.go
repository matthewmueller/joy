package idb

import (
	"github.com/matthewmueller/golly/dom2/mediastreamtrackstate"
	"github.com/matthewmueller/golly/dom2/mediatrackcapabilities"
	"github.com/matthewmueller/golly/dom2/mediatrackconstraints"
	"github.com/matthewmueller/golly/dom2/mediatracksettings"
	"github.com/matthewmueller/golly/js"
)

// MediaStreamTrack struct
// js:"MediaStreamTrack,omit"
type MediaStreamTrack struct {
	EventTarget
}

// ApplyConstraints
func (*MediaStreamTrack) ApplyConstraints(constraints *mediatrackconstraints.MediaTrackConstraints) {
	js.Rewrite("await $<.ApplyConstraints($1)", constraints)
}

// Clone
func (*MediaStreamTrack) Clone() (m *MediaStreamTrack) {
	js.Rewrite("$<.Clone()")
	return m
}

// GetCapabilities
func (*MediaStreamTrack) GetCapabilities() (m *mediatrackcapabilities.MediaTrackCapabilities) {
	js.Rewrite("$<.GetCapabilities()")
	return m
}

// GetConstraints
func (*MediaStreamTrack) GetConstraints() (m *mediatrackconstraints.MediaTrackConstraints) {
	js.Rewrite("$<.GetConstraints()")
	return m
}

// GetSettings
func (*MediaStreamTrack) GetSettings() (m *mediatracksettings.MediaTrackSettings) {
	js.Rewrite("$<.GetSettings()")
	return m
}

// Stop
func (*MediaStreamTrack) Stop() {
	js.Rewrite("$<.Stop()")
}

// Enabled
func (*MediaStreamTrack) Enabled() (enabled bool) {
	js.Rewrite("$<.Enabled")
	return enabled
}

// Enabled
func (*MediaStreamTrack) SetEnabled(enabled bool) {
	js.Rewrite("$<.Enabled = $1", enabled)
}

// ID
func (*MediaStreamTrack) ID() (id string) {
	js.Rewrite("$<.ID")
	return id
}

// Kind
func (*MediaStreamTrack) Kind() (kind string) {
	js.Rewrite("$<.Kind")
	return kind
}

// Label
func (*MediaStreamTrack) Label() (label string) {
	js.Rewrite("$<.Label")
	return label
}

// Muted
func (*MediaStreamTrack) Muted() (muted bool) {
	js.Rewrite("$<.Muted")
	return muted
}

// Onended
func (*MediaStreamTrack) Onended() (onended func(*MediaStreamErrorEvent)) {
	js.Rewrite("$<.Onended")
	return onended
}

// Onended
func (*MediaStreamTrack) SetOnended(onended func(*MediaStreamErrorEvent)) {
	js.Rewrite("$<.Onended = $1", onended)
}

// Onmute
func (*MediaStreamTrack) Onmute() (onmute func(Event)) {
	js.Rewrite("$<.Onmute")
	return onmute
}

// Onmute
func (*MediaStreamTrack) SetOnmute(onmute func(Event)) {
	js.Rewrite("$<.Onmute = $1", onmute)
}

// Onoverconstrained
func (*MediaStreamTrack) Onoverconstrained() (onoverconstrained func(*MediaStreamErrorEvent)) {
	js.Rewrite("$<.Onoverconstrained")
	return onoverconstrained
}

// Onoverconstrained
func (*MediaStreamTrack) SetOnoverconstrained(onoverconstrained func(*MediaStreamErrorEvent)) {
	js.Rewrite("$<.Onoverconstrained = $1", onoverconstrained)
}

// Onunmute
func (*MediaStreamTrack) Onunmute() (onunmute func(Event)) {
	js.Rewrite("$<.Onunmute")
	return onunmute
}

// Onunmute
func (*MediaStreamTrack) SetOnunmute(onunmute func(Event)) {
	js.Rewrite("$<.Onunmute = $1", onunmute)
}

// Readonly
func (*MediaStreamTrack) Readonly() (readonly bool) {
	js.Rewrite("$<.Readonly")
	return readonly
}

// ReadyState
func (*MediaStreamTrack) ReadyState() (readyState *mediastreamtrackstate.MediaStreamTrackState) {
	js.Rewrite("$<.ReadyState")
	return readyState
}

// Remote
func (*MediaStreamTrack) Remote() (remote bool) {
	js.Rewrite("$<.Remote")
	return remote
}
