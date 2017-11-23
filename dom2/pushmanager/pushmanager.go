package pushmanager

import (
	"github.com/matthewmueller/golly/dom2/pushpermissionstate"
	"github.com/matthewmueller/golly/dom2/pushsubscription"
	"github.com/matthewmueller/golly/dom2/pushsubscriptionoptionsinit"
	"github.com/matthewmueller/golly/js"
)

// js:"PushManager,omit"
type PushManager struct {
}

// GetSubscription
func (*PushManager) GetSubscription() (p *pushsubscription.PushSubscription) {
	js.Rewrite("await $<.GetSubscription()")
	return p
}

// PermissionState
func (*PushManager) PermissionState(options *pushsubscriptionoptionsinit.PushSubscriptionOptionsInit) (p *pushpermissionstate.PushPermissionState) {
	js.Rewrite("await $<.PermissionState($1)", options)
	return p
}

// Subscribe
func (*PushManager) Subscribe(options *pushsubscriptionoptionsinit.PushSubscriptionOptionsInit) (p *pushsubscription.PushSubscription) {
	js.Rewrite("await $<.Subscribe($1)", options)
	return p
}
