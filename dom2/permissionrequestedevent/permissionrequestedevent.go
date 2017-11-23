package permissionrequestedevent

import "github.com/matthewmueller/golly/js"

// js:"PermissionRequestedEvent,omit"
type PermissionRequestedEvent struct {
	window.Event
}

// PermissionRequest
func (*PermissionRequestedEvent) PermissionRequest() (permissionRequest *permissionrequest.PermissionRequest) {
	js.Rewrite("$<.PermissionRequest")
	return permissionRequest
}
