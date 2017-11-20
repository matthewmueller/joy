package idbtransaction

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domerror"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbdatabase"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbobjectstore"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbtransactionmode"
	"github.com/matthewmueller/golly/js"
)

type IDBTransaction struct {
	*eventtarget.EventTarget
}

func (*IDBTransaction) Abort() {
	js.Rewrite("$<.abort()")
}

func (*IDBTransaction) ObjectStore(name string) (i *idbobjectstore.IDBObjectStore) {
	js.Rewrite("$<.objectStore($1)", name)
	return i
}

func (*IDBTransaction) GetDb() (db *idbdatabase.IDBDatabase) {
	js.Rewrite("$<.db")
	return db
}

func (*IDBTransaction) GetError() (err *domerror.DOMError) {
	js.Rewrite("$<.error")
	return err
}

func (*IDBTransaction) GetMode() (mode *idbtransactionmode.IDBTransactionMode) {
	js.Rewrite("$<.mode")
	return mode
}

func (*IDBTransaction) GetOnabort() (abort *event.Event) {
	js.Rewrite("$<.onabort")
	return abort
}

func (*IDBTransaction) SetOnabort(abort *event.Event) {
	js.Rewrite("$<.onabort = $1", abort)
}

func (*IDBTransaction) GetOncomplete() (complete *event.Event) {
	js.Rewrite("$<.oncomplete")
	return complete
}

func (*IDBTransaction) SetOncomplete(complete *event.Event) {
	js.Rewrite("$<.oncomplete = $1", complete)
}

func (*IDBTransaction) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*IDBTransaction) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}
