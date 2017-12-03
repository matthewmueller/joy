package window

import (
	"github.com/matthewmueller/joy/dom/getnotificationoptions"
	"github.com/matthewmueller/joy/dom/notificationoptions"
	"github.com/matthewmueller/joy/dom/pushmanager"
	"github.com/matthewmueller/joy/dom/syncmanager"
	"github.com/matthewmueller/joy/js"
)

var _ EventTarget = (*ServiceWorkerRegistration)(nil)

// ServiceWorkerRegistration struct
// js:"ServiceWorkerRegistration,omit"
type ServiceWorkerRegistration struct {
}

// GetNotifications fn
// js:"getNotifications"
func (*ServiceWorkerRegistration) GetNotifications(filter *getnotificationoptions.GetNotificationOptions) (n []*Notification) {
	js.Rewrite("await $_.getNotifications($1)", filter)
	return n
}

// ShowNotification fn
// js:"showNotification"
func (*ServiceWorkerRegistration) ShowNotification(title string, options *notificationoptions.NotificationOptions) {
	js.Rewrite("await $_.showNotification($1, $2)", title, options)
}

// Unregister fn
// js:"unregister"
func (*ServiceWorkerRegistration) Unregister() (b bool) {
	js.Rewrite("await $_.unregister()")
	return b
}

// Update fn
// js:"update"
func (*ServiceWorkerRegistration) Update() {
	js.Rewrite("await $_.update()")
}

// AddEventListener fn
// js:"addEventListener"
func (*ServiceWorkerRegistration) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*ServiceWorkerRegistration) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*ServiceWorkerRegistration) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Active prop
// js:"active"
func (*ServiceWorkerRegistration) Active() (active *ServiceWorker) {
	js.Rewrite("$_.active")
	return active
}

// Installing prop
// js:"installing"
func (*ServiceWorkerRegistration) Installing() (installing *ServiceWorker) {
	js.Rewrite("$_.installing")
	return installing
}

// Onupdatefound prop
// js:"onupdatefound"
func (*ServiceWorkerRegistration) Onupdatefound() (onupdatefound func(Event)) {
	js.Rewrite("$_.onupdatefound")
	return onupdatefound
}

// SetOnupdatefound prop
// js:"onupdatefound"
func (*ServiceWorkerRegistration) SetOnupdatefound(onupdatefound func(Event)) {
	js.Rewrite("$_.onupdatefound = $1", onupdatefound)
}

// PushManager prop
// js:"pushManager"
func (*ServiceWorkerRegistration) PushManager() (pushManager *pushmanager.PushManager) {
	js.Rewrite("$_.pushManager")
	return pushManager
}

// Scope prop
// js:"scope"
func (*ServiceWorkerRegistration) Scope() (scope string) {
	js.Rewrite("$_.scope")
	return scope
}

// Sync prop
// js:"sync"
func (*ServiceWorkerRegistration) Sync() (sync *syncmanager.SyncManager) {
	js.Rewrite("$_.sync")
	return sync
}

// Waiting prop
// js:"waiting"
func (*ServiceWorkerRegistration) Waiting() (waiting *ServiceWorker) {
	js.Rewrite("$_.waiting")
	return waiting
}
