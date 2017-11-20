package readablestreamreader

import "github.com/matthewmueller/golly/js"

type ReadableStreamReader struct {
}

func (*ReadableStreamReader) Cancel() {
	js.Rewrite("await $<.cancel()")
}

func (*ReadableStreamReader) Read() (i interface{}) {
	js.Rewrite("await $<.read()")
	return i
}

func (*ReadableStreamReader) ReleaseLock() {
	js.Rewrite("$<.releaseLock()")
}
