package pushsubscriptionoptions

import "github.com/matthewmueller/golly/js"

type PushSubscriptionOptions struct {
}

func (*PushSubscriptionOptions) GetApplicationServerKey() (applicationServerKey []byte) {
	js.Rewrite("$<.applicationServerKey")
	return applicationServerKey
}

func (*PushSubscriptionOptions) GetUserVisibleOnly() (userVisibleOnly bool) {
	js.Rewrite("$<.userVisibleOnly")
	return userVisibleOnly
}
