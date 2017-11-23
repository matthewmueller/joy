package texttracklist

import (
	"github.com/matthewmueller/golly/dom/texttrack"
	"github.com/matthewmueller/golly/dom/trackevent"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// TextTrackList struct
// js:"TextTrackList,omit"
type TextTrackList struct {
	window.EventTarget
}

// Item fn
func (*TextTrackList) Item(index uint) (t *texttrack.TextTrack) {
	js.Rewrite("$<.item($1)", index)
	return t
}

// Length prop
func (*TextTrackList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}

// Onaddtrack prop
func (*TextTrackList) Onaddtrack() (onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.onaddtrack")
	return onaddtrack
}

// Onaddtrack prop
func (*TextTrackList) SetOnaddtrack(onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.onaddtrack = $1", onaddtrack)
}
