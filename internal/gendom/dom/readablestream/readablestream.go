package readablestream

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/readablestreamreader"
	"github.com/matthewmueller/golly/js"
)

type ReadableStream struct {
}

func (*ReadableStream) Cancel() {
	js.Rewrite("await $<.cancel()")
}

func (*ReadableStream) GetReader() (r *readablestreamreader.ReadableStreamReader) {
	js.Rewrite("$<.getReader()")
	return r
}

func (*ReadableStream) GetLocked() (locked bool) {
	js.Rewrite("$<.locked")
	return locked
}
