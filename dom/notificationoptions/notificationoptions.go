package notificationoptions

import "github.com/matthewmueller/joy/dom/notificationdirection"

type NotificationOptions struct {
	body *string
	dir  *notificationdirection.NotificationDirection
	icon *string
	lang *string
	tag  *string
}
