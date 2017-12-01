package texttracklist

import (
	"github.com/matthewmueller/golly/dom/texttrack"
	"github.com/matthewmueller/golly/dom/trackevent"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.EventTarget = (*TextTrackList)(nil)

// TextTrackList struct
// js:"TextTrackList,omit"
type TextTrackList struct {
}

// Item fn
// js:"item"
func (*TextTrackList) Item(index uint) (t *texttrack.TextTrack) {
	js.Rewrite("$_.item($1)", index)
	return t
}

// AddEventListener fn
// js:"addEventListener"
func (*TextTrackList) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*TextTrackList) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*TextTrackList) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Length prop
// js:"length"
func (*TextTrackList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}

// Onaddtrack prop
// js:"onaddtrack"
func (*TextTrackList) Onaddtrack() (onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$_.onaddtrack")
	return onaddtrack
}

// SetOnaddtrack prop
// js:"onaddtrack"
func (*TextTrackList) SetOnaddtrack(onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$_.onaddtrack = $1", onaddtrack)
}
