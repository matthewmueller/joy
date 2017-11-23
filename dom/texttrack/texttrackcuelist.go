package texttrack

import "github.com/matthewmueller/golly/js"

// TextTrackCueList struct
// js:"TextTrackCueList,omit"
type TextTrackCueList struct {
}

// GetCueByID fn
func (*TextTrackCueList) GetCueByID(id string) (t TextTrackCue) {
	js.Rewrite("$<.getCueById($1)", id)
	return t
}

// Item fn
func (*TextTrackCueList) Item(index uint) (t TextTrackCue) {
	js.Rewrite("$<.item($1)", index)
	return t
}

// Length prop
func (*TextTrackCueList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
