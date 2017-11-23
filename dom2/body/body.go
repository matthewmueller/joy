package body

import "github.com/matthewmueller/golly/dom2/blob"

// js:"Body,omit"
type Body interface {

	// ArrayBuffer
	ArrayBuffer() (b []byte)

	// Blob
	Blob() (b blob.Blob)

	// JSON
	JSON() (i interface{})

	// Text
	Text() (s string)

	// BodyUsed
	BodyUsed() (bodyUsed bool)
}
