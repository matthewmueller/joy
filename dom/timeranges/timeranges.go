package timeranges

import "github.com/matthewmueller/joy/js"

// TimeRanges struct
// js:"TimeRanges,omit"
type TimeRanges struct {
}

// End fn
// js:"end"
func (*TimeRanges) End(index uint) (f float32) {
	js.Rewrite("$_.end($1)", index)
	return f
}

// Start fn
// js:"start"
func (*TimeRanges) Start(index uint) (f float32) {
	js.Rewrite("$_.start($1)", index)
	return f
}

// Length prop
// js:"length"
func (*TimeRanges) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
