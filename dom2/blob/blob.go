package blob

// js:"Blob,omit"
type Blob interface {

	// MsClose
	MsClose()

	// MsDetachStream
	MsDetachStream() (i interface{})

	// Slice
	Slice(start *int64, end *int64, contentType *string) (b Blob)

	// Size
	Size() (size uint64)

	// Type
	Type() (kind string)
}
