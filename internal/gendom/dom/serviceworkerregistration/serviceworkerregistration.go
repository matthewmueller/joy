package serviceworkerregistration

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/getnotificationoptions"
	"github.com/matthewmueller/golly/internal/gendom/dom/notification"
	"github.com/matthewmueller/golly/internal/gendom/dom/notificationoptions"
	"github.com/matthewmueller/golly/internal/gendom/dom/pushmanager"
	"github.com/matthewmueller/golly/internal/gendom/dom/serviceworker"
	"github.com/matthewmueller/golly/internal/gendom/dom/syncmanager"
	"github.com/matthewmueller/golly/js"
)

type ServiceWorkerRegistration struct {
	*eventtarget.EventTarget
}

func (*ServiceWorkerRegistration) GetNotifications(filter *getnotificationoptions.GetNotificationOptions) (n []*notification.Notification) {
	js.Rewrite("await $<.getNotifications($1)", filter)
	return n
}

func (*ServiceWorkerRegistration) ShowNotification(title string, options *notificationoptions.NotificationOptions) {
	js.Rewrite("await $<.showNotification($1, $2)", title, options)
}

func (*ServiceWorkerRegistration) Unregister() (b bool) {
	js.Rewrite("await $<.unregister()")
	return b
}

func (*ServiceWorkerRegistration) Update() {
	js.Rewrite("await $<.update()")
}

func (*ServiceWorkerRegistration) GetActive() (active *serviceworker.ServiceWorker) {
	js.Rewrite("$<.active")
	return active
}

func (*ServiceWorkerRegistration) GetInstalling() (installing *serviceworker.ServiceWorker) {
	js.Rewrite("$<.installing")
	return installing
}

func (*ServiceWorkerRegistration) GetOnupdatefound() (updatefound *event.Event) {
	js.Rewrite("$<.onupdatefound")
	return updatefound
}

func (*ServiceWorkerRegistration) SetOnupdatefound(updatefound *event.Event) {
	js.Rewrite("$<.onupdatefound = $1", updatefound)
}

func (*ServiceWorkerRegistration) GetPushManager() (pushManager *pushmanager.PushManager) {
	js.Rewrite("$<.pushManager")
	return pushManager
}

func (*ServiceWorkerRegistration) GetScope() (scope string) {
	js.Rewrite("$<.scope")
	return scope
}

func (*ServiceWorkerRegistration) GetSync() (sync *syncmanager.SyncManager) {
	js.Rewrite("$<.sync")
	return sync
}

func (*ServiceWorkerRegistration) GetWaiting() (waiting *serviceworker.ServiceWorker) {
	js.Rewrite("$<.waiting")
	return waiting
}
