package clientrectlist

import (
	"github.com/matthewmueller/golly/dom2/clientrect"
	"github.com/matthewmueller/golly/js"
)

// ClientRectList struct
// js:"ClientRectList,omit"
type ClientRectList struct {
}

// Item fn
func (*ClientRectList) Item(index uint) (c *clientrect.ClientRect) {
	js.Rewrite("$<.item($1)", index)
	return c
}

// Length prop
func (*ClientRectList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
