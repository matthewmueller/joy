package datatransferitemlist

import (
	"github.com/matthewmueller/joy/dom/datatransferitem"
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/macro"
)

// DataTransferItemList struct
// js:"DataTransferItemList,omit"
type DataTransferItemList struct {
}

// Add fn
// js:"add"
func (*DataTransferItemList) Add(data *file.File) (d *datatransferitem.DataTransferItem) {
	macro.Rewrite("$_.add($1)", data)
	return d
}

// Clear fn
// js:"clear"
func (*DataTransferItemList) Clear() {
	macro.Rewrite("$_.clear()")
}

// Item fn
// js:"item"
func (*DataTransferItemList) Item(index uint) (f *file.File) {
	macro.Rewrite("$_.item($1)", index)
	return f
}

// Remove fn
// js:"remove"
func (*DataTransferItemList) Remove(index uint) {
	macro.Rewrite("$_.remove($1)", index)
}

// Length prop
// js:"length"
func (*DataTransferItemList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
