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
// js:"item"
func (*FileList) Item(index uint) (f *file.File) {
	js.Rewrite("$_.item($1)", index)
	return f
}

// Length prop
// js:"length"
func (*FileList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
