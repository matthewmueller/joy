package clientrectlist

import (
	"github.com/matthewmueller/joy/dom/clientrect"
	"github.com/matthewmueller/joy/macro"
)

// ClientRectList struct
// js:"ClientRectList,omit"
type ClientRectList struct {
}

// Item fn
// js:"item"
func (*ClientRectList) Item(index uint) (c *clientrect.ClientRect) {
	macro.Rewrite("$_.item($1)", index)
	return c
}

// Length prop
// js:"length"
func (*ClientRectList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
