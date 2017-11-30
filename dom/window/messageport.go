package window

import "github.com/matthewmueller/golly/js"

var _ EventTarget = (*MessagePort)(nil)

// MessagePort struct
// js:"MessagePort,omit"
type MessagePort struct {
}

// Close fn
// js:"close"
func (*MessagePort) Close() {
	js.Rewrite("$_.close()")
}

// PostMessage fn
// js:"postMessage"
func (*MessagePort) PostMessage(message *interface{}, transfer *[]interface{}) {
	js.Rewrite("$_.postMessage($1, $2)", message, transfer)
}

// Start fn
// js:"start"
func (*MessagePort) Start() {
	js.Rewrite("$_.start()")
}

// AddEventListener fn
// js:"addEventListener"
func (*MessagePort) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MessagePort) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MessagePort) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onmessage prop
// js:"onmessage"
func (*MessagePort) Onmessage() (onmessage func(*MessageEvent)) {
	js.Rewrite("$_.onmessage")
	return onmessage
}

// SetOnmessage prop
// js:"onmessage"
func (*MessagePort) SetOnmessage(onmessage func(*MessageEvent)) {
	js.Rewrite("$_.onmessage = $1", onmessage)
}
