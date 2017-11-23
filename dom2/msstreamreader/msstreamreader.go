package msstreamreader

import (
	"github.com/matthewmueller/golly/dom2/domerror"
	"github.com/matthewmueller/golly/dom2/msstream"
	"github.com/matthewmueller/golly/dom2/window"
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

// Abort
func (*MSStreamReader) Abort() {
	js.Rewrite("$<.Abort()")
}

// ReadAsArrayBuffer
func (*MSStreamReader) ReadAsArrayBuffer(stream *msstream.MSStream, size *int64) {
	js.Rewrite("$<.ReadAsArrayBuffer($1, $2)", stream, size)
}

// ReadAsBinaryString
func (*MSStreamReader) ReadAsBinaryString(stream *msstream.MSStream, size *int64) {
	js.Rewrite("$<.ReadAsBinaryString($1, $2)", stream, size)
}

// ReadAsBlob
func (*MSStreamReader) ReadAsBlob(stream *msstream.MSStream, size *int64) {
	js.Rewrite("$<.ReadAsBlob($1, $2)", stream, size)
}

// ReadAsDataURL
func (*MSStreamReader) ReadAsDataURL(stream *msstream.MSStream, size *int64) {
	js.Rewrite("$<.ReadAsDataURL($1, $2)", stream, size)
}

// ReadAsText
func (*MSStreamReader) ReadAsText(stream *msstream.MSStream, encoding *string, size *int64) {
	js.Rewrite("$<.ReadAsText($1, $2, $3)", stream, encoding, size)
}

// Onabort
func (*MSStreamReader) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$<.Onabort")
	return onabort
}

// Onabort
func (*MSStreamReader) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$<.Onabort = $1", onabort)
}

// Onerror
func (*MSStreamReader) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*MSStreamReader) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onload
func (*MSStreamReader) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.Onload")
	return onload
}

// Onload
func (*MSStreamReader) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.Onload = $1", onload)
}

// Onloadend
func (*MSStreamReader) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$<.Onloadend")
	return onloadend
}

// Onloadend
func (*MSStreamReader) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$<.Onloadend = $1", onloadend)
}

// Onloadstart
func (*MSStreamReader) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$<.Onloadstart")
	return onloadstart
}

// Onloadstart
func (*MSStreamReader) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$<.Onloadstart = $1", onloadstart)
}

// Onprogress
func (*MSStreamReader) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$<.Onprogress")
	return onprogress
}

// Onprogress
func (*MSStreamReader) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$<.Onprogress = $1", onprogress)
}

// ReadyState
func (*MSStreamReader) ReadyState() (readyState uint8) {
	js.Rewrite("$<.ReadyState")
	return readyState
}

// Result
func (*MSStreamReader) Result() (result interface{}) {
	js.Rewrite("$<.Result")
	return result
}

// Error
func (*MSStreamReader) Error() (err *domerror.DOMError) {
	js.Rewrite("$<.Error")
	return err
}
