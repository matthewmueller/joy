package permissionrequest

import (
	"github.com/matthewmueller/golly/dom/deferredpermissionrequest"
	"github.com/matthewmueller/golly/dom/mswebviewpermissionstate"
	"github.com/matthewmueller/golly/js"
)

// PermissionRequest struct
// js:"PermissionRequest,omit"
type PermissionRequest struct {
	deferredpermissionrequest.DeferredPermissionRequest
}

// Defer fn
func (*PermissionRequest) Defer() {
	js.Rewrite("$<.defer()")
}

// State prop
func (*PermissionRequest) State() (state *mswebviewpermissionstate.MSWebViewPermissionState) {
	js.Rewrite("$<.state")
	return state
}
