package texttrack

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// TextTrack struct
// js:"TextTrack,omit"
type TextTrack struct {
	window.EventTarget
}

// AddCue fn
func (*TextTrack) AddCue(cue TextTrackCue) {
	js.Rewrite("$<.addCue($1)", cue)
}

// RemoveCue fn
func (*TextTrack) RemoveCue(cue TextTrackCue) {
	js.Rewrite("$<.removeCue($1)", cue)
}

// ActiveCues prop
func (*TextTrack) ActiveCues() (activeCues *TextTrackCueList) {
	js.Rewrite("$<.activeCues")
	return activeCues
}

// Cues prop
func (*TextTrack) Cues() (cues *TextTrackCueList) {
	js.Rewrite("$<.cues")
	return cues
}

// InBandMetadataTrackDispatchType prop
func (*TextTrack) InBandMetadataTrackDispatchType() (inBandMetadataTrackDispatchType string) {
	js.Rewrite("$<.inBandMetadataTrackDispatchType")
	return inBandMetadataTrackDispatchType
}

// Kind prop
func (*TextTrack) Kind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

// Label prop
func (*TextTrack) Label() (label string) {
	js.Rewrite("$<.label")
	return label
}

// Language prop
func (*TextTrack) Language() (language string) {
	js.Rewrite("$<.language")
	return language
}

// Mode prop
func (*TextTrack) Mode() (mode interface{}) {
	js.Rewrite("$<.mode")
	return mode
}

// Mode prop
func (*TextTrack) SetMode(mode interface{}) {
	js.Rewrite("$<.mode = $1", mode)
}

// Oncuechange prop
func (*TextTrack) Oncuechange() (oncuechange func(window.Event)) {
	js.Rewrite("$<.oncuechange")
	return oncuechange
}

// Oncuechange prop
func (*TextTrack) SetOncuechange(oncuechange func(window.Event)) {
	js.Rewrite("$<.oncuechange = $1", oncuechange)
}

// Onerror prop
func (*TextTrack) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*TextTrack) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onload prop
func (*TextTrack) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// Onload prop
func (*TextTrack) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.onload = $1", onload)
}

// ReadyState prop
func (*TextTrack) ReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}
