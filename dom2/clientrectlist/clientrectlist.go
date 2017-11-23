package clientrectlist

import (
	"github.com/matthewmueller/golly/dom2/clientrect"
	"github.com/matthewmueller/golly/js"
)

// ClientRectList struct
// js:"ClientRectList,omit"
type ClientRectList struct {
}

// Item
func (*ClientRectList) Item(index uint) (c *clientrect.ClientRect) {
	js.Rewrite("$<.Item($1)", index)
	return c
}

// Length
func (*ClientRectList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
