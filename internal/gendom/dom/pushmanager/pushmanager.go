package pushmanager

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/pushpermissionstate"
	"github.com/matthewmueller/golly/internal/gendom/dom/pushsubscription"
	"github.com/matthewmueller/golly/internal/gendom/dom/pushsubscriptionoptionsinit"
	"github.com/matthewmueller/golly/js"
)

type PushManager struct {
}

func (*PushManager) GetSubscription() (p *pushsubscription.PushSubscription) {
	js.Rewrite("await $<.getSubscription()")
	return p
}

func (*PushManager) PermissionState(options *pushsubscriptionoptionsinit.PushSubscriptionOptionsInit) (p *pushpermissionstate.PushPermissionState) {
	js.Rewrite("await $<.permissionState($1)", options)
	return p
}

func (*PushManager) Subscribe(options *pushsubscriptionoptionsinit.PushSubscriptionOptionsInit) (p *pushsubscription.PushSubscription) {
	js.Rewrite("await $<.subscribe($1)", options)
	return p
}
