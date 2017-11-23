package datatransferitem

import (
	"github.com/matthewmueller/golly/dom2/file"
	"github.com/matthewmueller/golly/js"
)

// DataTransferItem struct
// js:"DataTransferItem,omit"
type DataTransferItem struct {
}

// GetAsFile fn
func (*DataTransferItem) GetAsFile() (f *file.File) {
	js.Rewrite("$<.getAsFile()")
	return f
}

// GetAsString fn
func (*DataTransferItem) GetAsString(_callback func(data string)) {
	js.Rewrite("$<.getAsString($1)", _callback)
}

// WebkitGetAsEntry fn
func (*DataTransferItem) WebkitGetAsEntry() (i interface{}) {
	js.Rewrite("$<.webkitGetAsEntry()")
	return i
}

// Kind prop
func (*DataTransferItem) Kind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

// Type prop
func (*DataTransferItem) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
