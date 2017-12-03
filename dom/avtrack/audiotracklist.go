package avtrack

import (
	"github.com/matthewmueller/joy/dom/trackevent"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.EventTarget = (*AudioTrackList)(nil)

// AudioTrackList struct
// js:"AudioTrackList,omit"
type AudioTrackList struct {
}

// GetTrackByID fn
// js:"getTrackById"
func (*AudioTrackList) GetTrackByID(id string) (a *AudioTrack) {
	js.Rewrite("$_.getTrackById($1)", id)
	return a
}

// Item fn
// js:"item"
func (*AudioTrackList) Item(index uint) (a *AudioTrack) {
	js.Rewrite("$_.item($1)", index)
	return a
}

// AddEventListener fn
// js:"addEventListener"
func (*AudioTrackList) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*AudioTrackList) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*AudioTrackList) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Length prop
// js:"length"
func (*AudioTrackList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}

// Onaddtrack prop
// js:"onaddtrack"
func (*AudioTrackList) Onaddtrack() (onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$_.onaddtrack")
	return onaddtrack
}

// SetOnaddtrack prop
// js:"onaddtrack"
func (*AudioTrackList) SetOnaddtrack(onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$_.onaddtrack = $1", onaddtrack)
}

// Onchange prop
// js:"onchange"
func (*AudioTrackList) Onchange() (onchange func(window.Event)) {
	js.Rewrite("$_.onchange")
	return onchange
}

// SetOnchange prop
// js:"onchange"
func (*AudioTrackList) SetOnchange(onchange func(window.Event)) {
	js.Rewrite("$_.onchange = $1", onchange)
}

// Onremovetrack prop
// js:"onremovetrack"
func (*AudioTrackList) Onremovetrack() (onremovetrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$_.onremovetrack")
	return onremovetrack
}

// SetOnremovetrack prop
// js:"onremovetrack"
func (*AudioTrackList) SetOnremovetrack(onremovetrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$_.onremovetrack = $1", onremovetrack)
}
