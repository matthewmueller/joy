package body

import "github.com/matthewmueller/golly/dom/blob"

// Body interface
// js:"Body"
type Body interface {

	// ArrayBuffer
	// js:"arrayBuffer"
	ArrayBuffer() (b []byte)

	// Blob
	// js:"blob"
	Blob() (b blob.Blob)

	// JSON
	// js:"json"
	JSON() (i interface{})

	// Text
	// js:"text"
	Text() (s string)

	// BodyUsed prop
	// js:"bodyUsed"
	BodyUsed() (bodyUsed bool)
}
