package window

import "github.com/matthewmueller/golly/js"

// NodeList struct
// js:"NodeList,omit"
type NodeList struct {
}

// Item fn
func (*NodeList) Item(index uint) (n Node) {
	js.Rewrite("$<.item($1)", index)
	return n
}

// Length prop
func (*NodeList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
