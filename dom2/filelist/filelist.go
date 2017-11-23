package filelist

import (
	"github.com/matthewmueller/golly/dom2/file"
	"github.com/matthewmueller/golly/js"
)

// js:"FileList,omit"
type FileList struct {
}

// Item
func (*FileList) Item(index uint) (f *file.File) {
	js.Rewrite("$<.Item($1)", index)
	return f
}

// Length
func (*FileList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
