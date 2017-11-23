package webglactiveinfo

import "github.com/matthewmueller/golly/js"

// WebGLActiveInfo struct
// js:"WebGLActiveInfo,omit"
type WebGLActiveInfo struct {
}

// Name prop
func (*WebGLActiveInfo) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Size prop
func (*WebGLActiveInfo) Size() (size int) {
	js.Rewrite("$<.size")
	return size
}

// Type prop
func (*WebGLActiveInfo) Type() (kind uint) {
	js.Rewrite("$<.type")
	return kind
}
