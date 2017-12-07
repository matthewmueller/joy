package datatransferitem

import (
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/macro"
)

// DataTransferItem struct
// js:"DataTransferItem,omit"
type DataTransferItem struct {
}

// GetAsFile fn
// js:"getAsFile"
func (*DataTransferItem) GetAsFile() (f *file.File) {
	macro.Rewrite("$_.getAsFile()")
	return f
}

// GetAsString fn
// js:"getAsString"
func (*DataTransferItem) GetAsString(Callback func(data string)) {
	macro.Rewrite("$_.getAsString($1)", Callback)
}

// WebkitGetAsEntry fn
// js:"webkitGetAsEntry"
func (*DataTransferItem) WebkitGetAsEntry() (i interface{}) {
	macro.Rewrite("$_.webkitGetAsEntry()")
	return i
}

// Kind prop
// js:"kind"
func (*DataTransferItem) Kind() (kind string) {
	macro.Rewrite("$_.kind")
	return kind
}

// Type prop
// js:"type"
func (*DataTransferItem) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
