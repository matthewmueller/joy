package window

import "github.com/matthewmueller/joy/macro"

var _ EventTarget = (*ApplicationCache)(nil)

// ApplicationCache struct
// js:"ApplicationCache,omit"
type ApplicationCache struct {
}

// Abort fn
// js:"abort"
func (*ApplicationCache) Abort() {
	macro.Rewrite("$_.abort()")
}

// SwapCache fn
// js:"swapCache"
func (*ApplicationCache) SwapCache() {
	macro.Rewrite("$_.swapCache()")
}

// Update fn
// js:"update"
func (*ApplicationCache) Update() {
	macro.Rewrite("$_.update()")
}

// AddEventListener fn
// js:"addEventListener"
func (*ApplicationCache) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*ApplicationCache) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*ApplicationCache) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Oncached prop
// js:"oncached"
func (*ApplicationCache) Oncached() (oncached func(Event)) {
	macro.Rewrite("$_.oncached")
	return oncached
}

// SetOncached prop
// js:"oncached"
func (*ApplicationCache) SetOncached(oncached func(Event)) {
	macro.Rewrite("$_.oncached = $1", oncached)
}

// Onchecking prop
// js:"onchecking"
func (*ApplicationCache) Onchecking() (onchecking func(Event)) {
	macro.Rewrite("$_.onchecking")
	return onchecking
}

// SetOnchecking prop
// js:"onchecking"
func (*ApplicationCache) SetOnchecking(onchecking func(Event)) {
	macro.Rewrite("$_.onchecking = $1", onchecking)
}

// Ondownloading prop
// js:"ondownloading"
func (*ApplicationCache) Ondownloading() (ondownloading func(Event)) {
	macro.Rewrite("$_.ondownloading")
	return ondownloading
}

// SetOndownloading prop
// js:"ondownloading"
func (*ApplicationCache) SetOndownloading(ondownloading func(Event)) {
	macro.Rewrite("$_.ondownloading = $1", ondownloading)
}

// Onerror prop
// js:"onerror"
func (*ApplicationCache) Onerror() (onerror func(Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*ApplicationCache) SetOnerror(onerror func(Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onnoupdate prop
// js:"onnoupdate"
func (*ApplicationCache) Onnoupdate() (onnoupdate func(Event)) {
	macro.Rewrite("$_.onnoupdate")
	return onnoupdate
}

// SetOnnoupdate prop
// js:"onnoupdate"
func (*ApplicationCache) SetOnnoupdate(onnoupdate func(Event)) {
	macro.Rewrite("$_.onnoupdate = $1", onnoupdate)
}

// Onobsolete prop
// js:"onobsolete"
func (*ApplicationCache) Onobsolete() (onobsolete func(Event)) {
	macro.Rewrite("$_.onobsolete")
	return onobsolete
}

// SetOnobsolete prop
// js:"onobsolete"
func (*ApplicationCache) SetOnobsolete(onobsolete func(Event)) {
	macro.Rewrite("$_.onobsolete = $1", onobsolete)
}

// Onprogress prop
// js:"onprogress"
func (*ApplicationCache) Onprogress() (onprogress func(*ProgressEvent)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*ApplicationCache) SetOnprogress(onprogress func(*ProgressEvent)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

// Onupdateready prop
// js:"onupdateready"
func (*ApplicationCache) Onupdateready() (onupdateready func(Event)) {
	macro.Rewrite("$_.onupdateready")
	return onupdateready
}

// SetOnupdateready prop
// js:"onupdateready"
func (*ApplicationCache) SetOnupdateready(onupdateready func(Event)) {
	macro.Rewrite("$_.onupdateready = $1", onupdateready)
}

// Status prop
// js:"status"
func (*ApplicationCache) Status() (status uint8) {
	macro.Rewrite("$_.status")
	return status
}
