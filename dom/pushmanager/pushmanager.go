package pushmanager

import (
	"github.com/matthewmueller/golly/dom/pushpermissionstate"
	"github.com/matthewmueller/golly/dom/pushsubscription"
	"github.com/matthewmueller/golly/dom/pushsubscriptionoptionsinit"
	"github.com/matthewmueller/golly/js"
)

// PushManager struct
// js:"PushManager,omit"
type PushManager struct {
}

// GetSubscription fn
func (*PushManager) GetSubscription() (p *pushsubscription.PushSubscription) {
	js.Rewrite("await $<.getSubscription()")
	return p
}

// PermissionState fn
func (*PushManager) PermissionState(options *pushsubscriptionoptionsinit.PushSubscriptionOptionsInit) (p *pushpermissionstate.PushPermissionState) {
	js.Rewrite("await $<.permissionState($1)", options)
	return p
}

// Subscribe fn
func (*PushManager) Subscribe(options *pushsubscriptionoptionsinit.PushSubscriptionOptionsInit) (p *pushsubscription.PushSubscription) {
	js.Rewrite("await $<.subscribe($1)", options)
	return p
}
