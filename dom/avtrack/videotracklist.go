package avtrack

import (
	"github.com/matthewmueller/joy/dom/trackevent"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.EventTarget = (*VideoTrackList)(nil)

// VideoTrackList struct
// js:"VideoTrackList,omit"
type VideoTrackList struct {
}

// GetTrackByID fn
// js:"getTrackById"
func (*VideoTrackList) GetTrackByID(id string) (v *VideoTrack) {
	js.Rewrite("$_.getTrackById($1)", id)
	return v
}

// Item fn
// js:"item"
func (*VideoTrackList) Item(index uint) (v *VideoTrack) {
	js.Rewrite("$_.item($1)", index)
	return v
}

// AddEventListener fn
// js:"addEventListener"
func (*VideoTrackList) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*VideoTrackList) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*VideoTrackList) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Length prop
// js:"length"
func (*VideoTrackList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}

// Onaddtrack prop
// js:"onaddtrack"
func (*VideoTrackList) Onaddtrack() (onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$_.onaddtrack")
	return onaddtrack
}

// SetOnaddtrack prop
// js:"onaddtrack"
func (*VideoTrackList) SetOnaddtrack(onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$_.onaddtrack = $1", onaddtrack)
}

// Onchange prop
// js:"onchange"
func (*VideoTrackList) Onchange() (onchange func(window.Event)) {
	js.Rewrite("$_.onchange")
	return onchange
}

// SetOnchange prop
// js:"onchange"
func (*VideoTrackList) SetOnchange(onchange func(window.Event)) {
	js.Rewrite("$_.onchange = $1", onchange)
}

// Onremovetrack prop
// js:"onremovetrack"
func (*VideoTrackList) Onremovetrack() (onremovetrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$_.onremovetrack")
	return onremovetrack
}

// SetOnremovetrack prop
// js:"onremovetrack"
func (*VideoTrackList) SetOnremovetrack(onremovetrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$_.onremovetrack = $1", onremovetrack)
}

// SelectedIndex prop
// js:"selectedIndex"
func (*VideoTrackList) SelectedIndex() (selectedIndex int) {
	js.Rewrite("$_.selectedIndex")
	return selectedIndex
}
