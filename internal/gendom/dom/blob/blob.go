package blob

import "github.com/matthewmueller/golly/js"

type Blob struct {
}

func (*Blob) MsClose() {
	js.Rewrite("$<.msClose()")
}

func (*Blob) MsDetachStream() (i interface{}) {
	js.Rewrite("$<.msDetachStream()")
	return i
}

func (*Blob) Slice(start int64, end int64, contentType string) (b *Blob) {
	js.Rewrite("$<.slice($1, $2, $3)", start, end, contentType)
	return b
}

func (*Blob) GetSize() (size uint64) {
	js.Rewrite("$<.size")
	return size
}

func (*Blob) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
