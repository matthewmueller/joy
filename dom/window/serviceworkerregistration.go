package window

import (
	"github.com/matthewmueller/golly/dom/getnotificationoptions"
	"github.com/matthewmueller/golly/dom/notificationoptions"
	"github.com/matthewmueller/golly/dom/pushmanager"
	"github.com/matthewmueller/golly/dom/syncmanager"
	"github.com/matthewmueller/golly/js"
)

// ServiceWorkerRegistration struct
// js:"ServiceWorkerRegistration,omit"
type ServiceWorkerRegistration struct {
	EventTarget
}

// GetNotifications fn
func (*ServiceWorkerRegistration) GetNotifications(filter *getnotificationoptions.GetNotificationOptions) (n []*Notification) {
	js.Rewrite("await $<.getNotifications($1)", filter)
	return n
}

// ShowNotification fn
func (*ServiceWorkerRegistration) ShowNotification(title string, options *notificationoptions.NotificationOptions) {
	js.Rewrite("await $<.showNotification($1, $2)", title, options)
}

// Unregister fn
func (*ServiceWorkerRegistration) Unregister() (b bool) {
	js.Rewrite("await $<.unregister()")
	return b
}

// Update fn
func (*ServiceWorkerRegistration) Update() {
	js.Rewrite("await $<.update()")
}

// Active prop
func (*ServiceWorkerRegistration) Active() (active *ServiceWorker) {
	js.Rewrite("$<.active")
	return active
}

// Installing prop
func (*ServiceWorkerRegistration) Installing() (installing *ServiceWorker) {
	js.Rewrite("$<.installing")
	return installing
}

// Onupdatefound prop
func (*ServiceWorkerRegistration) Onupdatefound() (onupdatefound func(Event)) {
	js.Rewrite("$<.onupdatefound")
	return onupdatefound
}

// Onupdatefound prop
func (*ServiceWorkerRegistration) SetOnupdatefound(onupdatefound func(Event)) {
	js.Rewrite("$<.onupdatefound = $1", onupdatefound)
}

// PushManager prop
func (*ServiceWorkerRegistration) PushManager() (pushManager *pushmanager.PushManager) {
	js.Rewrite("$<.pushManager")
	return pushManager
}

// Scope prop
func (*ServiceWorkerRegistration) Scope() (scope string) {
	js.Rewrite("$<.scope")
	return scope
}

// Sync prop
func (*ServiceWorkerRegistration) Sync() (sync *syncmanager.SyncManager) {
	js.Rewrite("$<.sync")
	return sync
}

// Waiting prop
func (*ServiceWorkerRegistration) Waiting() (waiting *ServiceWorker) {
	js.Rewrite("$<.waiting")
	return waiting
}
