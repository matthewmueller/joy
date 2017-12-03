package worker

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.EventTarget = (*Worker)(nil)

// New fn
func New(stringurl string) *Worker {
	macro.Rewrite("Worker")
	return &Worker{}
}

// Worker struct
// js:"Worker,omit"
type Worker struct {
}

// PostMessage fn
// js:"postMessage"
func (*Worker) PostMessage(message interface{}, transfer *[]interface{}) {
	macro.Rewrite("$_.postMessage($1, $2)", message, transfer)
}

// Terminate fn
// js:"terminate"
func (*Worker) Terminate() {
	macro.Rewrite("$_.terminate()")
}

// AddEventListener fn
// js:"addEventListener"
func (*Worker) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*Worker) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*Worker) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onmessage prop
// js:"onmessage"
func (*Worker) Onmessage() (onmessage func(*window.MessageEvent)) {
	macro.Rewrite("$_.onmessage")
	return onmessage
}

// SetOnmessage prop
// js:"onmessage"
func (*Worker) SetOnmessage(onmessage func(*window.MessageEvent)) {
	macro.Rewrite("$_.onmessage = $1", onmessage)
}

// Onerror prop
// js:"onerror"
func (*Worker) Onerror() (onerror func(window.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*Worker) SetOnerror(onerror func(window.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}
