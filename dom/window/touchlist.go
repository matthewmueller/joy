package window

import "github.com/matthewmueller/joy/macro"

// TouchList struct
// js:"TouchList,omit"
type TouchList struct {
}

// Item fn
// js:"item"
func (*TouchList) Item(index uint) (t *Touch) {
	macro.Rewrite("$_.item($1)", index)
	return t
}

// Length prop
// js:"length"
func (*TouchList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
