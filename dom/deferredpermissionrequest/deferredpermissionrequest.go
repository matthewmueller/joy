package deferredpermissionrequest

import "github.com/matthewmueller/golly/dom/mswebviewpermissiontype"

// js:"DeferredPermissionRequest,omit"
type DeferredPermissionRequest interface {

	// Allow
	// js:"allow"
	Allow()

	// Deny
	// js:"deny"
	Deny()

	// ID prop
	// js:"id"
	ID() (id uint)

	// Type prop
	// js:"type"
	Type() (kind *mswebviewpermissiontype.MSWebViewPermissionType)

	// URI prop
	// js:"uri"
	URI() (uri string)
}
