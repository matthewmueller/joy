package worker

import (
	"github.com/matthewmueller/golly/dom/window"
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

// PostMessage fn
func (*Worker) PostMessage(message interface{}, transfer *[]interface{}) {
	js.Rewrite("$<.postMessage($1, $2)", message, transfer)
}

// Terminate fn
func (*Worker) Terminate() {
	js.Rewrite("$<.terminate()")
}

// Onerror prop
func (*Worker) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*Worker) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onmessage prop
func (*Worker) Onmessage() (onmessage func(*window.MessageEvent)) {
	js.Rewrite("$<.onmessage")
	return onmessage
}

// Onmessage prop
func (*Worker) SetOnmessage(onmessage func(*window.MessageEvent)) {
	js.Rewrite("$<.onmessage = $1", onmessage)
}
