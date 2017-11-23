package datatransfer

import (
	"github.com/matthewmueller/golly/dom2/domstringlist"
	"github.com/matthewmueller/golly/js"
)

// js:"DataTransfer,omit"
type DataTransfer struct {
}

// ClearData
func (*DataTransfer) ClearData(format *string) (b bool) {
	js.Rewrite("$<.ClearData($1)", format)
	return b
}

// GetData
func (*DataTransfer) GetData(format string) (s string) {
	js.Rewrite("$<.GetData($1)", format)
	return s
}

// SetData
func (*DataTransfer) SetData(format string, data string) (b bool) {
	js.Rewrite("$<.SetData($1, $2)", format, data)
	return b
}

// DropEffect
func (*DataTransfer) DropEffect() (dropEffect string) {
	js.Rewrite("$<.DropEffect")
	return dropEffect
}

// DropEffect
func (*DataTransfer) SetDropEffect(dropEffect string) {
	js.Rewrite("$<.DropEffect = $1", dropEffect)
}

// EffectAllowed
func (*DataTransfer) EffectAllowed() (effectAllowed string) {
	js.Rewrite("$<.EffectAllowed")
	return effectAllowed
}

// EffectAllowed
func (*DataTransfer) SetEffectAllowed(effectAllowed string) {
	js.Rewrite("$<.EffectAllowed = $1", effectAllowed)
}

// Files
func (*DataTransfer) Files() (files *filelist.FileList) {
	js.Rewrite("$<.Files")
	return files
}

// Items
func (*DataTransfer) Items() (items *datatransferitemlist.DataTransferItemList) {
	js.Rewrite("$<.Items")
	return items
}

// Types
func (*DataTransfer) Types() (types *domstringlist.DOMStringList) {
	js.Rewrite("$<.Types")
	return types
}
