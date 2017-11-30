package blob

// Blob interface
// js:"Blob"
type Blob interface {

	// MsClose
	// js:"msClose"
	MsClose()

	// MsDetachStream
	// js:"msDetachStream"
	MsDetachStream() (i interface{})

	// Slice
	// js:"slice"
	Slice(start *int64, end *int64, contentType *string) (b Blob)

	// Size prop
	// js:"size"
	Size() (size uint64)

	// Type prop
	// js:"type"
	Type() (kind string)
}
