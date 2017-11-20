package filelist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/file"
	"github.com/matthewmueller/golly/js"
)

type FileList struct {
}

func (*FileList) Item(index uint) (f *file.File) {
	js.Rewrite("$<.item($1)", index)
	return f
}

func (*FileList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
