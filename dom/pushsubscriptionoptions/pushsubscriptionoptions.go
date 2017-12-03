package pushsubscriptionoptions

import "github.com/matthewmueller/joy/macro"

// PushSubscriptionOptions struct
// js:"PushSubscriptionOptions,omit"
type PushSubscriptionOptions struct {
}

// ApplicationServerKey prop
// js:"applicationServerKey"
func (*PushSubscriptionOptions) ApplicationServerKey() (applicationServerKey []byte) {
	macro.Rewrite("$_.applicationServerKey")
	return applicationServerKey
}

// UserVisibleOnly prop
// js:"userVisibleOnly"
func (*PushSubscriptionOptions) UserVisibleOnly() (userVisibleOnly bool) {
	macro.Rewrite("$_.userVisibleOnly")
	return userVisibleOnly
}
