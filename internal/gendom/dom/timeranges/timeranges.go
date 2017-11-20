package timeranges

import "github.com/matthewmueller/golly/js"

type TimeRanges struct {
}

func (*TimeRanges) End(index uint) (f float32) {
	js.Rewrite("$<.end($1)", index)
	return f
}

func (*TimeRanges) Start(index uint) (f float32) {
	js.Rewrite("$<.start($1)", index)
	return f
}

func (*TimeRanges) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
