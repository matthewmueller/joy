package window

import (
	"github.com/matthewmueller/golly/dom/domerror"
	"github.com/matthewmueller/golly/dom/idbtransactionmode"
	"github.com/matthewmueller/golly/js"
)

// IDBTransaction struct
// js:"IDBTransaction,omit"
type IDBTransaction struct {
	EventTarget
}

// Abort fn
func (*IDBTransaction) Abort() {
	js.Rewrite("$<.abort()")
}

// ObjectStore fn
func (*IDBTransaction) ObjectStore(name string) (i *IDBObjectStore) {
	js.Rewrite("$<.objectStore($1)", name)
	return i
}

// Db prop
func (*IDBTransaction) Db() (db *IDBDatabase) {
	js.Rewrite("$<.db")
	return db
}

// Error prop
func (*IDBTransaction) Error() (err *domerror.DOMError) {
	js.Rewrite("$<.error")
	return err
}

// Mode prop
func (*IDBTransaction) Mode() (mode *idbtransactionmode.IDBTransactionMode) {
	js.Rewrite("$<.mode")
	return mode
}

// Onabort prop
func (*IDBTransaction) Onabort() (onabort func(Event)) {
	js.Rewrite("$<.onabort")
	return onabort
}

// Onabort prop
func (*IDBTransaction) SetOnabort(onabort func(Event)) {
	js.Rewrite("$<.onabort = $1", onabort)
}

// Oncomplete prop
func (*IDBTransaction) Oncomplete() (oncomplete func(Event)) {
	js.Rewrite("$<.oncomplete")
	return oncomplete
}

// Oncomplete prop
func (*IDBTransaction) SetOncomplete(oncomplete func(Event)) {
	js.Rewrite("$<.oncomplete = $1", oncomplete)
}

// Onerror prop
func (*IDBTransaction) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*IDBTransaction) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}
