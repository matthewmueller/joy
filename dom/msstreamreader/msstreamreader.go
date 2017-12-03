package msstreamreader

import (
	"github.com/matthewmueller/joy/dom/domerror"
	"github.com/matthewmueller/joy/dom/msstream"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.EventTarget = (*MSStreamReader)(nil)

// New fn
func New() *MSStreamReader {
	macro.Rewrite("MSStreamReader")
	return &MSStreamReader{}
}

// MSStreamReader struct
// js:"MSStreamReader,omit"
type MSStreamReader struct {
}

// ReadAsArrayBuffer fn
// js:"readAsArrayBuffer"
func (*MSStreamReader) ReadAsArrayBuffer(stream *msstream.MSStream, size *int64) {
	macro.Rewrite("$_.readAsArrayBuffer($1, $2)", stream, size)
}

// ReadAsBinaryString fn
// js:"readAsBinaryString"
func (*MSStreamReader) ReadAsBinaryString(stream *msstream.MSStream, size *int64) {
	macro.Rewrite("$_.readAsBinaryString($1, $2)", stream, size)
}

// ReadAsBlob fn
// js:"readAsBlob"
func (*MSStreamReader) ReadAsBlob(stream *msstream.MSStream, size *int64) {
	macro.Rewrite("$_.readAsBlob($1, $2)", stream, size)
}

// ReadAsDataURL fn
// js:"readAsDataURL"
func (*MSStreamReader) ReadAsDataURL(stream *msstream.MSStream, size *int64) {
	macro.Rewrite("$_.readAsDataURL($1, $2)", stream, size)
}

// ReadAsText fn
// js:"readAsText"
func (*MSStreamReader) ReadAsText(stream *msstream.MSStream, encoding *string, size *int64) {
	macro.Rewrite("$_.readAsText($1, $2, $3)", stream, encoding, size)
}

// Abort fn
// js:"abort"
func (*MSStreamReader) Abort() {
	macro.Rewrite("$_.abort()")
}

// AddEventListener fn
// js:"addEventListener"
func (*MSStreamReader) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MSStreamReader) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MSStreamReader) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Error prop
// js:"error"
func (*MSStreamReader) Error() (err *domerror.DOMError) {
	macro.Rewrite("$_.error")
	return err
}

// Onabort prop
// js:"onabort"
func (*MSStreamReader) Onabort() (onabort func(window.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*MSStreamReader) SetOnabort(onabort func(window.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

// Onerror prop
// js:"onerror"
func (*MSStreamReader) Onerror() (onerror func(window.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*MSStreamReader) SetOnerror(onerror func(window.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onload prop
// js:"onload"
func (*MSStreamReader) Onload() (onload func(window.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*MSStreamReader) SetOnload(onload func(window.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

// Onloadend prop
// js:"onloadend"
func (*MSStreamReader) Onloadend() (onloadend func(window.Event)) {
	macro.Rewrite("$_.onloadend")
	return onloadend
}

// SetOnloadend prop
// js:"onloadend"
func (*MSStreamReader) SetOnloadend(onloadend func(window.Event)) {
	macro.Rewrite("$_.onloadend = $1", onloadend)
}

// Onloadstart prop
// js:"onloadstart"
func (*MSStreamReader) Onloadstart() (onloadstart func(window.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*MSStreamReader) SetOnloadstart(onloadstart func(window.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onprogress prop
// js:"onprogress"
func (*MSStreamReader) Onprogress() (onprogress func(window.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*MSStreamReader) SetOnprogress(onprogress func(window.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

// ReadyState prop
// js:"readyState"
func (*MSStreamReader) ReadyState() (readyState uint8) {
	macro.Rewrite("$_.readyState")
	return readyState
}

// Result prop
// js:"result"
func (*MSStreamReader) Result() (result interface{}) {
	macro.Rewrite("$_.result")
	return result
}
