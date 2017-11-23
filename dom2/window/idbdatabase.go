package window

import (
	"github.com/matthewmueller/golly/dom2/domstringlist"
	"github.com/matthewmueller/golly/dom2/idbobjectstoreparameters"
	"github.com/matthewmueller/golly/js"
)

// js:"IDBDatabase,omit"
type IDBDatabase struct {
	EventTarget
}

// Close
func (*IDBDatabase) Close() {
	js.Rewrite("$<.Close()")
}

// CreateObjectStore
func (*IDBDatabase) CreateObjectStore(name string, optionalParameters *idbobjectstoreparameters.IDBObjectStoreParameters) (i *IDBObjectStore) {
	js.Rewrite("$<.CreateObjectStore($1, $2)", name, optionalParameters)
	return i
}

// DeleteObjectStore
func (*IDBDatabase) DeleteObjectStore(name string) {
	js.Rewrite("$<.DeleteObjectStore($1)", name)
}

// Transaction
func (*IDBDatabase) Transaction(storeNames interface{}, mode *idbtransactionmode.IDBTransactionMode) (i *IDBTransaction) {
	js.Rewrite("$<.Transaction($1, $2)", storeNames, mode)
	return i
}

// Name
func (*IDBDatabase) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// ObjectStoreNames
func (*IDBDatabase) ObjectStoreNames() (objectStoreNames *domstringlist.DOMStringList) {
	js.Rewrite("$<.ObjectStoreNames")
	return objectStoreNames
}

// Onabort
func (*IDBDatabase) Onabort() (onabort func(Event)) {
	js.Rewrite("$<.Onabort")
	return onabort
}

// Onabort
func (*IDBDatabase) SetOnabort(onabort func(Event)) {
	js.Rewrite("$<.Onabort = $1", onabort)
}

// Onerror
func (*IDBDatabase) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*IDBDatabase) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Version
func (*IDBDatabase) Version() (version uint64) {
	js.Rewrite("$<.Version")
	return version
}
