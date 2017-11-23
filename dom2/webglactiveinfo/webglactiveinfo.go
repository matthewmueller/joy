package webglactiveinfo

import "github.com/matthewmueller/golly/js"

// js:"WebGLActiveInfo,omit"
type WebGLActiveInfo struct {
}

// Name
func (*WebGLActiveInfo) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Size
func (*WebGLActiveInfo) Size() (size int) {
	js.Rewrite("$<.Size")
	return size
}

// Type
func (*WebGLActiveInfo) Type() (kind uint) {
	js.Rewrite("$<.Type")
	return kind
}
