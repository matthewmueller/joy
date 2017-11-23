package idb

import "github.com/matthewmueller/golly/js"

// TouchList struct
// js:"TouchList,omit"
type TouchList struct {
}

// Item
func (*TouchList) Item(index uint) (t *Touch) {
	js.Rewrite("$<.Item($1)", index)
	return t
}

// Length
func (*TouchList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
