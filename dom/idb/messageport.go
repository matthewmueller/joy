package idb

import "github.com/matthewmueller/golly/js"

// MessagePort struct
// js:"MessagePort,omit"
type MessagePort struct {
	EventTarget
}

// Close
func (*MessagePort) Close() {
	js.Rewrite("$<.Close()")
}

// PostMessage
func (*MessagePort) PostMessage(message *interface{}, transfer *[]interface{}) {
	js.Rewrite("$<.PostMessage($1, $2)", message, transfer)
}

// Start
func (*MessagePort) Start() {
	js.Rewrite("$<.Start()")
}

// Onmessage
func (*MessagePort) Onmessage() (onmessage func(*MessageEvent)) {
	js.Rewrite("$<.Onmessage")
	return onmessage
}

// Onmessage
func (*MessagePort) SetOnmessage(onmessage func(*MessageEvent)) {
	js.Rewrite("$<.Onmessage = $1", onmessage)
}
