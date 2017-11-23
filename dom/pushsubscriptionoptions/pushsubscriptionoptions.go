package pushsubscriptionoptions

import "github.com/matthewmueller/golly/js"

// PushSubscriptionOptions struct
// js:"PushSubscriptionOptions,omit"
type PushSubscriptionOptions struct {
}

// ApplicationServerKey prop
func (*PushSubscriptionOptions) ApplicationServerKey() (applicationServerKey []byte) {
	js.Rewrite("$<.applicationServerKey")
	return applicationServerKey
}

// UserVisibleOnly prop
func (*PushSubscriptionOptions) UserVisibleOnly() (userVisibleOnly bool) {
	js.Rewrite("$<.userVisibleOnly")
	return userVisibleOnly
}
