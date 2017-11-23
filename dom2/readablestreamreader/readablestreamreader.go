package readablestreamreader

import "github.com/matthewmueller/golly/js"

// js:"ReadableStreamReader,omit"
type ReadableStreamReader struct {
}

// Cancel
func (*ReadableStreamReader) Cancel() {
	js.Rewrite("await $<.Cancel()")
}

// Read
func (*ReadableStreamReader) Read() (i interface{}) {
	js.Rewrite("await $<.Read()")
	return i
}

// ReleaseLock
func (*ReadableStreamReader) ReleaseLock() {
	js.Rewrite("$<.ReleaseLock()")
}
