package datatransfer

import (
	"github.com/matthewmueller/golly/dom/datatransferitemlist"
	"github.com/matthewmueller/golly/dom/domstringlist"
	"github.com/matthewmueller/golly/dom/filelist"
	"github.com/matthewmueller/golly/js"
)

// DataTransfer struct
// js:"DataTransfer,omit"
type DataTransfer struct {
}

// ClearData fn
// js:"clearData"
func (*DataTransfer) ClearData(format *string) (b bool) {
	js.Rewrite("$_.clearData($1)", format)
	return b
}

// GetData fn
// js:"getData"
func (*DataTransfer) GetData(format string) (s string) {
	js.Rewrite("$_.getData($1)", format)
	return s
}

// SetData fn
// js:"setData"
func (*DataTransfer) SetData(format string, data string) (b bool) {
	js.Rewrite("$_.setData($1, $2)", format, data)
	return b
}

// DropEffect prop
// js:"dropEffect"
func (*DataTransfer) DropEffect() (dropEffect string) {
	js.Rewrite("$_.dropEffect")
	return dropEffect
}

// SetDropEffect prop
// js:"dropEffect"
func (*DataTransfer) SetDropEffect(dropEffect string) {
	js.Rewrite("$_.dropEffect = $1", dropEffect)
}

// EffectAllowed prop
// js:"effectAllowed"
func (*DataTransfer) EffectAllowed() (effectAllowed string) {
	js.Rewrite("$_.effectAllowed")
	return effectAllowed
}

// SetEffectAllowed prop
// js:"effectAllowed"
func (*DataTransfer) SetEffectAllowed(effectAllowed string) {
	js.Rewrite("$_.effectAllowed = $1", effectAllowed)
}

// Files prop
// js:"files"
func (*DataTransfer) Files() (files *filelist.FileList) {
	js.Rewrite("$_.files")
	return files
}

// Items prop
// js:"items"
func (*DataTransfer) Items() (items *datatransferitemlist.DataTransferItemList) {
	js.Rewrite("$_.items")
	return items
}

// Types prop
// js:"types"
func (*DataTransfer) Types() (types *domstringlist.DOMStringList) {
	js.Rewrite("$_.types")
	return types
}
