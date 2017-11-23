package window

import "github.com/matthewmueller/golly/js"

// TouchList struct
// js:"TouchList,omit"
type TouchList struct {
}

// Item fn
func (*TouchList) Item(index uint) (t *Touch) {
	js.Rewrite("$<.item($1)", index)
	return t
}

// Length prop
func (*TouchList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
