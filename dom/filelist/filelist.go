package filelist

import (
	"github.com/matthewmueller/golly/dom/file"
	"github.com/matthewmueller/golly/js"
)

// FileList struct
// js:"FileList,omit"
type FileList struct {
}

// Item fn
func (*FileList) Item(index uint) (f *file.File) {
	js.Rewrite("$<.item($1)", index)
	return f
}

// Length prop
func (*FileList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
