package datatransfer

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/datatransferitemlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/domstringlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/filelist"
	"github.com/matthewmueller/golly/js"
)

type DataTransfer struct {
}

func (*DataTransfer) ClearData(format string) (b bool) {
	js.Rewrite("$<.clearData($1)", format)
	return b
}

func (*DataTransfer) GetData(format string) (s string) {
	js.Rewrite("$<.getData($1)", format)
	return s
}

func (*DataTransfer) SetData(format string, data string) (b bool) {
	js.Rewrite("$<.setData($1, $2)", format, data)
	return b
}

func (*DataTransfer) GetDropEffect() (dropEffect string) {
	js.Rewrite("$<.dropEffect")
	return dropEffect
}

func (*DataTransfer) SetDropEffect(dropEffect string) {
	js.Rewrite("$<.dropEffect = $1", dropEffect)
}

func (*DataTransfer) GetEffectAllowed() (effectAllowed string) {
	js.Rewrite("$<.effectAllowed")
	return effectAllowed
}

func (*DataTransfer) SetEffectAllowed(effectAllowed string) {
	js.Rewrite("$<.effectAllowed = $1", effectAllowed)
}

func (*DataTransfer) GetFiles() (files *filelist.FileList) {
	js.Rewrite("$<.files")
	return files
}

func (*DataTransfer) GetItems() (items *datatransferitemlist.DataTransferItemList) {
	js.Rewrite("$<.items")
	return items
}

func (*DataTransfer) GetTypes() (types *domstringlist.DOMStringList) {
	js.Rewrite("$<.types")
	return types
}
