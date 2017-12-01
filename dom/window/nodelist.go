package window

import "github.com/matthewmueller/golly/js"

// NodeList struct
// js:"NodeList,omit"
type NodeList struct {
}

// Item fn
// js:"item"
func (*NodeList) Item(index uint) (n Node) {
	js.Rewrite("$_.item($1)", index)
	return n
}

// Length prop
// js:"length"
func (*NodeList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
