package readablestreamreader

import "github.com/matthewmueller/golly/js"

// ReadableStreamReader struct
// js:"ReadableStreamReader,omit"
type ReadableStreamReader struct {
}

// Cancel fn
func (*ReadableStreamReader) Cancel() {
	js.Rewrite("await $<.cancel()")
}

// Read fn
func (*ReadableStreamReader) Read() (i interface{}) {
	js.Rewrite("await $<.read()")
	return i
}

// ReleaseLock fn
func (*ReadableStreamReader) ReleaseLock() {
	js.Rewrite("$<.releaseLock()")
}
