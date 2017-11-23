package body

import (
	"github.com/matthewmueller/golly/dom2/blob"
	"github.com/matthewmueller/golly/js"
)

// Body struct
// js:"Body,omit"
type Body struct {
}

// ArrayBuffer
func (*Body) ArrayBuffer() (b []byte) {
	js.Rewrite("await $<.ArrayBuffer()")
	return b
}

// Blob
func (*Body) Blob() (b blob.Blob) {
	js.Rewrite("await $<.Blob()")
	return b
}

// JSON
func (*Body) JSON() (i interface{}) {
	js.Rewrite("await $<.JSON()")
	return i
}

// Text
func (*Body) Text() (s string) {
	js.Rewrite("await $<.Text()")
	return s
}

// BodyUsed
func (*Body) BodyUsed() (bodyUsed bool) {
	js.Rewrite("$<.BodyUsed")
	return bodyUsed
}
