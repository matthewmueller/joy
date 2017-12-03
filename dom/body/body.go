package body

import "github.com/matthewmueller/joy/dom/blob"

// Body interface
// js:"Body"
type Body interface {

	// ArrayBuffer
	// js:"arrayBuffer"
	// jsrewrite:"await $_.arrayBuffer()"
	ArrayBuffer() (b []byte)

	// Blob
	// js:"blob"
	// jsrewrite:"await $_.blob()"
	Blob() (b blob.Blob)

	// JSON
	// js:"json"
	// jsrewrite:"await $_.json()"
	JSON() (i interface{})

	// Text
	// js:"text"
	// jsrewrite:"await $_.text()"
	Text() (s string)

	// BodyUsed prop
	// js:"bodyUsed"
	// jsrewrite:"$_.bodyUsed"
	BodyUsed() (bodyUsed bool)
}
