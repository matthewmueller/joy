package window

import (
	"github.com/matthewmueller/joy/dom/notificationdirection"
	"github.com/matthewmueller/joy/dom/notificationoptions"
	"github.com/matthewmueller/joy/dom/notificationpermission"
	"github.com/matthewmueller/joy/macro"
)

var _ EventTarget = (*Notification)(nil)

// NewNotification fn
func NewNotification(title string, options *notificationoptions.NotificationOptions) *Notification {
	macro.Rewrite("Notification")
	return &Notification{}
}

// Notification struct
// js:"Notification,omit"
type Notification struct {
}

// Close fn
// js:"close"
func (*Notification) Close() {
	macro.Rewrite("$_.close()")
}

// RequestPermission fn
// js:"requestPermission"
func (*Notification) RequestPermission(callback *func(permission *notificationpermission.NotificationPermission)) (n *notificationpermission.NotificationPermission) {
	macro.Rewrite("await $_.requestPermission($1)", callback)
	return n
}

// AddEventListener fn
// js:"addEventListener"
func (*Notification) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*Notification) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*Notification) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Body prop
// js:"body"
func (*Notification) Body() (body string) {
	macro.Rewrite("$_.body")
	return body
}

// Dir prop
// js:"dir"
func (*Notification) Dir() (dir *notificationdirection.NotificationDirection) {
	macro.Rewrite("$_.dir")
	return dir
}

// Icon prop
// js:"icon"
func (*Notification) Icon() (icon string) {
	macro.Rewrite("$_.icon")
	return icon
}

// Lang prop
// js:"lang"
func (*Notification) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

// Onclick prop
// js:"onclick"
func (*Notification) Onclick() (onclick func(Event)) {
	macro.Rewrite("$_.onclick")
	return onclick
}

// SetOnclick prop
// js:"onclick"
func (*Notification) SetOnclick(onclick func(Event)) {
	macro.Rewrite("$_.onclick = $1", onclick)
}

// Onclose prop
// js:"onclose"
func (*Notification) Onclose() (onclose func(Event)) {
	macro.Rewrite("$_.onclose")
	return onclose
}

// SetOnclose prop
// js:"onclose"
func (*Notification) SetOnclose(onclose func(Event)) {
	macro.Rewrite("$_.onclose = $1", onclose)
}

// Onerror prop
// js:"onerror"
func (*Notification) Onerror() (onerror func(Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*Notification) SetOnerror(onerror func(Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onshow prop
// js:"onshow"
func (*Notification) Onshow() (onshow func(Event)) {
	macro.Rewrite("$_.onshow")
	return onshow
}

// SetOnshow prop
// js:"onshow"
func (*Notification) SetOnshow(onshow func(Event)) {
	macro.Rewrite("$_.onshow = $1", onshow)
}

// Permission prop
// js:"permission"
func (*Notification) Permission() (permission *notificationpermission.NotificationPermission) {
	macro.Rewrite("$_.permission")
	return permission
}

// Tag prop
// js:"tag"
func (*Notification) Tag() (tag string) {
	macro.Rewrite("$_.tag")
	return tag
}

// Title prop
// js:"title"
func (*Notification) Title() (title string) {
	macro.Rewrite("$_.title")
	return title
}
