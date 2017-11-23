package window

import "github.com/matthewmueller/golly/js"

// MSRangeCollection struct
// js:"MSRangeCollection,omit"
type MSRangeCollection struct {
}

// Item fn
func (*MSRangeCollection) Item(index uint) (r *Range) {
	js.Rewrite("$<.item($1)", index)
	return r
}

// Length prop
func (*MSRangeCollection) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
