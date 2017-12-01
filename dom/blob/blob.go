package blob

// Blob interface
// js:"Blob"
type Blob interface {

	// MsClose
	// js:"msClose"
	// jsrewrite:"$_.msClose()"
	MsClose()

	// MsDetachStream
	// js:"msDetachStream"
	// jsrewrite:"$_.msDetachStream()"
	MsDetachStream() (i interface{})

	// Slice
	// js:"slice"
	// jsrewrite:"$_.slice($1, $2, $3)"
	Slice(start *int64, end *int64, contentType *string) (b Blob)

	// Size prop
	// js:"size"
	// jsrewrite:"$_.size"
	Size() (size uint64)

	// Type prop
	// js:"type"
	// jsrewrite:"$_.type"
	Type() (kind string)
}
