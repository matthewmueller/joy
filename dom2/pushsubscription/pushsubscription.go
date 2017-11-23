package pushsubscription

import (
	"github.com/matthewmueller/golly/dom2/pushencryptionkeyname"
	"github.com/matthewmueller/golly/dom2/pushsubscriptionoptions"
	"github.com/matthewmueller/golly/js"
)

// js:"PushSubscription,omit"
type PushSubscription struct {
}

// GetKey
func (*PushSubscription) GetKey(name *pushencryptionkeyname.PushEncryptionKeyName) (b []byte) {
	js.Rewrite("$<.GetKey($1)", name)
	return b
}

// ToJSON
func (*PushSubscription) ToJSON() (i interface{}) {
	js.Rewrite("$<.ToJSON()")
	return i
}

// Unsubscribe
func (*PushSubscription) Unsubscribe() (b bool) {
	js.Rewrite("await $<.Unsubscribe()")
	return b
}

// Endpoint
func (*PushSubscription) Endpoint() (endpoint string) {
	js.Rewrite("$<.Endpoint")
	return endpoint
}

// Options
func (*PushSubscription) Options() (options *pushsubscriptionoptions.PushSubscriptionOptions) {
	js.Rewrite("$<.Options")
	return options
}
