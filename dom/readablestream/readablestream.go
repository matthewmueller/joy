package readablestream

import (
	"github.com/matthewmueller/golly/dom2/readablestreamreader"
	"github.com/matthewmueller/golly/js"
)

// ReadableStream struct
// js:"ReadableStream,omit"
type ReadableStream struct {
}

// Cancel fn
func (*ReadableStream) Cancel() {
	js.Rewrite("await $<.cancel()")
}

// GetReader fn
func (*ReadableStream) GetReader() (r *readablestreamreader.ReadableStreamReader) {
	js.Rewrite("$<.getReader()")
	return r
}

// Locked prop
func (*ReadableStream) Locked() (locked bool) {
	js.Rewrite("$<.locked")
	return locked
}
