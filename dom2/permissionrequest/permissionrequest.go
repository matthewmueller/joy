package permissionrequest

import "github.com/matthewmueller/golly/js"

// js:"PermissionRequest,omit"
type PermissionRequest struct {
	deferredpermissionrequest.DeferredPermissionRequest
}

// Defer
func (*PermissionRequest) Defer() {
	js.Rewrite("$<.Defer()")
}

// State
func (*PermissionRequest) State() (state *mswebviewpermissionstate.MSWebViewPermissionState) {
	js.Rewrite("$<.State")
	return state
}
