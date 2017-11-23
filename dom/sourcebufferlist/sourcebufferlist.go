package sourcebufferlist

import (
	"github.com/matthewmueller/golly/dom2/avtrack"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// SourceBufferList struct
// js:"SourceBufferList,omit"
type SourceBufferList struct {
	window.EventTarget
}

// Item fn
func (*SourceBufferList) Item(index uint) (a *avtrack.SourceBuffer) {
	js.Rewrite("$<.item($1)", index)
	return a
}

// Length prop
func (*SourceBufferList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
