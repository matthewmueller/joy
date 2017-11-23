package deferredpermissionrequest

import "github.com/matthewmueller/golly/dom2/mswebviewpermissiontype"

// js:"DeferredPermissionRequest,omit"
type DeferredPermissionRequest interface {

	// Allow
	Allow()

	// Deny
	Deny()

	// ID
	ID() (id uint)

	// Type
	Type() (kind *mswebviewpermissiontype.MSWebViewPermissionType)

	// URI
	URI() (uri string)
}
