package body

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/blob"
	"github.com/matthewmueller/golly/js"
)

type Body struct {
}

func (*Body) ArrayBuffer() (b []byte) {
	js.Rewrite("await $<.arrayBuffer()")
	return b
}

func (*Body) Blob() (b *blob.Blob) {
	js.Rewrite("await $<.blob()")
	return b
}

func (*Body) JSON() (i interface{}) {
	js.Rewrite("await $<.json()")
	return i
}

func (*Body) Text() (s string) {
	js.Rewrite("await $<.text()")
	return s
}

func (*Body) GetBodyUsed() (bodyUsed bool) {
	js.Rewrite("$<.bodyUsed")
	return bodyUsed
}
