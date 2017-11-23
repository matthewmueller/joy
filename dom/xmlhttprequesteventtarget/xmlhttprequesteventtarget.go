package xmlhttprequesteventtarget

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// XMLHTTPRequestEventTarget struct
// js:"XMLHTTPRequestEventTarget,omit"
type XMLHTTPRequestEventTarget struct {
}

// Onabort prop
func (*XMLHttpRequestEventTarget) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$<.onabort")
	return onabort
}

// Onabort prop
func (*XMLHttpRequestEventTarget) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$<.onabort = $1", onabort)
}

// Onerror prop
func (*XMLHttpRequestEventTarget) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*XMLHttpRequestEventTarget) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onload prop
func (*XMLHttpRequestEventTarget) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// Onload prop
func (*XMLHttpRequestEventTarget) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.onload = $1", onload)
}

// Onloadend prop
func (*XMLHttpRequestEventTarget) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$<.onloadend")
	return onloadend
}

// Onloadend prop
func (*XMLHttpRequestEventTarget) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$<.onloadend = $1", onloadend)
}

// Onloadstart prop
func (*XMLHttpRequestEventTarget) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$<.onloadstart")
	return onloadstart
}

// Onloadstart prop
func (*XMLHttpRequestEventTarget) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$<.onloadstart = $1", onloadstart)
}

// Onprogress prop
func (*XMLHttpRequestEventTarget) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$<.onprogress")
	return onprogress
}

// Onprogress prop
func (*XMLHttpRequestEventTarget) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$<.onprogress = $1", onprogress)
}

// Ontimeout prop
func (*XMLHttpRequestEventTarget) Ontimeout() (ontimeout func(window.Event)) {
	js.Rewrite("$<.ontimeout")
	return ontimeout
}

// Ontimeout prop
func (*XMLHttpRequestEventTarget) SetOntimeout(ontimeout func(window.Event)) {
	js.Rewrite("$<.ontimeout = $1", ontimeout)
}
