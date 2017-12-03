package texttrack

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.EventTarget = (*TextTrack)(nil)

// TextTrack struct
// js:"TextTrack,omit"
type TextTrack struct {
}

// AddCue fn
// js:"addCue"
func (*TextTrack) AddCue(cue TextTrackCue) {
	js.Rewrite("$_.addCue($1)", cue)
}

// RemoveCue fn
// js:"removeCue"
func (*TextTrack) RemoveCue(cue TextTrackCue) {
	js.Rewrite("$_.removeCue($1)", cue)
}

// AddEventListener fn
// js:"addEventListener"
func (*TextTrack) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*TextTrack) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*TextTrack) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// ActiveCues prop
// js:"activeCues"
func (*TextTrack) ActiveCues() (activeCues *TextTrackCueList) {
	js.Rewrite("$_.activeCues")
	return activeCues
}

// Cues prop
// js:"cues"
func (*TextTrack) Cues() (cues *TextTrackCueList) {
	js.Rewrite("$_.cues")
	return cues
}

// InBandMetadataTrackDispatchType prop
// js:"inBandMetadataTrackDispatchType"
func (*TextTrack) InBandMetadataTrackDispatchType() (inBandMetadataTrackDispatchType string) {
	js.Rewrite("$_.inBandMetadataTrackDispatchType")
	return inBandMetadataTrackDispatchType
}

// Kind prop
// js:"kind"
func (*TextTrack) Kind() (kind string) {
	js.Rewrite("$_.kind")
	return kind
}

// Label prop
// js:"label"
func (*TextTrack) Label() (label string) {
	js.Rewrite("$_.label")
	return label
}

// Language prop
// js:"language"
func (*TextTrack) Language() (language string) {
	js.Rewrite("$_.language")
	return language
}

// Mode prop
// js:"mode"
func (*TextTrack) Mode() (mode interface{}) {
	js.Rewrite("$_.mode")
	return mode
}

// SetMode prop
// js:"mode"
func (*TextTrack) SetMode(mode interface{}) {
	js.Rewrite("$_.mode = $1", mode)
}

// Oncuechange prop
// js:"oncuechange"
func (*TextTrack) Oncuechange() (oncuechange func(window.Event)) {
	js.Rewrite("$_.oncuechange")
	return oncuechange
}

// SetOncuechange prop
// js:"oncuechange"
func (*TextTrack) SetOncuechange(oncuechange func(window.Event)) {
	js.Rewrite("$_.oncuechange = $1", oncuechange)
}

// Onerror prop
// js:"onerror"
func (*TextTrack) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*TextTrack) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onload prop
// js:"onload"
func (*TextTrack) Onload() (onload func(window.Event)) {
	js.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*TextTrack) SetOnload(onload func(window.Event)) {
	js.Rewrite("$_.onload = $1", onload)
}

// ReadyState prop
// js:"readyState"
func (*TextTrack) ReadyState() (readyState uint8) {
	js.Rewrite("$_.readyState")
	return readyState
}
