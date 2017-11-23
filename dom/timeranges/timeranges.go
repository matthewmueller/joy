package timeranges

import "github.com/matthewmueller/golly/js"

// TimeRanges struct
// js:"TimeRanges,omit"
type TimeRanges struct {
}

// End fn
func (*TimeRanges) End(index uint) (f float32) {
	js.Rewrite("$<.end($1)", index)
	return f
}

// Start fn
func (*TimeRanges) Start(index uint) (f float32) {
	js.Rewrite("$<.start($1)", index)
	return f
}

// Length prop
func (*TimeRanges) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
