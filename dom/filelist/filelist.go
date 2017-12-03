package filelist

import (
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/macro"
)

// FileList struct
// js:"FileList,omit"
type FileList struct {
}

// Item fn
// js:"item"
func (*FileList) Item(index uint) (f *file.File) {
	macro.Rewrite("$_.item($1)", index)
	return f
}

// Length prop
// js:"length"
func (*FileList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
