package mswebviewelement

import (
	"github.com/matthewmueller/golly/dom2/domerror"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// MSWebViewAsyncOperation struct
// js:"MSWebViewAsyncOperation,omit"
type MSWebViewAsyncOperation struct {
	window.EventTarget
}

// Start fn
func (*MSWebViewAsyncOperation) Start() {
	js.Rewrite("$<.start()")
}

// Error prop
func (*MSWebViewAsyncOperation) Error() (err *domerror.DOMError) {
	js.Rewrite("$<.error")
	return err
}

// Oncomplete prop
func (*MSWebViewAsyncOperation) Oncomplete() (oncomplete func(window.Event)) {
	js.Rewrite("$<.oncomplete")
	return oncomplete
}

// Oncomplete prop
func (*MSWebViewAsyncOperation) SetOncomplete(oncomplete func(window.Event)) {
	js.Rewrite("$<.oncomplete = $1", oncomplete)
}

// Onerror prop
func (*MSWebViewAsyncOperation) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*MSWebViewAsyncOperation) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// ReadyState prop
func (*MSWebViewAsyncOperation) ReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

// Result prop
func (*MSWebViewAsyncOperation) Result() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}

// Target prop
func (*MSWebViewAsyncOperation) Target() (target *MSHTMLWebViewElement) {
	js.Rewrite("$<.target")
	return target
}

// Type prop
func (*MSWebViewAsyncOperation) Type() (kind uint8) {
	js.Rewrite("$<.type")
	return kind
}
