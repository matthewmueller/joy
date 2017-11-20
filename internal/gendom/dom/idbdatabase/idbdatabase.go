package idbdatabase

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domstringlist"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbobjectstore"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbobjectstoreparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbtransaction"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbtransactionmode"
	"github.com/matthewmueller/golly/js"
)

type IDBDatabase struct {
	*eventtarget.EventTarget
}

func (*IDBDatabase) Close() {
	js.Rewrite("$<.close()")
}

func (*IDBDatabase) CreateObjectStore(name string, optionalParameters *idbobjectstoreparameters.IDBObjectStoreParameters) (i *idbobjectstore.IDBObjectStore) {
	js.Rewrite("$<.createObjectStore($1, $2)", name, optionalParameters)
	return i
}

func (*IDBDatabase) DeleteObjectStore(name string) {
	js.Rewrite("$<.deleteObjectStore($1)", name)
}

func (*IDBDatabase) Transaction(storeNames interface{}, mode *idbtransactionmode.IDBTransactionMode) (i *idbtransaction.IDBTransaction) {
	js.Rewrite("$<.transaction($1, $2)", storeNames, mode)
	return i
}

func (*IDBDatabase) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*IDBDatabase) GetObjectStoreNames() (objectStoreNames *domstringlist.DOMStringList) {
	js.Rewrite("$<.objectStoreNames")
	return objectStoreNames
}

func (*IDBDatabase) GetOnabort() (abort *event.Event) {
	js.Rewrite("$<.onabort")
	return abort
}

func (*IDBDatabase) SetOnabort(abort *event.Event) {
	js.Rewrite("$<.onabort = $1", abort)
}

func (*IDBDatabase) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*IDBDatabase) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*IDBDatabase) GetVersion() (version uint64) {
	js.Rewrite("$<.version")
	return version
}
