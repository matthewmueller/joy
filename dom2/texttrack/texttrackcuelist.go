package texttrack

import "github.com/matthewmueller/golly/js"

// TextTrackCueList struct
// js:"TextTrackCueList,omit"
type TextTrackCueList struct {
}

// GetCueByID
func (*TextTrackCueList) GetCueByID(id string) (t TextTrackCue) {
	js.Rewrite("$<.GetCueByID($1)", id)
	return t
}

// Item
func (*TextTrackCueList) Item(index uint) (t TextTrackCue) {
	js.Rewrite("$<.Item($1)", index)
	return t
}

// Length
func (*TextTrackCueList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
