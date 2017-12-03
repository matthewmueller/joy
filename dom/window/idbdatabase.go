package window

import (
	"github.com/matthewmueller/joy/dom/domstringlist"
	"github.com/matthewmueller/joy/dom/idbobjectstoreparameters"
	"github.com/matthewmueller/joy/dom/idbtransactionmode"
	"github.com/matthewmueller/joy/macro"
)

var _ EventTarget = (*IDBDatabase)(nil)

// IDBDatabase struct
// js:"IDBDatabase,omit"
type IDBDatabase struct {
}

// Close fn
// js:"close"
func (*IDBDatabase) Close() {
	macro.Rewrite("$_.close()")
}

// CreateObjectStore fn
// js:"createObjectStore"
func (*IDBDatabase) CreateObjectStore(name string, optionalParameters *idbobjectstoreparameters.IDBObjectStoreParameters) (i *IDBObjectStore) {
	macro.Rewrite("$_.createObjectStore($1, $2)", name, optionalParameters)
	return i
}

// DeleteObjectStore fn
// js:"deleteObjectStore"
func (*IDBDatabase) DeleteObjectStore(name string) {
	macro.Rewrite("$_.deleteObjectStore($1)", name)
}

// Transaction fn
// js:"transaction"
func (*IDBDatabase) Transaction(storeNames interface{}, mode *idbtransactionmode.IDBTransactionMode) (i *IDBTransaction) {
	macro.Rewrite("$_.transaction($1, $2)", storeNames, mode)
	return i
}

// AddEventListener fn
// js:"addEventListener"
func (*IDBDatabase) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*IDBDatabase) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*IDBDatabase) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Name prop
// js:"name"
func (*IDBDatabase) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// ObjectStoreNames prop
// js:"objectStoreNames"
func (*IDBDatabase) ObjectStoreNames() (objectStoreNames *domstringlist.DOMStringList) {
	macro.Rewrite("$_.objectStoreNames")
	return objectStoreNames
}

// Onabort prop
// js:"onabort"
func (*IDBDatabase) Onabort() (onabort func(Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*IDBDatabase) SetOnabort(onabort func(Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

// Onerror prop
// js:"onerror"
func (*IDBDatabase) Onerror() (onerror func(Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*IDBDatabase) SetOnerror(onerror func(Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Version prop
// js:"version"
func (*IDBDatabase) Version() (version uint64) {
	macro.Rewrite("$_.version")
	return version
}
