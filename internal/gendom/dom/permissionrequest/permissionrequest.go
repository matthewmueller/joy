package permissionrequest

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/deferredpermissionrequest"
	"github.com/matthewmueller/golly/internal/gendom/dom/mswebviewpermissionstate"
	"github.com/matthewmueller/golly/js"
)

type PermissionRequest struct {
	*deferredpermissionrequest.DeferredPermissionRequest
}

func (*PermissionRequest) Defer() {
	js.Rewrite("$<.defer()")
}

func (*PermissionRequest) GetState() (state *mswebviewpermissionstate.MSWebViewPermissionState) {
	js.Rewrite("$<.state")
	return state
}
