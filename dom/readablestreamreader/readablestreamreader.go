package readablestreamreader

import "github.com/matthewmueller/joy/macro"

// ReadableStreamReader struct
// js:"ReadableStreamReader,omit"
type ReadableStreamReader struct {
}

// Cancel fn
// js:"cancel"
func (*ReadableStreamReader) Cancel() {
	macro.Rewrite("await $_.cancel()")
}

// Read fn
// js:"read"
func (*ReadableStreamReader) Read() (i interface{}) {
	macro.Rewrite("await $_.read()")
	return i
}

// ReleaseLock fn
// js:"releaseLock"
func (*ReadableStreamReader) ReleaseLock() {
	macro.Rewrite("$_.releaseLock()")
}
