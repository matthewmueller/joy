package webglactiveinfo

import "github.com/matthewmueller/joy/macro"

// WebGLActiveInfo struct
// js:"WebGLActiveInfo,omit"
type WebGLActiveInfo struct {
}

// Name prop
// js:"name"
func (*WebGLActiveInfo) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// Size prop
// js:"size"
func (*WebGLActiveInfo) Size() (size int) {
	macro.Rewrite("$_.size")
	return size
}

// Type prop
// js:"type"
func (*WebGLActiveInfo) Type() (kind uint) {
	macro.Rewrite("$_.type")
	return kind
}
