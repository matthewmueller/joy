package datatransferitem

import (
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/js"
)

// DataTransferItem struct
// js:"DataTransferItem,omit"
type DataTransferItem struct {
}

// GetAsFile fn
// js:"getAsFile"
func (*DataTransferItem) GetAsFile() (f *file.File) {
	js.Rewrite("$_.getAsFile()")
	return f
}

// GetAsString fn
// js:"getAsString"
func (*DataTransferItem) GetAsString(_callback func(data string)) {
	js.Rewrite("$_.getAsString($1)", _callback)
}

// WebkitGetAsEntry fn
// js:"webkitGetAsEntry"
func (*DataTransferItem) WebkitGetAsEntry() (i interface{}) {
	js.Rewrite("$_.webkitGetAsEntry()")
	return i
}

// Kind prop
// js:"kind"
func (*DataTransferItem) Kind() (kind string) {
	js.Rewrite("$_.kind")
	return kind
}

// Type prop
// js:"type"
func (*DataTransferItem) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
