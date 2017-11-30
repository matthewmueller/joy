package blob

// Blob interface
// js:"Blob"
type Blob interface {

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
