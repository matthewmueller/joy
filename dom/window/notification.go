package window

import (
	"github.com/matthewmueller/golly/dom/notificationdirection"
	"github.com/matthewmueller/golly/dom/notificationoptions"
	"github.com/matthewmueller/golly/dom/notificationpermission"
	"github.com/matthewmueller/golly/js"
)

var _ EventTarget = (*Notification)(nil)

// NewNotification fn
func NewNotification(title string, options *notificationoptions.NotificationOptions) *Notification {
	js.Rewrite("Notification")
	return &Notification{}
}

// Notification struct
// js:"Notification,omit"
type Notification struct {
}

// Close fn
// js:"close"
func (*Notification) Close() {
	js.Rewrite("$_.close()")
}

// RequestPermission fn
// js:"requestPermission"
func (*Notification) RequestPermission(callback *func(permission *notificationpermission.NotificationPermission)) (n *notificationpermission.NotificationPermission) {
	js.Rewrite("await $_.requestPermission($1)", callback)
	return n
}

// AddEventListener fn
// js:"addEventListener"
func (*Notification) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*Notification) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*Notification) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Body prop
// js:"body"
func (*Notification) Body() (body string) {
	js.Rewrite("$_.body")
	return body
}

// Dir prop
// js:"dir"
func (*Notification) Dir() (dir *notificationdirection.NotificationDirection) {
	js.Rewrite("$_.dir")
	return dir
}

// Icon prop
// js:"icon"
func (*Notification) Icon() (icon string) {
	js.Rewrite("$_.icon")
	return icon
}

// Lang prop
// js:"lang"
func (*Notification) Lang() (lang string) {
	js.Rewrite("$_.lang")
	return lang
}

// Onclick prop
// js:"onclick"
func (*Notification) Onclick() (onclick func(Event)) {
	js.Rewrite("$_.onclick")
	return onclick
}

// SetOnclick prop
// js:"onclick"
func (*Notification) SetOnclick(onclick func(Event)) {
	js.Rewrite("$_.onclick = $1", onclick)
}

// Onclose prop
// js:"onclose"
func (*Notification) Onclose() (onclose func(Event)) {
	js.Rewrite("$_.onclose")
	return onclose
}

// SetOnclose prop
// js:"onclose"
func (*Notification) SetOnclose(onclose func(Event)) {
	js.Rewrite("$_.onclose = $1", onclose)
}

// Onerror prop
// js:"onerror"
func (*Notification) Onerror() (onerror func(Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*Notification) SetOnerror(onerror func(Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onshow prop
// js:"onshow"
func (*Notification) Onshow() (onshow func(Event)) {
	js.Rewrite("$_.onshow")
	return onshow
}

// SetOnshow prop
// js:"onshow"
func (*Notification) SetOnshow(onshow func(Event)) {
	js.Rewrite("$_.onshow = $1", onshow)
}

// Permission prop
// js:"permission"
func (*Notification) Permission() (permission *notificationpermission.NotificationPermission) {
	js.Rewrite("$_.permission")
	return permission
}

// Tag prop
// js:"tag"
func (*Notification) Tag() (tag string) {
	js.Rewrite("$_.tag")
	return tag
}

// Title prop
// js:"title"
func (*Notification) Title() (title string) {
	js.Rewrite("$_.title")
	return title
}
