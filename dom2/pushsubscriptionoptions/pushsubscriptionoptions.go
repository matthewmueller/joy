package pushsubscriptionoptions

import "github.com/matthewmueller/golly/js"

// js:"PushSubscriptionOptions,omit"
type PushSubscriptionOptions struct {
}

// ApplicationServerKey
func (*PushSubscriptionOptions) ApplicationServerKey() (applicationServerKey []byte) {
	js.Rewrite("$<.ApplicationServerKey")
	return applicationServerKey
}

// UserVisibleOnly
func (*PushSubscriptionOptions) UserVisibleOnly() (userVisibleOnly bool) {
	js.Rewrite("$<.UserVisibleOnly")
	return userVisibleOnly
}
