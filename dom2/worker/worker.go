package worker

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(stringurl string) *Worker {
	js.Rewrite("Worker")
	return &Worker{}
}

// Worker struct
// js:"Worker,omit"
type Worker struct {
	window.EventTarget
}

// PostMessage
func (*Worker) PostMessage(message interface{}, transfer *[]interface{}) {
	js.Rewrite("$<.PostMessage($1, $2)", message, transfer)
}

// Terminate
func (*Worker) Terminate() {
	js.Rewrite("$<.Terminate()")
}

// Onerror
func (*Worker) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*Worker) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
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
