package window

import (
	"github.com/matthewmueller/golly/dom2/notificationdirection"
	"github.com/matthewmueller/golly/dom2/notificationoptions"
	"github.com/matthewmueller/golly/dom2/notificationpermission"
	"github.com/matthewmueller/golly/js"
)

// NewNotification fn
func NewNotification(title string, options *notificationoptions.NotificationOptions) *Notification {
	js.Rewrite("Notification")
	return &Notification{}
}

// Notification struct
// js:"Notification,omit"
type Notification struct {
	EventTarget
}

// Close fn
func (*Notification) Close() {
	js.Rewrite("$<.close()")
}

// RequestPermission fn
func (*Notification) RequestPermission(callback *func(permission *notificationpermission.NotificationPermission)) (n *notificationpermission.NotificationPermission) {
	js.Rewrite("await $<.requestPermission($1)", callback)
	return n
}

// Body prop
func (*Notification) Body() (body string) {
	js.Rewrite("$<.body")
	return body
}

// Dir prop
func (*Notification) Dir() (dir *notificationdirection.NotificationDirection) {
	js.Rewrite("$<.dir")
	return dir
}

// Icon prop
func (*Notification) Icon() (icon string) {
	js.Rewrite("$<.icon")
	return icon
}

// Lang prop
func (*Notification) Lang() (lang string) {
	js.Rewrite("$<.lang")
	return lang
}

// Onclick prop
func (*Notification) Onclick() (onclick func(Event)) {
	js.Rewrite("$<.onclick")
	return onclick
}

// Onclick prop
func (*Notification) SetOnclick(onclick func(Event)) {
	js.Rewrite("$<.onclick = $1", onclick)
}

// Onclose prop
func (*Notification) Onclose() (onclose func(Event)) {
	js.Rewrite("$<.onclose")
	return onclose
}

// Onclose prop
func (*Notification) SetOnclose(onclose func(Event)) {
	js.Rewrite("$<.onclose = $1", onclose)
}

// Onerror prop
func (*Notification) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*Notification) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onshow prop
func (*Notification) Onshow() (onshow func(Event)) {
	js.Rewrite("$<.onshow")
	return onshow
}

// Onshow prop
func (*Notification) SetOnshow(onshow func(Event)) {
	js.Rewrite("$<.onshow = $1", onshow)
}

// Permission prop
func (*Notification) Permission() (permission *notificationpermission.NotificationPermission) {
	js.Rewrite("$<.permission")
	return permission
}

// Tag prop
func (*Notification) Tag() (tag string) {
	js.Rewrite("$<.tag")
	return tag
}

// Title prop
func (*Notification) Title() (title string) {
	js.Rewrite("$<.title")
	return title
}
