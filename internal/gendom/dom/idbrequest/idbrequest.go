package idbrequest

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domerror"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbrequestreadystate"
	"github.com/matthewmueller/golly/internal/gendom/dom/idbtransaction"
	"github.com/matthewmueller/golly/js"
)

type IDBRequest struct {
	*eventtarget.EventTarget
}

func (*IDBRequest) GetError() (err *domerror.DOMError) {
	js.Rewrite("$<.error")
	return err
}

func (*IDBRequest) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*IDBRequest) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*IDBRequest) GetOnsuccess() (success *event.Event) {
	js.Rewrite("$<.onsuccess")
	return success
}

func (*IDBRequest) SetOnsuccess(success *event.Event) {
	js.Rewrite("$<.onsuccess = $1", success)
}

func (*IDBRequest) GetReadyState() (readyState *idbrequestreadystate.IDBRequestReadyState) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*IDBRequest) GetResult() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}

func (*IDBRequest) GetSource() (source interface{}) {
	js.Rewrite("$<.source")
	return source
}

func (*IDBRequest) GetTransaction() (transaction *idbtransaction.IDBTransaction) {
	js.Rewrite("$<.transaction")
	return transaction
}
