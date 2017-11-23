package xmlhttprequesteventtarget

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// XMLHTTPRequestEventTarget struct
// js:"XMLHTTPRequestEventTarget,omit"
type XMLHTTPRequestEventTarget struct {
}

// Onabort
func (*XMLHttpRequestEventTarget) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$<.Onabort")
	return onabort
}

// Onabort
func (*XMLHttpRequestEventTarget) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$<.Onabort = $1", onabort)
}

// Onerror
func (*XMLHttpRequestEventTarget) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*XMLHttpRequestEventTarget) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onload
func (*XMLHttpRequestEventTarget) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.Onload")
	return onload
}

// Onload
func (*XMLHttpRequestEventTarget) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.Onload = $1", onload)
}

// Onloadend
func (*XMLHttpRequestEventTarget) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$<.Onloadend")
	return onloadend
}

// Onloadend
func (*XMLHttpRequestEventTarget) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$<.Onloadend = $1", onloadend)
}

// Onloadstart
func (*XMLHttpRequestEventTarget) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$<.Onloadstart")
	return onloadstart
}

// Onloadstart
func (*XMLHttpRequestEventTarget) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$<.Onloadstart = $1", onloadstart)
}

// Onprogress
func (*XMLHttpRequestEventTarget) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$<.Onprogress")
	return onprogress
}

// Onprogress
func (*XMLHttpRequestEventTarget) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$<.Onprogress = $1", onprogress)
}

// Ontimeout
func (*XMLHttpRequestEventTarget) Ontimeout() (ontimeout func(window.Event)) {
	js.Rewrite("$<.Ontimeout")
	return ontimeout
}

// Ontimeout
func (*XMLHttpRequestEventTarget) SetOntimeout(ontimeout func(window.Event)) {
	js.Rewrite("$<.Ontimeout = $1", ontimeout)
}
