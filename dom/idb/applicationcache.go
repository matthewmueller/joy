package idb

import "github.com/matthewmueller/golly/js"

// ApplicationCache struct
// js:"ApplicationCache,omit"
type ApplicationCache struct {
	EventTarget
}

// Abort
func (*ApplicationCache) Abort() {
	js.Rewrite("$<.Abort()")
}

// SwapCache
func (*ApplicationCache) SwapCache() {
	js.Rewrite("$<.SwapCache()")
}

// Update
func (*ApplicationCache) Update() {
	js.Rewrite("$<.Update()")
}

// Oncached
func (*ApplicationCache) Oncached() (oncached func(Event)) {
	js.Rewrite("$<.Oncached")
	return oncached
}

// Oncached
func (*ApplicationCache) SetOncached(oncached func(Event)) {
	js.Rewrite("$<.Oncached = $1", oncached)
}

// Onchecking
func (*ApplicationCache) Onchecking() (onchecking func(Event)) {
	js.Rewrite("$<.Onchecking")
	return onchecking
}

// Onchecking
func (*ApplicationCache) SetOnchecking(onchecking func(Event)) {
	js.Rewrite("$<.Onchecking = $1", onchecking)
}

// Ondownloading
func (*ApplicationCache) Ondownloading() (ondownloading func(Event)) {
	js.Rewrite("$<.Ondownloading")
	return ondownloading
}

// Ondownloading
func (*ApplicationCache) SetOndownloading(ondownloading func(Event)) {
	js.Rewrite("$<.Ondownloading = $1", ondownloading)
}

// Onerror
func (*ApplicationCache) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*ApplicationCache) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onnoupdate
func (*ApplicationCache) Onnoupdate() (onnoupdate func(Event)) {
	js.Rewrite("$<.Onnoupdate")
	return onnoupdate
}

// Onnoupdate
func (*ApplicationCache) SetOnnoupdate(onnoupdate func(Event)) {
	js.Rewrite("$<.Onnoupdate = $1", onnoupdate)
}

// Onobsolete
func (*ApplicationCache) Onobsolete() (onobsolete func(Event)) {
	js.Rewrite("$<.Onobsolete")
	return onobsolete
}

// Onobsolete
func (*ApplicationCache) SetOnobsolete(onobsolete func(Event)) {
	js.Rewrite("$<.Onobsolete = $1", onobsolete)
}

// Onprogress
func (*ApplicationCache) Onprogress() (onprogress func(*ProgressEvent)) {
	js.Rewrite("$<.Onprogress")
	return onprogress
}

// Onprogress
func (*ApplicationCache) SetOnprogress(onprogress func(*ProgressEvent)) {
	js.Rewrite("$<.Onprogress = $1", onprogress)
}

// Onupdateready
func (*ApplicationCache) Onupdateready() (onupdateready func(Event)) {
	js.Rewrite("$<.Onupdateready")
	return onupdateready
}

// Onupdateready
func (*ApplicationCache) SetOnupdateready(onupdateready func(Event)) {
	js.Rewrite("$<.Onupdateready = $1", onupdateready)
}

// Status
func (*ApplicationCache) Status() (status uint8) {
	js.Rewrite("$<.Status")
	return status
}
