package datatransferitemlist

import (
	"github.com/matthewmueller/golly/dom2/datatransferitem"
	"github.com/matthewmueller/golly/dom2/file"
	"github.com/matthewmueller/golly/js"
)

// js:"DataTransferItemList,omit"
type DataTransferItemList struct {
}

// Add
func (*DataTransferItemList) Add(data *file.File) (d *datatransferitem.DataTransferItem) {
	js.Rewrite("$<.Add($1)", data)
	return d
}

// Clear
func (*DataTransferItemList) Clear() {
	js.Rewrite("$<.Clear()")
}

// Item
func (*DataTransferItemList) Item(index uint) (f *file.File) {
	js.Rewrite("$<.Item($1)", index)
	return f
}

// Remove
func (*DataTransferItemList) Remove(index uint) {
	js.Rewrite("$<.Remove($1)", index)
}

// Length
func (*DataTransferItemList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
