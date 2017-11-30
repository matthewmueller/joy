package worker

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.EventTarget = (*Worker)(nil)

// Worker struct
// js:"Worker,omit"
type Worker struct {
}

// PostMessage fn
// js:"postMessage"
func (*Worker) PostMessage(message interface{}, transfer *[]interface{}) {
	js.Rewrite("$_.postMessage($1, $2)", message, transfer)
}

// Terminate fn
// js:"terminate"
func (*Worker) Terminate() {
	js.Rewrite("$_.terminate()")
}

// AddEventListener fn
// js:"addEventListener"
func (*Worker) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*Worker) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*Worker) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onmessage prop
// js:"onmessage"
func (*Worker) Onmessage() (onmessage func(*window.MessageEvent)) {
	js.Rewrite("$_.onmessage")
	return onmessage
}

// SetOnmessage prop
// js:"onmessage"
func (*Worker) SetOnmessage(onmessage func(*window.MessageEvent)) {
	js.Rewrite("$_.onmessage = $1", onmessage)
}

// Onerror prop
// js:"onerror"
func (*Worker) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*Worker) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}
