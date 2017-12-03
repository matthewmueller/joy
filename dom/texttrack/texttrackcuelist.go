package texttrack

import "github.com/matthewmueller/joy/macro"

// TextTrackCueList struct
// js:"TextTrackCueList,omit"
type TextTrackCueList struct {
}

// GetCueByID fn
// js:"getCueById"
func (*TextTrackCueList) GetCueByID(id string) (t TextTrackCue) {
	macro.Rewrite("$_.getCueById($1)", id)
	return t
}

// Item fn
// js:"item"
func (*TextTrackCueList) Item(index uint) (t TextTrackCue) {
	macro.Rewrite("$_.item($1)", index)
	return t
}

// Length prop
// js:"length"
func (*TextTrackCueList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
