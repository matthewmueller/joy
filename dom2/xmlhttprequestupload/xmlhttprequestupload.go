package xmlhttprequestupload

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// XMLHTTPRequestUpload struct
// js:"XMLHTTPRequestUpload,omit"
type XMLHTTPRequestUpload struct {
	window.EventTarget
}

// Onabort
func (*XMLHttpRequestUpload) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$<.Onabort")
	return onabort
}

// Onabort
func (*XMLHttpRequestUpload) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$<.Onabort = $1", onabort)
}

// Onerror
func (*XMLHttpRequestUpload) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*XMLHttpRequestUpload) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onload
func (*XMLHttpRequestUpload) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.Onload")
	return onload
}

// Onload
func (*XMLHttpRequestUpload) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.Onload = $1", onload)
}

// Onloadend
func (*XMLHttpRequestUpload) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$<.Onloadend")
	return onloadend
}

// Onloadend
func (*XMLHttpRequestUpload) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$<.Onloadend = $1", onloadend)
}

// Onloadstart
func (*XMLHttpRequestUpload) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$<.Onloadstart")
	return onloadstart
}

// Onloadstart
func (*XMLHttpRequestUpload) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$<.Onloadstart = $1", onloadstart)
}

// Onprogress
func (*XMLHttpRequestUpload) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$<.Onprogress")
	return onprogress
}

// Onprogress
func (*XMLHttpRequestUpload) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$<.Onprogress = $1", onprogress)
}

// Ontimeout
func (*XMLHttpRequestUpload) Ontimeout() (ontimeout func(window.Event)) {
	js.Rewrite("$<.Ontimeout")
	return ontimeout
}

// Ontimeout
func (*XMLHttpRequestUpload) SetOntimeout(ontimeout func(window.Event)) {
	js.Rewrite("$<.Ontimeout = $1", ontimeout)
}
