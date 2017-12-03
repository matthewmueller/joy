package pushsubscription

import (
	"github.com/matthewmueller/joy/dom/pushencryptionkeyname"
	"github.com/matthewmueller/joy/dom/pushsubscriptionoptions"
	"github.com/matthewmueller/joy/macro"
)

// PushSubscription struct
// js:"PushSubscription,omit"
type PushSubscription struct {
}

// GetKey fn
// js:"getKey"
func (*PushSubscription) GetKey(name *pushencryptionkeyname.PushEncryptionKeyName) (b []byte) {
	macro.Rewrite("$_.getKey($1)", name)
	return b
}

// ToJSON fn
// js:"toJSON"
func (*PushSubscription) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

// Unsubscribe fn
// js:"unsubscribe"
func (*PushSubscription) Unsubscribe() (b bool) {
	macro.Rewrite("await $_.unsubscribe()")
	return b
}

// Endpoint prop
// js:"endpoint"
func (*PushSubscription) Endpoint() (endpoint string) {
	macro.Rewrite("$_.endpoint")
	return endpoint
}

// Options prop
// js:"options"
func (*PushSubscription) Options() (options *pushsubscriptionoptions.PushSubscriptionOptions) {
	macro.Rewrite("$_.options")
	return options
}
