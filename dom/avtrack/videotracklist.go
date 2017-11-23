package avtrack

import (
	"github.com/matthewmueller/golly/dom2/trackevent"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// VideoTrackList struct
// js:"VideoTrackList,omit"
type VideoTrackList struct {
	window.EventTarget
}

// GetTrackByID fn
func (*VideoTrackList) GetTrackByID(id string) (v *VideoTrack) {
	js.Rewrite("$<.getTrackById($1)", id)
	return v
}

// Item fn
func (*VideoTrackList) Item(index uint) (v *VideoTrack) {
	js.Rewrite("$<.item($1)", index)
	return v
}

// Length prop
func (*VideoTrackList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}

// Onaddtrack prop
func (*VideoTrackList) Onaddtrack() (onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.onaddtrack")
	return onaddtrack
}

// Onaddtrack prop
func (*VideoTrackList) SetOnaddtrack(onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.onaddtrack = $1", onaddtrack)
}

// Onchange prop
func (*VideoTrackList) Onchange() (onchange func(window.Event)) {
	js.Rewrite("$<.onchange")
	return onchange
}

// Onchange prop
func (*VideoTrackList) SetOnchange(onchange func(window.Event)) {
	js.Rewrite("$<.onchange = $1", onchange)
}

// Onremovetrack prop
func (*VideoTrackList) Onremovetrack() (onremovetrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.onremovetrack")
	return onremovetrack
}

// Onremovetrack prop
func (*VideoTrackList) SetOnremovetrack(onremovetrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.onremovetrack = $1", onremovetrack)
}

// SelectedIndex prop
func (*VideoTrackList) SelectedIndex() (selectedIndex int) {
	js.Rewrite("$<.selectedIndex")
	return selectedIndex
}
