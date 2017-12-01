package permissionrequest

import (
	"github.com/matthewmueller/golly/dom/deferredpermissionrequest"
	"github.com/matthewmueller/golly/dom/mswebviewpermissionstate"
	"github.com/matthewmueller/golly/dom/mswebviewpermissiontype"
	"github.com/matthewmueller/golly/js"
)

var _ deferredpermissionrequest.DeferredPermissionRequest = (*PermissionRequest)(nil)

// PermissionRequest struct
// js:"PermissionRequest,omit"
type PermissionRequest struct {
}

// Defer fn
// js:"defer"
func (*PermissionRequest) Defer() {
	js.Rewrite("$_.defer()")
}

// Allow fn
// js:"allow"
func (*PermissionRequest) Allow() {
	js.Rewrite("$_.allow()")
}

// Deny fn
// js:"deny"
func (*PermissionRequest) Deny() {
	js.Rewrite("$_.deny()")
}

// State prop
// js:"state"
func (*PermissionRequest) State() (state *mswebviewpermissionstate.MSWebViewPermissionState) {
	js.Rewrite("$_.state")
	return state
}

// ID prop
// js:"id"
func (*PermissionRequest) ID() (id uint) {
	js.Rewrite("$_.id")
	return id
}

// Type prop
// js:"type"
func (*PermissionRequest) Type() (kind *mswebviewpermissiontype.MSWebViewPermissionType) {
	js.Rewrite("$_.type")
	return kind
}

// URI prop
// js:"uri"
func (*PermissionRequest) URI() (uri string) {
	js.Rewrite("$_.uri")
	return uri
}
