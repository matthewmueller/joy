package msbasereader

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// MSBaseReader struct
// js:"MSBaseReader,omit"
type MSBaseReader struct {
}

// Abort fn
func (*MSBaseReader) Abort() {
	js.Rewrite("$<.abort()")
}

// Onabort prop
func (*MSBaseReader) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$<.onabort")
	return onabort
}

// Onabort prop
func (*MSBaseReader) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$<.onabort = $1", onabort)
}

// Onerror prop
func (*MSBaseReader) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*MSBaseReader) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onload prop
func (*MSBaseReader) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// Onload prop
func (*MSBaseReader) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.onload = $1", onload)
}

// Onloadend prop
func (*MSBaseReader) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$<.onloadend")
	return onloadend
}

// Onloadend prop
func (*MSBaseReader) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$<.onloadend = $1", onloadend)
}

// Onloadstart prop
func (*MSBaseReader) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$<.onloadstart")
	return onloadstart
}

// Onloadstart prop
func (*MSBaseReader) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$<.onloadstart = $1", onloadstart)
}

// Onprogress prop
func (*MSBaseReader) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$<.onprogress")
	return onprogress
}

// Onprogress prop
func (*MSBaseReader) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$<.onprogress = $1", onprogress)
}

// ReadyState prop
func (*MSBaseReader) ReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

// Result prop
func (*MSBaseReader) Result() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}
