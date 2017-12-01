package msappasyncoperation

import (
	"github.com/matthewmueller/golly/dom/domerror"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.EventTarget = (*MSAppAsyncOperation)(nil)

// MSAppAsyncOperation struct
// js:"MSAppAsyncOperation,omit"
type MSAppAsyncOperation struct {
}

// Start fn
// js:"start"
func (*MSAppAsyncOperation) Start() {
	js.Rewrite("$_.start()")
}

// AddEventListener fn
// js:"addEventListener"
func (*MSAppAsyncOperation) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MSAppAsyncOperation) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MSAppAsyncOperation) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Error prop
// js:"error"
func (*MSAppAsyncOperation) Error() (err *domerror.DOMError) {
	js.Rewrite("$_.error")
	return err
}

// Oncomplete prop
// js:"oncomplete"
func (*MSAppAsyncOperation) Oncomplete() (oncomplete func(window.Event)) {
	js.Rewrite("$_.oncomplete")
	return oncomplete
}

// SetOncomplete prop
// js:"oncomplete"
func (*MSAppAsyncOperation) SetOncomplete(oncomplete func(window.Event)) {
	js.Rewrite("$_.oncomplete = $1", oncomplete)
}

// Onerror prop
// js:"onerror"
func (*MSAppAsyncOperation) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*MSAppAsyncOperation) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// ReadyState prop
// js:"readyState"
func (*MSAppAsyncOperation) ReadyState() (readyState uint8) {
	js.Rewrite("$_.readyState")
	return readyState
}

// Result prop
// js:"result"
func (*MSAppAsyncOperation) Result() (result interface{}) {
	js.Rewrite("$_.result")
	return result
}
