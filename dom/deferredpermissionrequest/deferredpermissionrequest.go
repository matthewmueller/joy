package deferredpermissionrequest

import (
	"github.com/matthewmueller/golly/dom/mswebviewpermissiontype"
	"github.com/matthewmueller/golly/js"
)

// js:"DeferredPermissionRequest,omit"
type DeferredPermissionRequest interface {

	// Allow
	// js:"allow,rewrite=allow"
	Allow()

	// Deny
	// js:"deny,rewrite=deny"
	Deny()

	// ID prop
	// js:"id,rewrite=id"
	ID() (id uint)

	// Type prop
	// js:"type,rewrite=kind"
	Type() (kind *mswebviewpermissiontype.MSWebViewPermissionType)

	// URI prop
	// js:"uri,rewrite=uri"
	URI() (uri string)
}

// allow fn
func allow() {
	js.Rewrite("$<.allow()")
}

// deny fn
func deny() {
	js.Rewrite("$<.deny()")
}

// id prop
func id() (id uint) {
	js.Rewrite("$<.id")
	return id
}

// kind prop
func kind() (kind *mswebviewpermissiontype.MSWebViewPermissionType) {
	js.Rewrite("$<.type")
	return kind
}

// uri prop
func uri() (uri string) {
	js.Rewrite("$<.uri")
	return uri
}
