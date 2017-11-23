package window

import "github.com/matthewmueller/golly/js"

// MessagePort struct
// js:"MessagePort,omit"
type MessagePort struct {
	EventTarget
}

// Close fn
func (*MessagePort) Close() {
	js.Rewrite("$<.close()")
}

// PostMessage fn
func (*MessagePort) PostMessage(message *interface{}, transfer *[]interface{}) {
	js.Rewrite("$<.postMessage($1, $2)", message, transfer)
}

// Start fn
func (*MessagePort) Start() {
	js.Rewrite("$<.start()")
}

// Onmessage prop
func (*MessagePort) Onmessage() (onmessage func(*MessageEvent)) {
	js.Rewrite("$<.onmessage")
	return onmessage
}

// Onmessage prop
func (*MessagePort) SetOnmessage(onmessage func(*MessageEvent)) {
	js.Rewrite("$<.onmessage = $1", onmessage)
}
