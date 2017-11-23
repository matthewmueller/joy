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

// Item
func (*SourceBufferList) Item(index uint) (a *avtrack.SourceBuffer) {
	js.Rewrite("$<.Item($1)", index)
	return a
}

// Length
func (*SourceBufferList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
