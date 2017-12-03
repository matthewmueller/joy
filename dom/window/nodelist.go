package window

import "github.com/matthewmueller/joy/macro"

// NodeList struct
// js:"NodeList,omit"
type NodeList struct {
}

// Item fn
// js:"item"
func (*NodeList) Item(index uint) (n Node) {
	macro.Rewrite("$_.item($1)", index)
	return n
}

// Length prop
// js:"length"
func (*NodeList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
