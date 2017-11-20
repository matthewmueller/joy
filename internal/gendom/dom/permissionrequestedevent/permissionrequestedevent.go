package permissionrequestedevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/permissionrequest"
	"github.com/matthewmueller/golly/js"
)

type PermissionRequestedEvent struct {
	*event.Event
}

func (*PermissionRequestedEvent) GetPermissionRequest() (permissionRequest *permissionrequest.PermissionRequest) {
	js.Rewrite("$<.permissionRequest")
	return permissionRequest
}
