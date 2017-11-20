package clientrectlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/clientrect"
	"github.com/matthewmueller/golly/js"
)

type ClientRectList struct {
}

func (*ClientRectList) Item(index uint) (c *clientrect.ClientRect) {
	js.Rewrite("$<.item($1)", index)
	return c
}

func (*ClientRectList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
