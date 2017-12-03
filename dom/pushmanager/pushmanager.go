package pushmanager

import (
	"github.com/matthewmueller/joy/dom/pushpermissionstate"
	"github.com/matthewmueller/joy/dom/pushsubscription"
	"github.com/matthewmueller/joy/dom/pushsubscriptionoptionsinit"
	"github.com/matthewmueller/joy/js"
)

// PushManager struct
// js:"PushManager,omit"
type PushManager struct {
}

// GetSubscription fn
// js:"getSubscription"
func (*PushManager) GetSubscription() (p *pushsubscription.PushSubscription) {
	js.Rewrite("await $_.getSubscription()")
	return p
}

// PermissionState fn
// js:"permissionState"
func (*PushManager) PermissionState(options *pushsubscriptionoptionsinit.PushSubscriptionOptionsInit) (p *pushpermissionstate.PushPermissionState) {
	js.Rewrite("await $_.permissionState($1)", options)
	return p
}

// Subscribe fn
// js:"subscribe"
func (*PushManager) Subscribe(options *pushsubscriptionoptionsinit.PushSubscriptionOptionsInit) (p *pushsubscription.PushSubscription) {
	js.Rewrite("await $_.subscribe($1)", options)
	return p
}
