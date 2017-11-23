package permissionrequestedevent

import (
	"github.com/matthewmueller/golly/dom2/permissionrequest"
	"github.com/matthewmueller/golly/dom2/window"
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
