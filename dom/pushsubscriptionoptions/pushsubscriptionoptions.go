package pushsubscriptionoptions

import "github.com/matthewmueller/golly/js"

// PushSubscriptionOptions struct
// js:"PushSubscriptionOptions,omit"
type PushSubscriptionOptions struct {
}

// ApplicationServerKey prop
// js:"applicationServerKey"
func (*PushSubscriptionOptions) ApplicationServerKey() (applicationServerKey []byte) {
	js.Rewrite("$_.applicationServerKey")
	return applicationServerKey
}

// UserVisibleOnly prop
// js:"userVisibleOnly"
func (*PushSubscriptionOptions) UserVisibleOnly() (userVisibleOnly bool) {
	js.Rewrite("$_.userVisibleOnly")
	return userVisibleOnly
}
