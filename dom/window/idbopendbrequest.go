package window

import (
	"github.com/matthewmueller/joy/dom/domerror"
	"github.com/matthewmueller/joy/dom/idbrequestreadystate"
	"github.com/matthewmueller/joy/macro"
)

var _ IDBRequest = (*IDBOpenDBRequest)(nil)
var _ EventTarget = (*IDBOpenDBRequest)(nil)

// IDBOpenDBRequest struct
// js:"IDBOpenDBRequest,omit"
type IDBOpenDBRequest struct {
}

// AddEventListener fn
// js:"addEventListener"
func (*IDBOpenDBRequest) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*IDBOpenDBRequest) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*IDBOpenDBRequest) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onblocked prop
// js:"onblocked"
func (*IDBOpenDBRequest) Onblocked() (onblocked func(Event)) {
	macro.Rewrite("$_.onblocked")
	return onblocked
}

// SetOnblocked prop
// js:"onblocked"
func (*IDBOpenDBRequest) SetOnblocked(onblocked func(Event)) {
	macro.Rewrite("$_.onblocked = $1", onblocked)
}

// Onupgradeneeded prop
// js:"onupgradeneeded"
func (*IDBOpenDBRequest) Onupgradeneeded() (onupgradeneeded func(*IDBVersionChangeEvent)) {
	macro.Rewrite("$_.onupgradeneeded")
	return onupgradeneeded
}

// SetOnupgradeneeded prop
// js:"onupgradeneeded"
func (*IDBOpenDBRequest) SetOnupgradeneeded(onupgradeneeded func(*IDBVersionChangeEvent)) {
	macro.Rewrite("$_.onupgradeneeded = $1", onupgradeneeded)
}

// Error prop
// js:"error"
func (*IDBOpenDBRequest) Error() (err *domerror.DOMError) {
	macro.Rewrite("$_.error")
	return err
}

// Onerror prop
// js:"onerror"
func (*IDBOpenDBRequest) Onerror() (onerror func(Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*IDBOpenDBRequest) SetOnerror(onerror func(Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onsuccess prop
// js:"onsuccess"
func (*IDBOpenDBRequest) Onsuccess() (onsuccess func(Event)) {
	macro.Rewrite("$_.onsuccess")
	return onsuccess
}

// SetOnsuccess prop
// js:"onsuccess"
func (*IDBOpenDBRequest) SetOnsuccess(onsuccess func(Event)) {
	macro.Rewrite("$_.onsuccess = $1", onsuccess)
}

// ReadyState prop
// js:"readyState"
func (*IDBOpenDBRequest) ReadyState() (readyState *idbrequestreadystate.IDBRequestReadyState) {
	macro.Rewrite("$_.readyState")
	return readyState
}

// Result prop
// js:"result"
func (*IDBOpenDBRequest) Result() (result interface{}) {
	macro.Rewrite("$_.result")
	return result
}

// Source prop
// js:"source"
func (*IDBOpenDBRequest) Source() (source interface{}) {
	macro.Rewrite("$_.source")
	return source
}

// Transaction prop
// js:"transaction"
func (*IDBOpenDBRequest) Transaction() (transaction *IDBTransaction) {
	macro.Rewrite("$_.transaction")
	return transaction
}
