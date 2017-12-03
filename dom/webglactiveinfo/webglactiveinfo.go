package webglactiveinfo

import "github.com/matthewmueller/joy/js"

// WebGLActiveInfo struct
// js:"WebGLActiveInfo,omit"
type WebGLActiveInfo struct {
}

// Name prop
// js:"name"
func (*WebGLActiveInfo) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// Size prop
// js:"size"
func (*WebGLActiveInfo) Size() (size int) {
	js.Rewrite("$_.size")
	return size
}

// Type prop
// js:"type"
func (*WebGLActiveInfo) Type() (kind uint) {
	js.Rewrite("$_.type")
	return kind
}
