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

// Start
func (*MSAppAsyncOperation) Start() {
	js.Rewrite("$<.Start()")
}

// Error
func (*MSAppAsyncOperation) Error() (err *domerror.DOMError) {
	js.Rewrite("$<.Error")
	return err
}

// Oncomplete
func (*MSAppAsyncOperation) Oncomplete() (oncomplete func(window.Event)) {
	js.Rewrite("$<.Oncomplete")
	return oncomplete
}

// Oncomplete
func (*MSAppAsyncOperation) SetOncomplete(oncomplete func(window.Event)) {
	js.Rewrite("$<.Oncomplete = $1", oncomplete)
}

// Onerror
func (*MSAppAsyncOperation) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*MSAppAsyncOperation) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// ReadyState
func (*MSAppAsyncOperation) ReadyState() (readyState uint8) {
	js.Rewrite("$<.ReadyState")
	return readyState
}

// Result
func (*MSAppAsyncOperation) Result() (result interface{}) {
	js.Rewrite("$<.Result")
	return result
}
