package texttrack

import "github.com/matthewmueller/joy/js"

// TextTrackCueList struct
// js:"TextTrackCueList,omit"
type TextTrackCueList struct {
}

// GetCueByID fn
// js:"getCueById"
func (*TextTrackCueList) GetCueByID(id string) (t TextTrackCue) {
	js.Rewrite("$_.getCueById($1)", id)
	return t
}

// Item fn
// js:"item"
func (*TextTrackCueList) Item(index uint) (t TextTrackCue) {
	js.Rewrite("$_.item($1)", index)
	return t
}

// Length prop
// js:"length"
func (*TextTrackCueList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
