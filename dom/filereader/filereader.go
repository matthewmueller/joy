package filereader

import (
	"github.com/matthewmueller/golly/dom/blob"
	"github.com/matthewmueller/golly/dom/domerror"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.EventTarget = (*FileReader)(nil)

// New fn
func New() *FileReader {
	js.Rewrite("FileReader")
	return &FileReader{}
}

// FileReader struct
// js:"FileReader,omit"
type FileReader struct {
}

// ReadAsArrayBuffer fn
// js:"readAsArrayBuffer"
func (*FileReader) ReadAsArrayBuffer(blob blob.Blob) {
	js.Rewrite("$_.readAsArrayBuffer($1)", blob)
}

// ReadAsBinaryString fn
// js:"readAsBinaryString"
func (*FileReader) ReadAsBinaryString(blob blob.Blob) {
	js.Rewrite("$_.readAsBinaryString($1)", blob)
}

// ReadAsDataURL fn
// js:"readAsDataURL"
func (*FileReader) ReadAsDataURL(blob blob.Blob) {
	js.Rewrite("$_.readAsDataURL($1)", blob)
}

// ReadAsText fn
// js:"readAsText"
func (*FileReader) ReadAsText(blob blob.Blob, encoding *string) {
	js.Rewrite("$_.readAsText($1, $2)", blob, encoding)
}

// Abort fn
// js:"abort"
func (*FileReader) Abort() {
	js.Rewrite("$_.abort()")
}

// AddEventListener fn
// js:"addEventListener"
func (*FileReader) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*FileReader) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*FileReader) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Error prop
// js:"error"
func (*FileReader) Error() (err *domerror.DOMError) {
	js.Rewrite("$_.error")
	return err
}

// Onabort prop
// js:"onabort"
func (*FileReader) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*FileReader) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$_.onabort = $1", onabort)
}

// Onerror prop
// js:"onerror"
func (*FileReader) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*FileReader) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onload prop
// js:"onload"
func (*FileReader) Onload() (onload func(window.Event)) {
	js.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*FileReader) SetOnload(onload func(window.Event)) {
	js.Rewrite("$_.onload = $1", onload)
}

// Onloadend prop
// js:"onloadend"
func (*FileReader) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$_.onloadend")
	return onloadend
}

// SetOnloadend prop
// js:"onloadend"
func (*FileReader) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$_.onloadend = $1", onloadend)
}

// Onloadstart prop
// js:"onloadstart"
func (*FileReader) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*FileReader) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onprogress prop
// js:"onprogress"
func (*FileReader) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*FileReader) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress = $1", onprogress)
}

// ReadyState prop
// js:"readyState"
func (*FileReader) ReadyState() (readyState uint8) {
	js.Rewrite("$_.readyState")
	return readyState
}

// Result prop
// js:"result"
func (*FileReader) Result() (result interface{}) {
	js.Rewrite("$_.result")
	return result
}
