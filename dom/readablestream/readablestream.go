package readablestream

import (
	"github.com/matthewmueller/joy/dom/readablestreamreader"
	"github.com/matthewmueller/joy/macro"
)

// ReadableStream struct
// js:"ReadableStream,omit"
type ReadableStream struct {
}

// Cancel fn
// js:"cancel"
func (*ReadableStream) Cancel() {
	macro.Rewrite("await $_.cancel()")
}

// GetReader fn
// js:"getReader"
func (*ReadableStream) GetReader() (r *readablestreamreader.ReadableStreamReader) {
	macro.Rewrite("$_.getReader()")
	return r
}

// Locked prop
// js:"locked"
func (*ReadableStream) Locked() (locked bool) {
	macro.Rewrite("$_.locked")
	return locked
}
