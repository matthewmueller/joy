package window

import "github.com/matthewmueller/golly/js"

var _ EventTarget = (*ApplicationCache)(nil)

// ApplicationCache struct
// js:"ApplicationCache,omit"
type ApplicationCache struct {
}

// Abort fn
// js:"abort"
func (*ApplicationCache) Abort() {
	js.Rewrite("$_.abort()")
}

// SwapCache fn
// js:"swapCache"
func (*ApplicationCache) SwapCache() {
	js.Rewrite("$_.swapCache()")
}

// Update fn
// js:"update"
func (*ApplicationCache) Update() {
	js.Rewrite("$_.update()")
}

// AddEventListener fn
// js:"addEventListener"
func (*ApplicationCache) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*ApplicationCache) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*ApplicationCache) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Oncached prop
// js:"oncached"
func (*ApplicationCache) Oncached() (oncached func(Event)) {
	js.Rewrite("$_.oncached")
	return oncached
}

// SetOncached prop
// js:"oncached"
func (*ApplicationCache) SetOncached(oncached func(Event)) {
	js.Rewrite("$_.oncached = $1", oncached)
}

// Onchecking prop
// js:"onchecking"
func (*ApplicationCache) Onchecking() (onchecking func(Event)) {
	js.Rewrite("$_.onchecking")
	return onchecking
}

// SetOnchecking prop
// js:"onchecking"
func (*ApplicationCache) SetOnchecking(onchecking func(Event)) {
	js.Rewrite("$_.onchecking = $1", onchecking)
}

// Ondownloading prop
// js:"ondownloading"
func (*ApplicationCache) Ondownloading() (ondownloading func(Event)) {
	js.Rewrite("$_.ondownloading")
	return ondownloading
}

// SetOndownloading prop
// js:"ondownloading"
func (*ApplicationCache) SetOndownloading(ondownloading func(Event)) {
	js.Rewrite("$_.ondownloading = $1", ondownloading)
}

// Onerror prop
// js:"onerror"
func (*ApplicationCache) Onerror() (onerror func(Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*ApplicationCache) SetOnerror(onerror func(Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onnoupdate prop
// js:"onnoupdate"
func (*ApplicationCache) Onnoupdate() (onnoupdate func(Event)) {
	js.Rewrite("$_.onnoupdate")
	return onnoupdate
}

// SetOnnoupdate prop
// js:"onnoupdate"
func (*ApplicationCache) SetOnnoupdate(onnoupdate func(Event)) {
	js.Rewrite("$_.onnoupdate = $1", onnoupdate)
}

// Onobsolete prop
// js:"onobsolete"
func (*ApplicationCache) Onobsolete() (onobsolete func(Event)) {
	js.Rewrite("$_.onobsolete")
	return onobsolete
}

// SetOnobsolete prop
// js:"onobsolete"
func (*ApplicationCache) SetOnobsolete(onobsolete func(Event)) {
	js.Rewrite("$_.onobsolete = $1", onobsolete)
}

// Onprogress prop
// js:"onprogress"
func (*ApplicationCache) Onprogress() (onprogress func(*ProgressEvent)) {
	js.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*ApplicationCache) SetOnprogress(onprogress func(*ProgressEvent)) {
	js.Rewrite("$_.onprogress = $1", onprogress)
}

// Onupdateready prop
// js:"onupdateready"
func (*ApplicationCache) Onupdateready() (onupdateready func(Event)) {
	js.Rewrite("$_.onupdateready")
	return onupdateready
}

// SetOnupdateready prop
// js:"onupdateready"
func (*ApplicationCache) SetOnupdateready(onupdateready func(Event)) {
	js.Rewrite("$_.onupdateready = $1", onupdateready)
}

// Status prop
// js:"status"
func (*ApplicationCache) Status() (status uint8) {
	js.Rewrite("$_.status")
	return status
}
