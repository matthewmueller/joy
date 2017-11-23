package worker

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"Worker,omit"
type Worker struct {
	window.EventTarget
	window.AbstractWorker
}

// PostMessage
func (*Worker) PostMessage(message interface{}, transfer *[]interface{}) {
	js.Rewrite("$<.PostMessage($1, $2)", message, transfer)
}

// Terminate
func (*Worker) Terminate() {
	js.Rewrite("$<.Terminate()")
}

// Onmessage
func (*Worker) Onmessage() (onmessage func(*window.MessageEvent)) {
	js.Rewrite("$<.Onmessage")
	return onmessage
}

// Onmessage
func (*Worker) SetOnmessage(onmessage func(*window.MessageEvent)) {
	js.Rewrite("$<.Onmessage = $1", onmessage)
}
