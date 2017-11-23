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

// Start
func (*MSWebViewAsyncOperation) Start() {
	js.Rewrite("$<.Start()")
}

// Error
func (*MSWebViewAsyncOperation) Error() (err *domerror.DOMError) {
	js.Rewrite("$<.Error")
	return err
}

// Oncomplete
func (*MSWebViewAsyncOperation) Oncomplete() (oncomplete func(window.Event)) {
	js.Rewrite("$<.Oncomplete")
	return oncomplete
}

// Oncomplete
func (*MSWebViewAsyncOperation) SetOncomplete(oncomplete func(window.Event)) {
	js.Rewrite("$<.Oncomplete = $1", oncomplete)
}

// Onerror
func (*MSWebViewAsyncOperation) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*MSWebViewAsyncOperation) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// ReadyState
func (*MSWebViewAsyncOperation) ReadyState() (readyState uint8) {
	js.Rewrite("$<.ReadyState")
	return readyState
}

// Result
func (*MSWebViewAsyncOperation) Result() (result interface{}) {
	js.Rewrite("$<.Result")
	return result
}

// Target
func (*MSWebViewAsyncOperation) Target() (target *MSHTMLWebViewElement) {
	js.Rewrite("$<.Target")
	return target
}

// Type
func (*MSWebViewAsyncOperation) Type() (kind uint8) {
	js.Rewrite("$<.Type")
	return kind
}
