package notification

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/notificationdirection"
	"github.com/matthewmueller/golly/internal/gendom/dom/notificationpermission"
	"github.com/matthewmueller/golly/js"
)

type Notification struct {
	*eventtarget.EventTarget
}

func (*Notification) Close() {
	js.Rewrite("$<.close()")
}

func (*Notification) RequestPermission(callback func(permission *notificationpermission.NotificationPermission)) (n *notificationpermission.NotificationPermission) {
	js.Rewrite("await $<.requestPermission($1)", callback)
	return n
}

func (*Notification) GetBody() (body string) {
	js.Rewrite("$<.body")
	return body
}

func (*Notification) GetDir() (dir *notificationdirection.NotificationDirection) {
	js.Rewrite("$<.dir")
	return dir
}

func (*Notification) GetIcon() (icon string) {
	js.Rewrite("$<.icon")
	return icon
}

func (*Notification) GetLang() (lang string) {
	js.Rewrite("$<.lang")
	return lang
}

func (*Notification) GetOnclick() (click *event.Event) {
	js.Rewrite("$<.onclick")
	return click
}

func (*Notification) SetOnclick(click *event.Event) {
	js.Rewrite("$<.onclick = $1", click)
}

func (*Notification) GetOnclose() (cls *event.Event) {
	js.Rewrite("$<.onclose")
	return cls
}

func (*Notification) SetOnclose(cls *event.Event) {
	js.Rewrite("$<.onclose = $1", cls)
}

func (*Notification) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*Notification) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*Notification) GetOnshow() (show *event.Event) {
	js.Rewrite("$<.onshow")
	return show
}

func (*Notification) SetOnshow(show *event.Event) {
	js.Rewrite("$<.onshow = $1", show)
}

func (*Notification) GetPermission() (permission *notificationpermission.NotificationPermission) {
	js.Rewrite("$<.permission")
	return permission
}

func (*Notification) GetTag() (tag string) {
	js.Rewrite("$<.tag")
	return tag
}

func (*Notification) GetTitle() (title string) {
	js.Rewrite("$<.title")
	return title
}
