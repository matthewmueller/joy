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
// js:"add"
func (*DataTransferItemList) Add(data *file.File) (d *datatransferitem.DataTransferItem) {
	js.Rewrite("$_.add($1)", data)
	return d
}

// Clear fn
// js:"clear"
func (*DataTransferItemList) Clear() {
	js.Rewrite("$_.clear()")
}

// Item fn
// js:"item"
func (*DataTransferItemList) Item(index uint) (f *file.File) {
	js.Rewrite("$_.item($1)", index)
	return f
}

// Remove fn
// js:"remove"
func (*DataTransferItemList) Remove(index uint) {
	js.Rewrite("$_.remove($1)", index)
}

// Length prop
// js:"length"
func (*DataTransferItemList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
