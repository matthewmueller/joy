package readablestream

import (
	"github.com/matthewmueller/golly/dom2/readablestreamreader"
	"github.com/matthewmueller/golly/js"
)

// js:"ReadableStream,omit"
type ReadableStream struct {
}

// Cancel
func (*ReadableStream) Cancel() {
	js.Rewrite("await $<.Cancel()")
}

// GetReader
func (*ReadableStream) GetReader() (r *readablestreamreader.ReadableStreamReader) {
	js.Rewrite("$<.GetReader()")
	return r
}

// Locked
func (*ReadableStream) Locked() (locked bool) {
	js.Rewrite("$<.Locked")
	return locked
}
