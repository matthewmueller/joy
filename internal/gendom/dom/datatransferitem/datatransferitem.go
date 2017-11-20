package datatransferitem

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/file"
	"github.com/matthewmueller/golly/js"
)

type DataTransferItem struct {
}

func (*DataTransferItem) GetAsFile() (f *file.File) {
	js.Rewrite("$<.getAsFile()")
	return f
}

func (*DataTransferItem) GetAsString(_callback func(data string)) {
	js.Rewrite("$<.getAsString($1)", _callback)
}

func (*DataTransferItem) WebkitGetAsEntry() (i interface{}) {
	js.Rewrite("$<.webkitGetAsEntry()")
	return i
}

func (*DataTransferItem) GetKind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

func (*DataTransferItem) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}
