package texttracklist

import (
	"github.com/matthewmueller/golly/dom2/texttrack"
	"github.com/matthewmueller/golly/dom2/trackevent"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"TextTrackList,omit"
type TextTrackList struct {
	window.EventTarget
}

// Item
func (*TextTrackList) Item(index uint) (t *texttrack.TextTrack) {
	js.Rewrite("$<.Item($1)", index)
	return t
}

// Length
func (*TextTrackList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}

// Onaddtrack
func (*TextTrackList) Onaddtrack() (onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.Onaddtrack")
	return onaddtrack
}

// Onaddtrack
func (*TextTrackList) SetOnaddtrack(onaddtrack func(*trackevent.TrackEvent)) {
	js.Rewrite("$<.Onaddtrack = $1", onaddtrack)
}
