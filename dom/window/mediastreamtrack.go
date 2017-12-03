package window

import (
	"github.com/matthewmueller/joy/dom/mediastreamtrackstate"
	"github.com/matthewmueller/joy/dom/mediatrackcapabilities"
	"github.com/matthewmueller/joy/dom/mediatrackconstraints"
	"github.com/matthewmueller/joy/dom/mediatracksettings"
	"github.com/matthewmueller/joy/macro"
)

var _ EventTarget = (*MediaStreamTrack)(nil)

// MediaStreamTrack struct
// js:"MediaStreamTrack,omit"
type MediaStreamTrack struct {
}

// ApplyConstraints fn
// js:"applyConstraints"
func (*MediaStreamTrack) ApplyConstraints(constraints *mediatrackconstraints.MediaTrackConstraints) {
	macro.Rewrite("await $_.applyConstraints($1)", constraints)
}

// Clone fn
// js:"clone"
func (*MediaStreamTrack) Clone() (m *MediaStreamTrack) {
	macro.Rewrite("$_.clone()")
	return m
}

// GetCapabilities fn
// js:"getCapabilities"
func (*MediaStreamTrack) GetCapabilities() (m *mediatrackcapabilities.MediaTrackCapabilities) {
	macro.Rewrite("$_.getCapabilities()")
	return m
}

// GetConstraints fn
// js:"getConstraints"
func (*MediaStreamTrack) GetConstraints() (m *mediatrackconstraints.MediaTrackConstraints) {
	macro.Rewrite("$_.getConstraints()")
	return m
}

// GetSettings fn
// js:"getSettings"
func (*MediaStreamTrack) GetSettings() (m *mediatracksettings.MediaTrackSettings) {
	macro.Rewrite("$_.getSettings()")
	return m
}

// Stop fn
// js:"stop"
func (*MediaStreamTrack) Stop() {
	macro.Rewrite("$_.stop()")
}

// AddEventListener fn
// js:"addEventListener"
func (*MediaStreamTrack) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MediaStreamTrack) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MediaStreamTrack) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Enabled prop
// js:"enabled"
func (*MediaStreamTrack) Enabled() (enabled bool) {
	macro.Rewrite("$_.enabled")
	return enabled
}

// SetEnabled prop
// js:"enabled"
func (*MediaStreamTrack) SetEnabled(enabled bool) {
	macro.Rewrite("$_.enabled = $1", enabled)
}

// ID prop
// js:"id"
func (*MediaStreamTrack) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// Kind prop
// js:"kind"
func (*MediaStreamTrack) Kind() (kind string) {
	macro.Rewrite("$_.kind")
	return kind
}

// Label prop
// js:"label"
func (*MediaStreamTrack) Label() (label string) {
	macro.Rewrite("$_.label")
	return label
}

// Muted prop
// js:"muted"
func (*MediaStreamTrack) Muted() (muted bool) {
	macro.Rewrite("$_.muted")
	return muted
}

// Onended prop
// js:"onended"
func (*MediaStreamTrack) Onended() (onended func(*MediaStreamErrorEvent)) {
	macro.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*MediaStreamTrack) SetOnended(onended func(*MediaStreamErrorEvent)) {
	macro.Rewrite("$_.onended = $1", onended)
}

// Onmute prop
// js:"onmute"
func (*MediaStreamTrack) Onmute() (onmute func(Event)) {
	macro.Rewrite("$_.onmute")
	return onmute
}

// SetOnmute prop
// js:"onmute"
func (*MediaStreamTrack) SetOnmute(onmute func(Event)) {
	macro.Rewrite("$_.onmute = $1", onmute)
}

// Onoverconstrained prop
// js:"onoverconstrained"
func (*MediaStreamTrack) Onoverconstrained() (onoverconstrained func(*MediaStreamErrorEvent)) {
	macro.Rewrite("$_.onoverconstrained")
	return onoverconstrained
}

// SetOnoverconstrained prop
// js:"onoverconstrained"
func (*MediaStreamTrack) SetOnoverconstrained(onoverconstrained func(*MediaStreamErrorEvent)) {
	macro.Rewrite("$_.onoverconstrained = $1", onoverconstrained)
}

// Onunmute prop
// js:"onunmute"
func (*MediaStreamTrack) Onunmute() (onunmute func(Event)) {
	macro.Rewrite("$_.onunmute")
	return onunmute
}

// SetOnunmute prop
// js:"onunmute"
func (*MediaStreamTrack) SetOnunmute(onunmute func(Event)) {
	macro.Rewrite("$_.onunmute = $1", onunmute)
}

// Readonly prop
// js:"readonly"
func (*MediaStreamTrack) Readonly() (readonly bool) {
	macro.Rewrite("$_.readonly")
	return readonly
}

// ReadyState prop
// js:"readyState"
func (*MediaStreamTrack) ReadyState() (readyState *mediastreamtrackstate.MediaStreamTrackState) {
	macro.Rewrite("$_.readyState")
	return readyState
}

// Remote prop
// js:"remote"
func (*MediaStreamTrack) Remote() (remote bool) {
	macro.Rewrite("$_.remote")
	return remote
}
