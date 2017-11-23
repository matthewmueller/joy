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

// Onabort prop
func (*XMLHttpRequestUpload) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$<.onabort")
	return onabort
}

// Onabort prop
func (*XMLHttpRequestUpload) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$<.onabort = $1", onabort)
}

// Onerror prop
func (*XMLHttpRequestUpload) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*XMLHttpRequestUpload) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onload prop
func (*XMLHttpRequestUpload) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// Onload prop
func (*XMLHttpRequestUpload) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.onload = $1", onload)
}

// Onloadend prop
func (*XMLHttpRequestUpload) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$<.onloadend")
	return onloadend
}

// Onloadend prop
func (*XMLHttpRequestUpload) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$<.onloadend = $1", onloadend)
}

// Onloadstart prop
func (*XMLHttpRequestUpload) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$<.onloadstart")
	return onloadstart
}

// Onloadstart prop
func (*XMLHttpRequestUpload) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$<.onloadstart = $1", onloadstart)
}

// Onprogress prop
func (*XMLHttpRequestUpload) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$<.onprogress")
	return onprogress
}

// Onprogress prop
func (*XMLHttpRequestUpload) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$<.onprogress = $1", onprogress)
}

// Ontimeout prop
func (*XMLHttpRequestUpload) Ontimeout() (ontimeout func(window.Event)) {
	js.Rewrite("$<.ontimeout")
	return ontimeout
}

// Ontimeout prop
func (*XMLHttpRequestUpload) SetOntimeout(ontimeout func(window.Event)) {
	js.Rewrite("$<.ontimeout = $1", ontimeout)
}
