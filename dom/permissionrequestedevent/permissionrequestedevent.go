package permissionrequestedevent

import (
	"github.com/matthewmueller/golly/dom/permissionrequest"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// PermissionRequestedEvent struct
// js:"PermissionRequestedEvent,omit"
type PermissionRequestedEvent struct {
	window.Event
}

// PermissionRequest prop
func (*PermissionRequestedEvent) PermissionRequest() (permissionRequest *permissionrequest.PermissionRequest) {
	js.Rewrite("$<.permissionRequest")
	return permissionRequest
}
