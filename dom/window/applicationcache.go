package window

import "github.com/matthewmueller/golly/js"

// ApplicationCache struct
// js:"ApplicationCache,omit"
type ApplicationCache struct {
	EventTarget
}

// Abort fn
func (*ApplicationCache) Abort() {
	js.Rewrite("$<.abort()")
}

// SwapCache fn
func (*ApplicationCache) SwapCache() {
	js.Rewrite("$<.swapCache()")
}

// Update fn
func (*ApplicationCache) Update() {
	js.Rewrite("$<.update()")
}

// Oncached prop
func (*ApplicationCache) Oncached() (oncached func(Event)) {
	js.Rewrite("$<.oncached")
	return oncached
}

// Oncached prop
func (*ApplicationCache) SetOncached(oncached func(Event)) {
	js.Rewrite("$<.oncached = $1", oncached)
}

// Onchecking prop
func (*ApplicationCache) Onchecking() (onchecking func(Event)) {
	js.Rewrite("$<.onchecking")
	return onchecking
}

// Onchecking prop
func (*ApplicationCache) SetOnchecking(onchecking func(Event)) {
	js.Rewrite("$<.onchecking = $1", onchecking)
}

// Ondownloading prop
func (*ApplicationCache) Ondownloading() (ondownloading func(Event)) {
	js.Rewrite("$<.ondownloading")
	return ondownloading
}

// Ondownloading prop
func (*ApplicationCache) SetOndownloading(ondownloading func(Event)) {
	js.Rewrite("$<.ondownloading = $1", ondownloading)
}

// Onerror prop
func (*ApplicationCache) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*ApplicationCache) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onnoupdate prop
func (*ApplicationCache) Onnoupdate() (onnoupdate func(Event)) {
	js.Rewrite("$<.onnoupdate")
	return onnoupdate
}

// Onnoupdate prop
func (*ApplicationCache) SetOnnoupdate(onnoupdate func(Event)) {
	js.Rewrite("$<.onnoupdate = $1", onnoupdate)
}

// Onobsolete prop
func (*ApplicationCache) Onobsolete() (onobsolete func(Event)) {
	js.Rewrite("$<.onobsolete")
	return onobsolete
}

// Onobsolete prop
func (*ApplicationCache) SetOnobsolete(onobsolete func(Event)) {
	js.Rewrite("$<.onobsolete = $1", onobsolete)
}

// Onprogress prop
func (*ApplicationCache) Onprogress() (onprogress func(*ProgressEvent)) {
	js.Rewrite("$<.onprogress")
	return onprogress
}

// Onprogress prop
func (*ApplicationCache) SetOnprogress(onprogress func(*ProgressEvent)) {
	js.Rewrite("$<.onprogress = $1", onprogress)
}

// Onupdateready prop
func (*ApplicationCache) Onupdateready() (onupdateready func(Event)) {
	js.Rewrite("$<.onupdateready")
	return onupdateready
}

// Onupdateready prop
func (*ApplicationCache) SetOnupdateready(onupdateready func(Event)) {
	js.Rewrite("$<.onupdateready = $1", onupdateready)
}

// Status prop
func (*ApplicationCache) Status() (status uint8) {
	js.Rewrite("$<.status")
	return status
}
