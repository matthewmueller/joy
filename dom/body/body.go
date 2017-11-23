package body

import (
	"github.com/matthewmueller/golly/dom/blob"
	"github.com/matthewmueller/golly/js"
)

// Body struct
// js:"Body,omit"
type Body struct {
}

// ArrayBuffer fn
func (*Body) ArrayBuffer() (b []byte) {
	js.Rewrite("await $<.arrayBuffer()")
	return b
}

// Blob fn
func (*Body) Blob() (b blob.Blob) {
	js.Rewrite("await $<.blob()")
	return b
}

// JSON fn
func (*Body) JSON() (i interface{}) {
	js.Rewrite("await $<.json()")
	return i
}

// Text fn
func (*Body) Text() (s string) {
	js.Rewrite("await $<.text()")
	return s
}

// BodyUsed prop
func (*Body) BodyUsed() (bodyUsed bool) {
	js.Rewrite("$<.bodyUsed")
	return bodyUsed
}
