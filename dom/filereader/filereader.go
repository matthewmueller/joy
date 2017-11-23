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

// Abort fn
func (*FileReader) Abort() {
	js.Rewrite("$<.abort()")
}

// ReadAsArrayBuffer fn
func (*FileReader) ReadAsArrayBuffer(blob blob.Blob) {
	js.Rewrite("$<.readAsArrayBuffer($1)", blob)
}

// ReadAsBinaryString fn
func (*FileReader) ReadAsBinaryString(blob blob.Blob) {
	js.Rewrite("$<.readAsBinaryString($1)", blob)
}

// ReadAsDataURL fn
func (*FileReader) ReadAsDataURL(blob blob.Blob) {
	js.Rewrite("$<.readAsDataURL($1)", blob)
}

// ReadAsText fn
func (*FileReader) ReadAsText(blob blob.Blob, encoding *string) {
	js.Rewrite("$<.readAsText($1, $2)", blob, encoding)
}

// Onabort prop
func (*FileReader) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$<.onabort")
	return onabort
}

// Onabort prop
func (*FileReader) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$<.onabort = $1", onabort)
}

// Onerror prop
func (*FileReader) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*FileReader) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onload prop
func (*FileReader) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// Onload prop
func (*FileReader) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.onload = $1", onload)
}

// Onloadend prop
func (*FileReader) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$<.onloadend")
	return onloadend
}

// Onloadend prop
func (*FileReader) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$<.onloadend = $1", onloadend)
}

// Onloadstart prop
func (*FileReader) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$<.onloadstart")
	return onloadstart
}

// Onloadstart prop
func (*FileReader) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$<.onloadstart = $1", onloadstart)
}

// Onprogress prop
func (*FileReader) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$<.onprogress")
	return onprogress
}

// Onprogress prop
func (*FileReader) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$<.onprogress = $1", onprogress)
}

// ReadyState prop
func (*FileReader) ReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

// Result prop
func (*FileReader) Result() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}

// Error prop
func (*FileReader) Error() (err *domerror.DOMError) {
	js.Rewrite("$<.error")
	return err
}
