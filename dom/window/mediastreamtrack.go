package window

import (
	"github.com/matthewmueller/golly/dom/mediastreamtrackstate"
	"github.com/matthewmueller/golly/dom/mediatrackcapabilities"
	"github.com/matthewmueller/golly/dom/mediatrackconstraints"
	"github.com/matthewmueller/golly/dom/mediatracksettings"
	"github.com/matthewmueller/golly/js"
)

var _ EventTarget = (*MediaStreamTrack)(nil)

// MediaStreamTrack struct
// js:"MediaStreamTrack,omit"
type MediaStreamTrack struct {
}

// ApplyConstraints fn
// js:"applyConstraints"
func (*MediaStreamTrack) ApplyConstraints(constraints *mediatrackconstraints.MediaTrackConstraints) {
	js.Rewrite("await $_.applyConstraints($1)", constraints)
}

// Clone fn
// js:"clone"
func (*MediaStreamTrack) Clone() (m *MediaStreamTrack) {
	js.Rewrite("$_.clone()")
	return m
}

// GetCapabilities fn
// js:"getCapabilities"
func (*MediaStreamTrack) GetCapabilities() (m *mediatrackcapabilities.MediaTrackCapabilities) {
	js.Rewrite("$_.getCapabilities()")
	return m
}

// GetConstraints fn
// js:"getConstraints"
func (*MediaStreamTrack) GetConstraints() (m *mediatrackconstraints.MediaTrackConstraints) {
	js.Rewrite("$_.getConstraints()")
	return m
}

// GetSettings fn
// js:"getSettings"
func (*MediaStreamTrack) GetSettings() (m *mediatracksettings.MediaTrackSettings) {
	js.Rewrite("$_.getSettings()")
	return m
}

// Stop fn
// js:"stop"
func (*MediaStreamTrack) Stop() {
	js.Rewrite("$_.stop()")
}

// AddEventListener fn
// js:"addEventListener"
func (*MediaStreamTrack) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MediaStreamTrack) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MediaStreamTrack) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Enabled prop
// js:"enabled"
func (*MediaStreamTrack) Enabled() (enabled bool) {
	js.Rewrite("$_.enabled")
	return enabled
}

// SetEnabled prop
// js:"enabled"
func (*MediaStreamTrack) SetEnabled(enabled bool) {
	js.Rewrite("$_.enabled = $1", enabled)
}

// ID prop
// js:"id"
func (*MediaStreamTrack) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// Kind prop
// js:"kind"
func (*MediaStreamTrack) Kind() (kind string) {
	js.Rewrite("$_.kind")
	return kind
}

// Label prop
// js:"label"
func (*MediaStreamTrack) Label() (label string) {
	js.Rewrite("$_.label")
	return label
}

// Muted prop
// js:"muted"
func (*MediaStreamTrack) Muted() (muted bool) {
	js.Rewrite("$_.muted")
	return muted
}

// Onended prop
// js:"onended"
func (*MediaStreamTrack) Onended() (onended func(*MediaStreamErrorEvent)) {
	js.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*MediaStreamTrack) SetOnended(onended func(*MediaStreamErrorEvent)) {
	js.Rewrite("$_.onended = $1", onended)
}

// Onmute prop
// js:"onmute"
func (*MediaStreamTrack) Onmute() (onmute func(Event)) {
	js.Rewrite("$_.onmute")
	return onmute
}

// SetOnmute prop
// js:"onmute"
func (*MediaStreamTrack) SetOnmute(onmute func(Event)) {
	js.Rewrite("$_.onmute = $1", onmute)
}

// Onoverconstrained prop
// js:"onoverconstrained"
func (*MediaStreamTrack) Onoverconstrained() (onoverconstrained func(*MediaStreamErrorEvent)) {
	js.Rewrite("$_.onoverconstrained")
	return onoverconstrained
}

// SetOnoverconstrained prop
// js:"onoverconstrained"
func (*MediaStreamTrack) SetOnoverconstrained(onoverconstrained func(*MediaStreamErrorEvent)) {
	js.Rewrite("$_.onoverconstrained = $1", onoverconstrained)
}

// Onunmute prop
// js:"onunmute"
func (*MediaStreamTrack) Onunmute() (onunmute func(Event)) {
	js.Rewrite("$_.onunmute")
	return onunmute
}

// SetOnunmute prop
// js:"onunmute"
func (*MediaStreamTrack) SetOnunmute(onunmute func(Event)) {
	js.Rewrite("$_.onunmute = $1", onunmute)
}

// Readonly prop
// js:"readonly"
func (*MediaStreamTrack) Readonly() (readonly bool) {
	js.Rewrite("$_.readonly")
	return readonly
}

// ReadyState prop
// js:"readyState"
func (*MediaStreamTrack) ReadyState() (readyState *mediastreamtrackstate.MediaStreamTrackState) {
	js.Rewrite("$_.readyState")
	return readyState
}

// Remote prop
// js:"remote"
func (*MediaStreamTrack) Remote() (remote bool) {
	js.Rewrite("$_.remote")
	return remote
}
