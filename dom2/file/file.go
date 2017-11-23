package file

import (
	"github.com/matthewmueller/golly/dom2/blob"
	"github.com/matthewmueller/golly/js"
)

// File struct
// js:"File,omit"
type File struct {
	blob.Blob
}

// LastModifiedDate
func (*File) LastModifiedDate() (lastModifiedDate interface{}) {
	js.Rewrite("$<.LastModifiedDate")
	return lastModifiedDate
}

// Name
func (*File) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// WebkitRelativePath
func (*File) WebkitRelativePath() (webkitRelativePath string) {
	js.Rewrite("$<.WebkitRelativePath")
	return webkitRelativePath
}
