package window

import (
	"github.com/matthewmueller/golly/dom2/getnotificationoptions"
	"github.com/matthewmueller/golly/dom2/notificationoptions"
	"github.com/matthewmueller/golly/dom2/pushmanager"
	"github.com/matthewmueller/golly/dom2/syncmanager"
	"github.com/matthewmueller/golly/js"
)

// ServiceWorkerRegistration struct
// js:"ServiceWorkerRegistration,omit"
type ServiceWorkerRegistration struct {
	EventTarget
}

// GetNotifications
func (*ServiceWorkerRegistration) GetNotifications(filter *getnotificationoptions.GetNotificationOptions) (n []*Notification) {
	js.Rewrite("await $<.GetNotifications($1)", filter)
	return n
}

// ShowNotification
func (*ServiceWorkerRegistration) ShowNotification(title string, options *notificationoptions.NotificationOptions) {
	js.Rewrite("await $<.ShowNotification($1, $2)", title, options)
}

// Unregister
func (*ServiceWorkerRegistration) Unregister() (b bool) {
	js.Rewrite("await $<.Unregister()")
	return b
}

// Update
func (*ServiceWorkerRegistration) Update() {
	js.Rewrite("await $<.Update()")
}

// Active
func (*ServiceWorkerRegistration) Active() (active *ServiceWorker) {
	js.Rewrite("$<.Active")
	return active
}

// Installing
func (*ServiceWorkerRegistration) Installing() (installing *ServiceWorker) {
	js.Rewrite("$<.Installing")
	return installing
}

// Onupdatefound
func (*ServiceWorkerRegistration) Onupdatefound() (onupdatefound func(Event)) {
	js.Rewrite("$<.Onupdatefound")
	return onupdatefound
}

// Onupdatefound
func (*ServiceWorkerRegistration) SetOnupdatefound(onupdatefound func(Event)) {
	js.Rewrite("$<.Onupdatefound = $1", onupdatefound)
}

// PushManager
func (*ServiceWorkerRegistration) PushManager() (pushManager *pushmanager.PushManager) {
	js.Rewrite("$<.PushManager")
	return pushManager
}

// Scope
func (*ServiceWorkerRegistration) Scope() (scope string) {
	js.Rewrite("$<.Scope")
	return scope
}

// Sync
func (*ServiceWorkerRegistration) Sync() (sync *syncmanager.SyncManager) {
	js.Rewrite("$<.Sync")
	return sync
}

// Waiting
func (*ServiceWorkerRegistration) Waiting() (waiting *ServiceWorker) {
	js.Rewrite("$<.Waiting")
	return waiting
}
