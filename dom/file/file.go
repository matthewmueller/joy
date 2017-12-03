package file

import (
	"github.com/matthewmueller/joy/dom/blob"
	"github.com/matthewmueller/joy/macro"
)

var _ blob.Blob = (*File)(nil)

// File struct
// js:"File,omit"
type File struct {
}

// MsClose fn
// js:"msClose"
func (*File) MsClose() {
	macro.Rewrite("$_.msClose()")
}

// MsDetachStream fn
// js:"msDetachStream"
func (*File) MsDetachStream() (i interface{}) {
	macro.Rewrite("$_.msDetachStream()")
	return i
}

// Slice fn
// js:"slice"
func (*File) Slice(start *int64, end *int64, contentType *string) (b blob.Blob) {
	macro.Rewrite("$_.slice($1, $2, $3)", start, end, contentType)
	return b
}

// LastModifiedDate prop
// js:"lastModifiedDate"
func (*File) LastModifiedDate() (lastModifiedDate interface{}) {
	macro.Rewrite("$_.lastModifiedDate")
	return lastModifiedDate
}

// Name prop
// js:"name"
func (*File) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// WebkitRelativePath prop
// js:"webkitRelativePath"
func (*File) WebkitRelativePath() (webkitRelativePath string) {
	macro.Rewrite("$_.webkitRelativePath")
	return webkitRelativePath
}

// Size prop
// js:"size"
func (*File) Size() (size uint64) {
	macro.Rewrite("$_.size")
	return size
}

// Type prop
// js:"type"
func (*File) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
