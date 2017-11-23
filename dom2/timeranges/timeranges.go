package timeranges

import "github.com/matthewmueller/golly/js"

// TimeRanges struct
// js:"TimeRanges,omit"
type TimeRanges struct {
}

// End
func (*TimeRanges) End(index uint) (f float32) {
	js.Rewrite("$<.End($1)", index)
	return f
}

// Start
func (*TimeRanges) Start(index uint) (f float32) {
	js.Rewrite("$<.Start($1)", index)
	return f
}

// Length
func (*TimeRanges) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
