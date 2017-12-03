package sourcebufferlist

import (
	"github.com/matthewmueller/joy/dom/avtrack"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.EventTarget = (*SourceBufferList)(nil)

// SourceBufferList struct
// js:"SourceBufferList,omit"
type SourceBufferList struct {
}

// Item fn
// js:"item"
func (*SourceBufferList) Item(index uint) (a *avtrack.SourceBuffer) {
	macro.Rewrite("$_.item($1)", index)
	return a
}

// AddEventListener fn
// js:"addEventListener"
func (*SourceBufferList) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SourceBufferList) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SourceBufferList) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Length prop
// js:"length"
func (*SourceBufferList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
