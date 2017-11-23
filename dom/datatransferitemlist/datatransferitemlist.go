package datatransferitemlist

import (
	"github.com/matthewmueller/golly/dom/datatransferitem"
	"github.com/matthewmueller/golly/dom/file"
	"github.com/matthewmueller/golly/js"
)

// DataTransferItemList struct
// js:"DataTransferItemList,omit"
type DataTransferItemList struct {
}

// Add fn
func (*DataTransferItemList) Add(data *file.File) (d *datatransferitem.DataTransferItem) {
	js.Rewrite("$<.add($1)", data)
	return d
}

// Clear fn
func (*DataTransferItemList) Clear() {
	js.Rewrite("$<.clear()")
}

// Item fn
func (*DataTransferItemList) Item(index uint) (f *file.File) {
	js.Rewrite("$<.item($1)", index)
	return f
}

// Remove fn
func (*DataTransferItemList) Remove(index uint) {
	js.Rewrite("$<.remove($1)", index)
}

// Length prop
func (*DataTransferItemList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
