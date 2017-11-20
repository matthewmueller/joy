package datatransferitemlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/datatransferitem"
	"github.com/matthewmueller/golly/internal/gendom/dom/file"
	"github.com/matthewmueller/golly/js"
)

type DataTransferItemList struct {
}

func (*DataTransferItemList) Add(data *file.File) (d *datatransferitem.DataTransferItem) {
	js.Rewrite("$<.add($1)", data)
	return d
}

func (*DataTransferItemList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*DataTransferItemList) Item(index uint) (f *file.File) {
	js.Rewrite("$<.item($1)", index)
	return f
}

func (*DataTransferItemList) Remove(index uint) {
	js.Rewrite("$<.remove($1)", index)
}

func (*DataTransferItemList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}
