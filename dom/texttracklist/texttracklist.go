package texttracklist

import (
	"github.com/matthewmueller/joy/dom/texttrack"
	"github.com/matthewmueller/joy/dom/trackevent"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.EventTarget = (*TextTrackList)(nil)

// TextTrackList struct
// js:"TextTrackList,omit"
type TextTrackList struct {
}

// Item fn
// js:"item"
func (*TextTrackList) Item(index uint) (t *texttrack.TextTrack) {
	macro.Rewrite("$_.item($1)", index)
	return t
}

// AddEventListener fn
// js:"addEventListener"
func (*TextTrackList) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*TextTrackList) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*TextTrackList) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Length prop
// js:"length"
func (*TextTrackList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

// Onaddtrack prop
// js:"onaddtrack"
func (*TextTrackList) Onaddtrack() (onaddtrack func(*trackevent.TrackEvent)) {
	macro.Rewrite("$_.onaddtrack")
	return onaddtrack
}

// SetOnaddtrack prop
// js:"onaddtrack"
func (*TextTrackList) SetOnaddtrack(onaddtrack func(*trackevent.TrackEvent)) {
	macro.Rewrite("$_.onaddtrack = $1", onaddtrack)
}
