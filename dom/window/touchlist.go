package window

import "github.com/matthewmueller/golly/js"

// TouchList struct
// js:"TouchList,omit"
type TouchList struct {
}

// Item fn
// js:"item"
func (*TouchList) Item(index uint) (t *Touch) {
	js.Rewrite("$_.item($1)", index)
	return t
}

// Length prop
// js:"length"
func (*TouchList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
