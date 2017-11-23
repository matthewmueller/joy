package window

import "github.com/matthewmueller/golly/js"

// NodeList struct
// js:"NodeList,omit"
type NodeList struct {
}

// Item
func (*NodeList) Item(index uint) (n Node) {
	js.Rewrite("$<.Item($1)", index)
	return n
}

// Length
func (*NodeList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
