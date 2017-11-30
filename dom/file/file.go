package file

import (
	"github.com/matthewmueller/golly/dom/blob"
	"github.com/matthewmueller/golly/js"
)

var _ blob.Blob = (*File)(nil)

// File struct
// js:"File,omit"
type File struct {
}

// Slice fn
// js:"slice"
func (*File) Slice(start *int64, end *int64, contentType *string) (b blob.Blob) {
	js.Rewrite("$_.slice($1, $2, $3)", start, end, contentType)
	return b
}

// LastModifiedDate prop
// js:"lastModifiedDate"
func (*File) LastModifiedDate() (lastModifiedDate interface{}) {
	js.Rewrite("$_.lastModifiedDate")
	return lastModifiedDate
}

// Name prop
// js:"name"
func (*File) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// WebkitRelativePath prop
// js:"webkitRelativePath"
func (*File) WebkitRelativePath() (webkitRelativePath string) {
	js.Rewrite("$_.webkitRelativePath")
	return webkitRelativePath
}

// Size prop
// js:"size"
func (*File) Size() (size uint64) {
	js.Rewrite("$_.size")
	return size
}

// Type prop
// js:"type"
func (*File) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
