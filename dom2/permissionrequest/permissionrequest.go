package permissionrequest

import (
	"github.com/matthewmueller/golly/dom2/deferredpermissionrequest"
	"github.com/matthewmueller/golly/dom2/mswebviewpermissionstate"
	"github.com/matthewmueller/golly/js"
)

// PermissionRequest struct
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
