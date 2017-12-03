package window

import "github.com/matthewmueller/joy/macro"

// MSRangeCollection struct
// js:"MSRangeCollection,omit"
type MSRangeCollection struct {
}

// Item fn
// js:"item"
func (*MSRangeCollection) Item(index uint) (r *Range) {
	macro.Rewrite("$_.item($1)", index)
	return r
}

// Length prop
// js:"length"
func (*MSRangeCollection) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
