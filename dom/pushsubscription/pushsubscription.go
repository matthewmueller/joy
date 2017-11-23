package pushsubscription

import (
	"github.com/matthewmueller/golly/dom/pushencryptionkeyname"
	"github.com/matthewmueller/golly/dom/pushsubscriptionoptions"
	"github.com/matthewmueller/golly/js"
)

// PushSubscription struct
// js:"PushSubscription,omit"
type PushSubscription struct {
}

// GetKey fn
func (*PushSubscription) GetKey(name *pushencryptionkeyname.PushEncryptionKeyName) (b []byte) {
	js.Rewrite("$<.getKey($1)", name)
	return b
}

// ToJSON fn
func (*PushSubscription) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

// Unsubscribe fn
func (*PushSubscription) Unsubscribe() (b bool) {
	js.Rewrite("await $<.unsubscribe()")
	return b
}

// Endpoint prop
func (*PushSubscription) Endpoint() (endpoint string) {
	js.Rewrite("$<.endpoint")
	return endpoint
}

// Options prop
func (*PushSubscription) Options() (options *pushsubscriptionoptions.PushSubscriptionOptions) {
	js.Rewrite("$<.options")
	return options
}
