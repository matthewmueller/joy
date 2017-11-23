package msappasyncoperation

import (
	"github.com/matthewmueller/golly/dom2/domerror"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// MSAppAsyncOperation struct
// js:"MSAppAsyncOperation,omit"
type MSAppAsyncOperation struct {
	window.EventTarget
}

// Start fn
func (*MSAppAsyncOperation) Start() {
	js.Rewrite("$<.start()")
}

// Error prop
func (*MSAppAsyncOperation) Error() (err *domerror.DOMError) {
	js.Rewrite("$<.error")
	return err
}

// Oncomplete prop
func (*MSAppAsyncOperation) Oncomplete() (oncomplete func(window.Event)) {
	js.Rewrite("$<.oncomplete")
	return oncomplete
}

// Oncomplete prop
func (*MSAppAsyncOperation) SetOncomplete(oncomplete func(window.Event)) {
	js.Rewrite("$<.oncomplete = $1", oncomplete)
}

// Onerror prop
func (*MSAppAsyncOperation) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*MSAppAsyncOperation) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// ReadyState prop
func (*MSAppAsyncOperation) ReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

// Result prop
func (*MSAppAsyncOperation) Result() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}
