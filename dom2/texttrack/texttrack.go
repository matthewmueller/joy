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

// AddCue
func (*TextTrack) AddCue(cue TextTrackCue) {
	js.Rewrite("$<.AddCue($1)", cue)
}

// RemoveCue
func (*TextTrack) RemoveCue(cue TextTrackCue) {
	js.Rewrite("$<.RemoveCue($1)", cue)
}

// ActiveCues
func (*TextTrack) ActiveCues() (activeCues *TextTrackCueList) {
	js.Rewrite("$<.ActiveCues")
	return activeCues
}

// Cues
func (*TextTrack) Cues() (cues *TextTrackCueList) {
	js.Rewrite("$<.Cues")
	return cues
}

// InBandMetadataTrackDispatchType
func (*TextTrack) InBandMetadataTrackDispatchType() (inBandMetadataTrackDispatchType string) {
	js.Rewrite("$<.InBandMetadataTrackDispatchType")
	return inBandMetadataTrackDispatchType
}

// Kind
func (*TextTrack) Kind() (kind string) {
	js.Rewrite("$<.Kind")
	return kind
}

// Label
func (*TextTrack) Label() (label string) {
	js.Rewrite("$<.Label")
	return label
}

// Language
func (*TextTrack) Language() (language string) {
	js.Rewrite("$<.Language")
	return language
}

// Mode
func (*TextTrack) Mode() (mode interface{}) {
	js.Rewrite("$<.Mode")
	return mode
}

// Mode
func (*TextTrack) SetMode(mode interface{}) {
	js.Rewrite("$<.Mode = $1", mode)
}

// Oncuechange
func (*TextTrack) Oncuechange() (oncuechange func(window.Event)) {
	js.Rewrite("$<.Oncuechange")
	return oncuechange
}

// Oncuechange
func (*TextTrack) SetOncuechange(oncuechange func(window.Event)) {
	js.Rewrite("$<.Oncuechange = $1", oncuechange)
}

// Onerror
func (*TextTrack) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*TextTrack) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onload
func (*TextTrack) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.Onload")
	return onload
}

// Onload
func (*TextTrack) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.Onload = $1", onload)
}

// ReadyState
func (*TextTrack) ReadyState() (readyState uint8) {
	js.Rewrite("$<.ReadyState")
	return readyState
}
