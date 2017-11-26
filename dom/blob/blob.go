package blob

import "github.com/matthewmueller/golly/js"

// js:"Blob,omit"
type Blob interface {

	// MsClose
	// js:"msClose,rewrite=msclose"
	MsClose()

	// MsDetachStream
	// js:"msDetachStream,rewrite=msdetachstream"
	MsDetachStream() (i interface{})

	// Slice
	// js:"slice,rewrite=slice"
	Slice(start *int64, end *int64, contentType *string) (b Blob)

	// Size prop
	// js:"size,rewrite=size"
	Size() (size uint64)

	// Type prop
	// js:"type,rewrite=kind"
	Type() (kind string)
}

// msclose fn
func msclose() {
	js.Rewrite("$<.msClose()")
}

// msdetachstream fn
func msdetachstream() (i interface{}) {
	js.Rewrite("$<.msDetachStream()")
	return i
}

// slice fn
func slice(start *int64, end *int64, contentType *string) (b Blob) {
	js.Rewrite("$<.slice($1, $2, $3)", start, end, contentType)
	return b
}

// size prop
func size() (size uint64) {
	js.Rewrite("$<.size")
	return size
}

// kind prop
func kind() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
