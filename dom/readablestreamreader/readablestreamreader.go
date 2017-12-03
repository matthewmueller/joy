package readablestreamreader

import "github.com/matthewmueller/joy/js"

// ReadableStreamReader struct
// js:"ReadableStreamReader,omit"
type ReadableStreamReader struct {
}

// Cancel fn
// js:"cancel"
func (*ReadableStreamReader) Cancel() {
	js.Rewrite("await $_.cancel()")
}

// Read fn
// js:"read"
func (*ReadableStreamReader) Read() (i interface{}) {
	js.Rewrite("await $_.read()")
	return i
}

// ReleaseLock fn
// js:"releaseLock"
func (*ReadableStreamReader) ReleaseLock() {
	js.Rewrite("$_.releaseLock()")
}
