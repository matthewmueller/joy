package avtrack

import (
	"github.com/matthewmueller/golly/dom/trackevent"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// AudioTrackList struct
// js:"AudioTrackList,omit"
type AudioTrackList struct {
	window.EventTarget
}

// GetTrackByID fn
func (*AudioTrackList) GetTrackByID(id string) (a *AudioTrack) {
	js.Rewrite("$<.getTrackById($1)", id)
	return a
}

// Item fn
func (*AudioTrackList) Item(index uint) (a *AudioTrack) {
	js.Rewrite("$<.item($1)", index)
	return a
}

// Length prop
func (*AudioTrackList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}

// Onaddtrack prop
func (*AudioTrackList) Onaddtrack() (onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.onaddtrack")
	return onaddtrack
}

// Onaddtrack prop
func (*AudioTrackList) SetOnaddtrack(onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.onaddtrack = $1", onaddtrack)
}

// Onchange prop
func (*AudioTrackList) Onchange() (onchange func(window.Event)) {
	js.Rewrite("$<.onchange")
	return onchange
}

// Onchange prop
func (*AudioTrackList) SetOnchange(onchange func(window.Event)) {
	js.Rewrite("$<.onchange = $1", onchange)
}

// Onremovetrack prop
func (*AudioTrackList) Onremovetrack() (onremovetrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.onremovetrack")
	return onremovetrack
}

// Onremovetrack prop
func (*AudioTrackList) SetOnremovetrack(onremovetrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.onremovetrack = $1", onremovetrack)
}
