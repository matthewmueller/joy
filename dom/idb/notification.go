package idb

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

// Close
func (*Notification) Close() {
	js.Rewrite("$<.Close()")
}

// RequestPermission
func (*Notification) RequestPermission(callback *func(permission *notificationpermission.NotificationPermission)) (n *notificationpermission.NotificationPermission) {
	js.Rewrite("await $<.RequestPermission($1)", callback)
	return n
}

// Body
func (*Notification) Body() (body string) {
	js.Rewrite("$<.Body")
	return body
}

// Dir
func (*Notification) Dir() (dir *notificationdirection.NotificationDirection) {
	js.Rewrite("$<.Dir")
	return dir
}

// Icon
func (*Notification) Icon() (icon string) {
	js.Rewrite("$<.Icon")
	return icon
}

// Lang
func (*Notification) Lang() (lang string) {
	js.Rewrite("$<.Lang")
	return lang
}

// Onclick
func (*Notification) Onclick() (onclick func(Event)) {
	js.Rewrite("$<.Onclick")
	return onclick
}

// Onclick
func (*Notification) SetOnclick(onclick func(Event)) {
	js.Rewrite("$<.Onclick = $1", onclick)
}

// Onclose
func (*Notification) Onclose() (onclose func(Event)) {
	js.Rewrite("$<.Onclose")
	return onclose
}

// Onclose
func (*Notification) SetOnclose(onclose func(Event)) {
	js.Rewrite("$<.Onclose = $1", onclose)
}

// Onerror
func (*Notification) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*Notification) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onshow
func (*Notification) Onshow() (onshow func(Event)) {
	js.Rewrite("$<.Onshow")
	return onshow
}

// Onshow
func (*Notification) SetOnshow(onshow func(Event)) {
	js.Rewrite("$<.Onshow = $1", onshow)
}

// Permission
func (*Notification) Permission() (permission *notificationpermission.NotificationPermission) {
	js.Rewrite("$<.Permission")
	return permission
}

// Tag
func (*Notification) Tag() (tag string) {
	js.Rewrite("$<.Tag")
	return tag
}

// Title
func (*Notification) Title() (title string) {
	js.Rewrite("$<.Title")
	return title
}
