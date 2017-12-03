package readablestream

import (
	"github.com/matthewmueller/joy/dom/readablestreamreader"
	"github.com/matthewmueller/joy/js"
)

// ReadableStream struct
// js:"ReadableStream,omit"
type ReadableStream struct {
}

// Cancel fn
// js:"cancel"
func (*ReadableStream) Cancel() {
	js.Rewrite("await $_.cancel()")
}

// GetReader fn
// js:"getReader"
func (*ReadableStream) GetReader() (r *readablestreamreader.ReadableStreamReader) {
	js.Rewrite("$_.getReader()")
	return r
}

// Locked prop
// js:"locked"
func (*ReadableStream) Locked() (locked bool) {
	js.Rewrite("$_.locked")
	return locked
}
