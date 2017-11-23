package notificationoptions

import "github.com/matthewmueller/golly/dom2/notificationdirection"

type NotificationOptions struct {
	body *string
	dir  *notificationdirection.NotificationDirection
	icon *string
	lang *string
	tag  *string
}
