package avtrack

import (
	"github.com/matthewmueller/golly/dom2/trackevent"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"AudioTrackList,omit"
type AudioTrackList struct {
	window.EventTarget
}

// GetTrackByID
func (*AudioTrackList) GetTrackByID(id string) (a *AudioTrack) {
	js.Rewrite("$<.GetTrackByID($1)", id)
	return a
}

// Item
func (*AudioTrackList) Item(index uint) (a *AudioTrack) {
	js.Rewrite("$<.Item($1)", index)
	return a
}

// Length
func (*AudioTrackList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}

// Onaddtrack
func (*AudioTrackList) Onaddtrack() (onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.Onaddtrack")
	return onaddtrack
}

// Onaddtrack
func (*AudioTrackList) SetOnaddtrack(onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.Onaddtrack = $1", onaddtrack)
}

// Onchange
func (*AudioTrackList) Onchange() (onchange func(window.Event)) {
	js.Rewrite("$<.Onchange")
	return onchange
}

// Onchange
func (*AudioTrackList) SetOnchange(onchange func(window.Event)) {
	js.Rewrite("$<.Onchange = $1", onchange)
}

// Onremovetrack
func (*AudioTrackList) Onremovetrack() (onremovetrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.Onremovetrack")
	return onremovetrack
}

// Onremovetrack
func (*AudioTrackList) SetOnremovetrack(onremovetrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.Onremovetrack = $1", onremovetrack)
}
