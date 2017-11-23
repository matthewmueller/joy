package file

import (
	"github.com/matthewmueller/golly/dom/blob"
	"github.com/matthewmueller/golly/js"
)

// File struct
// js:"File,omit"
type File struct {
	blob.Blob
}

// LastModifiedDate prop
func (*File) LastModifiedDate() (lastModifiedDate interface{}) {
	js.Rewrite("$<.lastModifiedDate")
	return lastModifiedDate
}

// Name prop
func (*File) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// WebkitRelativePath prop
func (*File) WebkitRelativePath() (webkitRelativePath string) {
	js.Rewrite("$<.webkitRelativePath")
	return webkitRelativePath
}
