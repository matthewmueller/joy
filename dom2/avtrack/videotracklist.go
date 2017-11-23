package avtrack

import (
	"github.com/matthewmueller/golly/dom2/trackevent"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"VideoTrackList,omit"
type VideoTrackList struct {
	window.EventTarget
}

// GetTrackByID
func (*VideoTrackList) GetTrackByID(id string) (v *VideoTrack) {
	js.Rewrite("$<.GetTrackByID($1)", id)
	return v
}

// Item
func (*VideoTrackList) Item(index uint) (v *VideoTrack) {
	js.Rewrite("$<.Item($1)", index)
	return v
}

// Length
func (*VideoTrackList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}

// Onaddtrack
func (*VideoTrackList) Onaddtrack() (onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.Onaddtrack")
	return onaddtrack
}

// Onaddtrack
func (*VideoTrackList) SetOnaddtrack(onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.Onaddtrack = $1", onaddtrack)
}

// Onchange
func (*VideoTrackList) Onchange() (onchange func(window.Event)) {
	js.Rewrite("$<.Onchange")
	return onchange
}

// Onchange
func (*VideoTrackList) SetOnchange(onchange func(window.Event)) {
	js.Rewrite("$<.Onchange = $1", onchange)
}

// Onremovetrack
func (*VideoTrackList) Onremovetrack() (onremovetrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.Onremovetrack")
	return onremovetrack
}

// Onremovetrack
func (*VideoTrackList) SetOnremovetrack(onremovetrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.Onremovetrack = $1", onremovetrack)
}

// SelectedIndex
func (*VideoTrackList) SelectedIndex() (selectedIndex int) {
	js.Rewrite("$<.SelectedIndex")
	return selectedIndex
}
