package window

import (
	"github.com/matthewmueller/golly/dom/document"
	"github.com/matthewmueller/golly/dom/event"
	"github.com/matthewmueller/golly/dom/eventtarget"
	"github.com/matthewmueller/golly/js"
)

var _ eventtarget.EventTarget = (*Window)(nil)

// Window struct
// js:"Window,omit"
type Window struct {
}

// AddEventListener fn
// js:"addEventListener"
func (*Window) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*Window) DispatchEvent(evt event.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// Document prop
// js:"document"
func (*Window) Document() (document document.Document) {
	js.Rewrite("$_.document")
	return document
}
