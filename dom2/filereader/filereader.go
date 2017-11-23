package filereader

import (
	"github.com/matthewmueller/golly/dom2/blob"
	"github.com/matthewmueller/golly/dom2/domerror"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New() *FileReader {
	js.Rewrite("FileReader")
	return &FileReader{}
}

// FileReader struct
// js:"FileReader,omit"
type FileReader struct {
	window.EventTarget
}

// Abort
func (*FileReader) Abort() {
	js.Rewrite("$<.Abort()")
}

// ReadAsArrayBuffer
func (*FileReader) ReadAsArrayBuffer(blob blob.Blob) {
	js.Rewrite("$<.ReadAsArrayBuffer($1)", blob)
}

// ReadAsBinaryString
func (*FileReader) ReadAsBinaryString(blob blob.Blob) {
	js.Rewrite("$<.ReadAsBinaryString($1)", blob)
}

// ReadAsDataURL
func (*FileReader) ReadAsDataURL(blob blob.Blob) {
	js.Rewrite("$<.ReadAsDataURL($1)", blob)
}

// ReadAsText
func (*FileReader) ReadAsText(blob blob.Blob, encoding *string) {
	js.Rewrite("$<.ReadAsText($1, $2)", blob, encoding)
}

// Onabort
func (*FileReader) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$<.Onabort")
	return onabort
}

// Onabort
func (*FileReader) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$<.Onabort = $1", onabort)
}

// Onerror
func (*FileReader) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*FileReader) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onload
func (*FileReader) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.Onload")
	return onload
}

// Onload
func (*FileReader) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.Onload = $1", onload)
}

// Onloadend
func (*FileReader) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$<.Onloadend")
	return onloadend
}

// Onloadend
func (*FileReader) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$<.Onloadend = $1", onloadend)
}

// Onloadstart
func (*FileReader) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$<.Onloadstart")
	return onloadstart
}

// Onloadstart
func (*FileReader) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$<.Onloadstart = $1", onloadstart)
}

// Onprogress
func (*FileReader) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$<.Onprogress")
	return onprogress
}

// Onprogress
func (*FileReader) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$<.Onprogress = $1", onprogress)
}

// ReadyState
func (*FileReader) ReadyState() (readyState uint8) {
	js.Rewrite("$<.ReadyState")
	return readyState
}

// Result
func (*FileReader) Result() (result interface{}) {
	js.Rewrite("$<.Result")
	return result
}

// Error
func (*FileReader) Error() (err *domerror.DOMError) {
	js.Rewrite("$<.Error")
	return err
}
