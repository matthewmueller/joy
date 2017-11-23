package msbasereader

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// MSBaseReader struct
// js:"MSBaseReader,omit"
type MSBaseReader struct {
}

// Abort
func (*MSBaseReader) Abort() {
	js.Rewrite("$<.Abort()")
}

// Onabort
func (*MSBaseReader) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$<.Onabort")
	return onabort
}

// Onabort
func (*MSBaseReader) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$<.Onabort = $1", onabort)
}

// Onerror
func (*MSBaseReader) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*MSBaseReader) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onload
func (*MSBaseReader) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.Onload")
	return onload
}

// Onload
func (*MSBaseReader) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.Onload = $1", onload)
}

// Onloadend
func (*MSBaseReader) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$<.Onloadend")
	return onloadend
}

// Onloadend
func (*MSBaseReader) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$<.Onloadend = $1", onloadend)
}

// Onloadstart
func (*MSBaseReader) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$<.Onloadstart")
	return onloadstart
}

// Onloadstart
func (*MSBaseReader) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$<.Onloadstart = $1", onloadstart)
}

// Onprogress
func (*MSBaseReader) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$<.Onprogress")
	return onprogress
}

// Onprogress
func (*MSBaseReader) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$<.Onprogress = $1", onprogress)
}

// ReadyState
func (*MSBaseReader) ReadyState() (readyState uint8) {
	js.Rewrite("$<.ReadyState")
	return readyState
}

// Result
func (*MSBaseReader) Result() (result interface{}) {
	js.Rewrite("$<.Result")
	return result
}
