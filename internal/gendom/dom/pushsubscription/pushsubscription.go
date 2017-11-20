package pushsubscription

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/pushencryptionkeyname"
	"github.com/matthewmueller/golly/internal/gendom/dom/pushsubscriptionoptions"
	"github.com/matthewmueller/golly/js"
)

type PushSubscription struct {
}

func (*PushSubscription) GetKey(name *pushencryptionkeyname.PushEncryptionKeyName) (b []byte) {
	js.Rewrite("$<.getKey($1)", name)
	return b
}

func (*PushSubscription) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*PushSubscription) Unsubscribe() (b bool) {
	js.Rewrite("await $<.unsubscribe()")
	return b
}

func (*PushSubscription) GetEndpoint() (endpoint string) {
	js.Rewrite("$<.endpoint")
	return endpoint
}

func (*PushSubscription) GetOptions() (options *pushsubscriptionoptions.PushSubscriptionOptions) {
	js.Rewrite("$<.options")
	return options
}
