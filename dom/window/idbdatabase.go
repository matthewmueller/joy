package window

import (
	"github.com/matthewmueller/golly/dom2/domstringlist"
	"github.com/matthewmueller/golly/dom2/idbobjectstoreparameters"
	"github.com/matthewmueller/golly/dom2/idbtransactionmode"
	"github.com/matthewmueller/golly/js"
)

// IDBDatabase struct
// js:"IDBDatabase,omit"
type IDBDatabase struct {
	EventTarget
}

// Close fn
func (*IDBDatabase) Close() {
	js.Rewrite("$<.close()")
}

// CreateObjectStore fn
func (*IDBDatabase) CreateObjectStore(name string, optionalParameters *idbobjectstoreparameters.IDBObjectStoreParameters) (i *IDBObjectStore) {
	js.Rewrite("$<.createObjectStore($1, $2)", name, optionalParameters)
	return i
}

// DeleteObjectStore fn
func (*IDBDatabase) DeleteObjectStore(name string) {
	js.Rewrite("$<.deleteObjectStore($1)", name)
}

// Transaction fn
func (*IDBDatabase) Transaction(storeNames interface{}, mode *idbtransactionmode.IDBTransactionMode) (i *IDBTransaction) {
	js.Rewrite("$<.transaction($1, $2)", storeNames, mode)
	return i
}

// Name prop
func (*IDBDatabase) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// ObjectStoreNames prop
func (*IDBDatabase) ObjectStoreNames() (objectStoreNames *domstringlist.DOMStringList) {
	js.Rewrite("$<.objectStoreNames")
	return objectStoreNames
}

// Onabort prop
func (*IDBDatabase) Onabort() (onabort func(Event)) {
	js.Rewrite("$<.onabort")
	return onabort
}

// Onabort prop
func (*IDBDatabase) SetOnabort(onabort func(Event)) {
	js.Rewrite("$<.onabort = $1", onabort)
}

// Onerror prop
func (*IDBDatabase) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*IDBDatabase) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Version prop
func (*IDBDatabase) Version() (version uint64) {
	js.Rewrite("$<.version")
	return version
}
