package notificationoptions

import "github.com/matthewmueller/golly/internal/gendom/dom/notificationdirection"

type NotificationOptions struct {
	body *string
	dir  *notificationdirection.NotificationDirection
	icon *string
	lang *string
	tag  *string
}
