package datatransferitem

import (
	"github.com/matthewmueller/golly/dom2/file"
	"github.com/matthewmueller/golly/js"
)

// js:"DataTransferItem,omit"
type DataTransferItem struct {
}

// GetAsFile
func (*DataTransferItem) GetAsFile() (f *file.File) {
	js.Rewrite("$<.GetAsFile()")
	return f
}

// GetAsString
func (*DataTransferItem) GetAsString(_callback func(data string)) {
	js.Rewrite("$<.GetAsString($1)", _callback)
}

// WebkitGetAsEntry
func (*DataTransferItem) WebkitGetAsEntry() (i interface{}) {
	js.Rewrite("$<.WebkitGetAsEntry()")
	return i
}

// Kind
func (*DataTransferItem) Kind() (kind string) {
	js.Rewrite("$<.Kind")
	return kind
}

// Type
func (*DataTransferItem) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}
