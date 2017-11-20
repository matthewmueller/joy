package deferredpermissionrequest

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/mswebviewpermissiontype"
	"github.com/matthewmueller/golly/js"
)

type DeferredPermissionRequest struct {
}

func (*DeferredPermissionRequest) Allow() {
	js.Rewrite("$<.allow()")
}

func (*DeferredPermissionRequest) Deny() {
	js.Rewrite("$<.deny()")
}

func (*DeferredPermissionRequest) GetID() (id uint) {
	js.Rewrite("$<.id")
	return id
}

func (*DeferredPermissionRequest) GetType() (kind *mswebviewpermissiontype.MSWebViewPermissionType) {
	js.Rewrite("$<.type")
	return kind
}

func (*DeferredPermissionRequest) GetURI() (uri string) {
	js.Rewrite("$<.uri")
	return uri
}
