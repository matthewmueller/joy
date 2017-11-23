package datatransfer

import (
	"github.com/matthewmueller/golly/dom2/datatransferitemlist"
	"github.com/matthewmueller/golly/dom2/domstringlist"
	"github.com/matthewmueller/golly/dom2/filelist"
	"github.com/matthewmueller/golly/js"
)

// DataTransfer struct
// js:"DataTransfer,omit"
type DataTransfer struct {
}

// ClearData fn
func (*DataTransfer) ClearData(format *string) (b bool) {
	js.Rewrite("$<.clearData($1)", format)
	return b
}

// GetData fn
func (*DataTransfer) GetData(format string) (s string) {
	js.Rewrite("$<.getData($1)", format)
	return s
}

// SetData fn
func (*DataTransfer) SetData(format string, data string) (b bool) {
	js.Rewrite("$<.setData($1, $2)", format, data)
	return b
}

// DropEffect prop
func (*DataTransfer) DropEffect() (dropEffect string) {
	js.Rewrite("$<.dropEffect")
	return dropEffect
}

// DropEffect prop
func (*DataTransfer) SetDropEffect(dropEffect string) {
	js.Rewrite("$<.dropEffect = $1", dropEffect)
}

// EffectAllowed prop
func (*DataTransfer) EffectAllowed() (effectAllowed string) {
	js.Rewrite("$<.effectAllowed")
	return effectAllowed
}

// EffectAllowed prop
func (*DataTransfer) SetEffectAllowed(effectAllowed string) {
	js.Rewrite("$<.effectAllowed = $1", effectAllowed)
}

// Files prop
func (*DataTransfer) Files() (files *filelist.FileList) {
	js.Rewrite("$<.files")
	return files
}

// Items prop
func (*DataTransfer) Items() (items *datatransferitemlist.DataTransferItemList) {
	js.Rewrite("$<.items")
	return items
}

// Types prop
func (*DataTransfer) Types() (types *domstringlist.DOMStringList) {
	js.Rewrite("$<.types")
	return types
}
