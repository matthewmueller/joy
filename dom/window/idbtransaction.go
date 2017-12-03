package window

import (
	"github.com/matthewmueller/joy/dom/domerror"
	"github.com/matthewmueller/joy/dom/idbtransactionmode"
	"github.com/matthewmueller/joy/macro"
)

var _ EventTarget = (*IDBTransaction)(nil)

// IDBTransaction struct
// js:"IDBTransaction,omit"
type IDBTransaction struct {
}

// Abort fn
// js:"abort"
func (*IDBTransaction) Abort() {
	macro.Rewrite("$_.abort()")
}

// ObjectStore fn
// js:"objectStore"
func (*IDBTransaction) ObjectStore(name string) (i *IDBObjectStore) {
	macro.Rewrite("$_.objectStore($1)", name)
	return i
}

// AddEventListener fn
// js:"addEventListener"
func (*IDBTransaction) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*IDBTransaction) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*IDBTransaction) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Db prop
// js:"db"
func (*IDBTransaction) Db() (db *IDBDatabase) {
	macro.Rewrite("$_.db")
	return db
}

// Error prop
// js:"error"
func (*IDBTransaction) Error() (err *domerror.DOMError) {
	macro.Rewrite("$_.error")
	return err
}

// Mode prop
// js:"mode"
func (*IDBTransaction) Mode() (mode *idbtransactionmode.IDBTransactionMode) {
	macro.Rewrite("$_.mode")
	return mode
}

// Onabort prop
// js:"onabort"
func (*IDBTransaction) Onabort() (onabort func(Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*IDBTransaction) SetOnabort(onabort func(Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

// Oncomplete prop
// js:"oncomplete"
func (*IDBTransaction) Oncomplete() (oncomplete func(Event)) {
	macro.Rewrite("$_.oncomplete")
	return oncomplete
}

// SetOncomplete prop
// js:"oncomplete"
func (*IDBTransaction) SetOncomplete(oncomplete func(Event)) {
	macro.Rewrite("$_.oncomplete = $1", oncomplete)
}

// Onerror prop
// js:"onerror"
func (*IDBTransaction) Onerror() (onerror func(Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*IDBTransaction) SetOnerror(onerror func(Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}
