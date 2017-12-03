package mswebviewelement

import (
	"github.com/matthewmueller/joy/dom/domerror"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.EventTarget = (*MSWebViewAsyncOperation)(nil)

// MSWebViewAsyncOperation struct
// js:"MSWebViewAsyncOperation,omit"
type MSWebViewAsyncOperation struct {
}

// Start fn
// js:"start"
func (*MSWebViewAsyncOperation) Start() {
	macro.Rewrite("$_.start()")
}

// AddEventListener fn
// js:"addEventListener"
func (*MSWebViewAsyncOperation) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MSWebViewAsyncOperation) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MSWebViewAsyncOperation) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Error prop
// js:"error"
func (*MSWebViewAsyncOperation) Error() (err *domerror.DOMError) {
	macro.Rewrite("$_.error")
	return err
}

// Oncomplete prop
// js:"oncomplete"
func (*MSWebViewAsyncOperation) Oncomplete() (oncomplete func(window.Event)) {
	macro.Rewrite("$_.oncomplete")
	return oncomplete
}

// SetOncomplete prop
// js:"oncomplete"
func (*MSWebViewAsyncOperation) SetOncomplete(oncomplete func(window.Event)) {
	macro.Rewrite("$_.oncomplete = $1", oncomplete)
}

// Onerror prop
// js:"onerror"
func (*MSWebViewAsyncOperation) Onerror() (onerror func(window.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*MSWebViewAsyncOperation) SetOnerror(onerror func(window.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// ReadyState prop
// js:"readyState"
func (*MSWebViewAsyncOperation) ReadyState() (readyState uint8) {
	macro.Rewrite("$_.readyState")
	return readyState
}

// Result prop
// js:"result"
func (*MSWebViewAsyncOperation) Result() (result interface{}) {
	macro.Rewrite("$_.result")
	return result
}

// Target prop
// js:"target"
func (*MSWebViewAsyncOperation) Target() (target *MSHTMLWebViewElement) {
	macro.Rewrite("$_.target")
	return target
}

// Type prop
// js:"type"
func (*MSWebViewAsyncOperation) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
