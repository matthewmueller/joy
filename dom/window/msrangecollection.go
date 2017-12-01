package window

import "github.com/matthewmueller/golly/js"

// MSRangeCollection struct
// js:"MSRangeCollection,omit"
type MSRangeCollection struct {
}

// Item fn
// js:"item"
func (*MSRangeCollection) Item(index uint) (r *Range) {
	js.Rewrite("$_.item($1)", index)
	return r
}

// Length prop
// js:"length"
func (*MSRangeCollection) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
