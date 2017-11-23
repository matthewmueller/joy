package window

import (
	"github.com/matthewmueller/golly/dom2/domerror"
	"github.com/matthewmueller/golly/dom2/idbtransactionmode"
	"github.com/matthewmueller/golly/js"
)

// js:"IDBTransaction,omit"
type IDBTransaction struct {
	EventTarget
}

// Abort
func (*IDBTransaction) Abort() {
	js.Rewrite("$<.Abort()")
}

// ObjectStore
func (*IDBTransaction) ObjectStore(name string) (i *IDBObjectStore) {
	js.Rewrite("$<.ObjectStore($1)", name)
	return i
}

// Db
func (*IDBTransaction) Db() (db *IDBDatabase) {
	js.Rewrite("$<.Db")
	return db
}

// Error
func (*IDBTransaction) Error() (err *domerror.DOMError) {
	js.Rewrite("$<.Error")
	return err
}

// Mode
func (*IDBTransaction) Mode() (mode *idbtransactionmode.IDBTransactionMode) {
	js.Rewrite("$<.Mode")
	return mode
}

// Onabort
func (*IDBTransaction) Onabort() (onabort func(Event)) {
	js.Rewrite("$<.Onabort")
	return onabort
}

// Onabort
func (*IDBTransaction) SetOnabort(onabort func(Event)) {
	js.Rewrite("$<.Onabort = $1", onabort)
}

// Oncomplete
func (*IDBTransaction) Oncomplete() (oncomplete func(Event)) {
	js.Rewrite("$<.Oncomplete")
	return oncomplete
}

// Oncomplete
func (*IDBTransaction) SetOncomplete(oncomplete func(Event)) {
	js.Rewrite("$<.Oncomplete = $1", oncomplete)
}

// Onerror
func (*IDBTransaction) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*IDBTransaction) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}
