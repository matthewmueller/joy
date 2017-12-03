package sourcebufferlist

import (
	"github.com/matthewmueller/joy/dom/avtrack"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.EventTarget = (*SourceBufferList)(nil)

// SourceBufferList struct
// js:"SourceBufferList,omit"
type SourceBufferList struct {
}

// Item fn
// js:"item"
func (*SourceBufferList) Item(index uint) (a *avtrack.SourceBuffer) {
	js.Rewrite("$_.item($1)", index)
	return a
}

// AddEventListener fn
// js:"addEventListener"
func (*SourceBufferList) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SourceBufferList) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SourceBufferList) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Length prop
// js:"length"
func (*SourceBufferList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
