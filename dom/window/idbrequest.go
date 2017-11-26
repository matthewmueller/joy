package window

import (
	"github.com/matthewmueller/golly/dom/domerror"
	"github.com/matthewmueller/golly/dom/idbrequestreadystate"
	"github.com/matthewmueller/golly/js"
)

// js:"IDBRequest,omit"
type IDBRequest interface {
	EventTarget

	// Error prop
	// js:"error,rewrite=err"
	Error() (err *domerror.DOMError)

	// Onerror prop
	// js:"onerror,rewrite=onerror"
	Onerror() (onerror func(Event))

	// Onerror prop
	// js:"setonerror,rewrite=setonerror"
	SetOnerror(onerror func(Event))

	// Onsuccess prop
	// js:"onsuccess,rewrite=onsuccess"
	Onsuccess() (onsuccess func(Event))

	// Onsuccess prop
	// js:"setonsuccess,rewrite=setonsuccess"
	SetOnsuccess(onsuccess func(Event))

	// ReadyState prop
	// js:"readyState,rewrite=readystate"
	ReadyState() (readyState *idbrequestreadystate.IDBRequestReadyState)

	// Result prop
	// js:"result,rewrite=result"
	Result() (result interface{})

	// Source prop
	// js:"source,rewrite=source"
	Source() (source interface{})

	// Transaction prop
	// js:"transaction,rewrite=transaction"
	Transaction() (transaction *IDBTransaction)
}

// err prop
func err() (err *domerror.DOMError) {
	js.Rewrite("$<.error")
	return err
}

// // onerror prop
// func onerror() (onerror func(Event)) {
// 	js.Rewrite("$<.onerror")
// 	return onerror
// }

// // setonerror prop
// func setonerror(onerror func(Event)) {
// 	js.Rewrite("$<.onerror = onerror")
// }

// onsuccess prop
func onsuccess() (onsuccess func(Event)) {
	js.Rewrite("$<.onsuccess")
	return onsuccess
}

// setonsuccess prop
func setonsuccess(onsuccess func(Event)) {
	js.Rewrite("$<.onsuccess = onsuccess")
}

// // readystate prop
// func readystate() (readyState *idbrequestreadystate.IDBRequestReadyState) {
// 	js.Rewrite("$<.readyState")
// 	return readyState
// }

// result prop
func result() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}

// source prop
func source() (source interface{}) {
	js.Rewrite("$<.source")
	return source
}

// transaction prop
func transaction() (transaction *IDBTransaction) {
	js.Rewrite("$<.transaction")
	return transaction
}
