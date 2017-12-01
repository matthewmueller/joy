package xmlhttprequestupload

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.EventTarget = (*XMLHTTPRequestUpload)(nil)

// XMLHTTPRequestUpload struct
// js:"XMLHTTPRequestUpload,omit"
type XMLHTTPRequestUpload struct {
}

// AddEventListener fn
// js:"addEventListener"
func (*XMLHTTPRequestUpload) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*XMLHTTPRequestUpload) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*XMLHTTPRequestUpload) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onabort prop
// js:"onabort"
func (*XMLHTTPRequestUpload) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*XMLHTTPRequestUpload) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$_.onabort = $1", onabort)
}

// Onerror prop
// js:"onerror"
func (*XMLHTTPRequestUpload) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*XMLHTTPRequestUpload) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onload prop
// js:"onload"
func (*XMLHTTPRequestUpload) Onload() (onload func(window.Event)) {
	js.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*XMLHTTPRequestUpload) SetOnload(onload func(window.Event)) {
	js.Rewrite("$_.onload = $1", onload)
}

// Onloadend prop
// js:"onloadend"
func (*XMLHTTPRequestUpload) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$_.onloadend")
	return onloadend
}

// SetOnloadend prop
// js:"onloadend"
func (*XMLHTTPRequestUpload) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$_.onloadend = $1", onloadend)
}

// Onloadstart prop
// js:"onloadstart"
func (*XMLHTTPRequestUpload) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*XMLHTTPRequestUpload) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onprogress prop
// js:"onprogress"
func (*XMLHTTPRequestUpload) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*XMLHTTPRequestUpload) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress = $1", onprogress)
}

// Ontimeout prop
// js:"ontimeout"
func (*XMLHTTPRequestUpload) Ontimeout() (ontimeout func(window.Event)) {
	js.Rewrite("$_.ontimeout")
	return ontimeout
}

// SetOntimeout prop
// js:"ontimeout"
func (*XMLHTTPRequestUpload) SetOntimeout(ontimeout func(window.Event)) {
	js.Rewrite("$_.ontimeout = $1", ontimeout)
}
