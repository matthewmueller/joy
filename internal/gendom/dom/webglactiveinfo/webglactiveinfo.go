package webglactiveinfo

import "github.com/matthewmueller/golly/js"

type WebGLActiveInfo struct {
}

func (*WebGLActiveInfo) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*WebGLActiveInfo) GetSize() (size int) {
	js.Rewrite("$<.size")
	return size
}

func (*WebGLActiveInfo) GetType() (kind uint) {
	js.Rewrite("$<.type")
	return kind
}
