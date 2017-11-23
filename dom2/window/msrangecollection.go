package window

import "github.com/matthewmueller/golly/js"

// MSRangeCollection struct
// js:"MSRangeCollection,omit"
type MSRangeCollection struct {
}

// Item
func (*MSRangeCollection) Item(index uint) (r *Range) {
	js.Rewrite("$<.Item($1)", index)
	return r
}

// Length
func (*MSRangeCollection) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
