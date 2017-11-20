package msrangecollection

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/rng"
	"github.com/matthewmueller/golly/js"
)

type MSRangeCollection struct {
}

func (*MSRangeCollection) Item(index uint) (r *rng.Range) {
	js.Rewrite("$<.item($1)", index)
	return r
}

func (*MSRangeCollection) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
