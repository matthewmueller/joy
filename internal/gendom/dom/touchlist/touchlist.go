package touchlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/touch"
	"github.com/matthewmueller/golly/js"
)

type TouchList struct {
}

func (*TouchList) Item(index uint) (t *touch.Touch) {
	js.Rewrite("$<.item($1)", index)
	return t
}

func (*TouchList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
