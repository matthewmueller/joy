package msstreamreader

import (
	"github.com/matthewmueller/golly/dom/domerror"
	"github.com/matthewmueller/golly/dom/msstream"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New() *MSStreamReader {
	js.Rewrite("MSStreamReader")
	return &MSStreamReader{}
}

// MSStreamReader struct
// js:"MSStreamReader,omit"
type MSStreamReader struct {
	window.EventTarget
}

// Abort fn
func (*MSStreamReader) Abort() {
	js.Rewrite("$<.abort()")
}

// ReadAsArrayBuffer fn
func (*MSStreamReader) ReadAsArrayBuffer(stream *msstream.MSStream, size *int64) {
	js.Rewrite("$<.readAsArrayBuffer($1, $2)", stream, size)
}

// ReadAsBinaryString fn
func (*MSStreamReader) ReadAsBinaryString(stream *msstream.MSStream, size *int64) {
	js.Rewrite("$<.readAsBinaryString($1, $2)", stream, size)
}

// ReadAsBlob fn
func (*MSStreamReader) ReadAsBlob(stream *msstream.MSStream, size *int64) {
	js.Rewrite("$<.readAsBlob($1, $2)", stream, size)
}

// ReadAsDataURL fn
func (*MSStreamReader) ReadAsDataURL(stream *msstream.MSStream, size *int64) {
	js.Rewrite("$<.readAsDataURL($1, $2)", stream, size)
}

// ReadAsText fn
func (*MSStreamReader) ReadAsText(stream *msstream.MSStream, encoding *string, size *int64) {
	js.Rewrite("$<.readAsText($1, $2, $3)", stream, encoding, size)
}

// Onabort prop
func (*MSStreamReader) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$<.onabort")
	return onabort
}

// Onabort prop
func (*MSStreamReader) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$<.onabort = $1", onabort)
}

// Onerror prop
func (*MSStreamReader) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*MSStreamReader) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onload prop
func (*MSStreamReader) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// Onload prop
func (*MSStreamReader) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.onload = $1", onload)
}

// Onloadend prop
func (*MSStreamReader) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$<.onloadend")
	return onloadend
}

// Onloadend prop
func (*MSStreamReader) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$<.onloadend = $1", onloadend)
}

// Onloadstart prop
func (*MSStreamReader) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$<.onloadstart")
	return onloadstart
}

// Onloadstart prop
func (*MSStreamReader) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$<.onloadstart = $1", onloadstart)
}

// Onprogress prop
func (*MSStreamReader) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$<.onprogress")
	return onprogress
}

// Onprogress prop
func (*MSStreamReader) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$<.onprogress = $1", onprogress)
}

// ReadyState prop
func (*MSStreamReader) ReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

// Result prop
func (*MSStreamReader) Result() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}

// Error prop
func (*MSStreamReader) Error() (err *domerror.DOMError) {
	js.Rewrite("$<.error")
	return err
}
