package window

import (
	"github.com/matthewmueller/golly/dom/domerror"
	"github.com/matthewmueller/golly/dom/idbrequestreadystate"
)

// IDBRequest interface
// js:"IDBRequest"
type IDBRequest interface {
	EventTarget

	// Error prop
	// js:"error"
	// jsrewrite:"$_.error"
	Error() (err *domerror.DOMError)

	// Onerror prop
	// js:"onerror"
	// jsrewrite:"$_.onerror"
	Onerror() (onerror func(Event))

	// SetOnerror prop
	// js:"onerror"
	// jsrewrite:"$_.onerror = $1"
	SetOnerror(onerror func(Event))

	// Onsuccess prop
	// js:"onsuccess"
	// jsrewrite:"$_.onsuccess"
	Onsuccess() (onsuccess func(Event))

	// SetOnsuccess prop
	// js:"onsuccess"
	// jsrewrite:"$_.onsuccess = $1"
	SetOnsuccess(onsuccess func(Event))

	// ReadyState prop
	// js:"readyState"
	// jsrewrite:"$_.readyState"
	ReadyState() (readyState *idbrequestreadystate.IDBRequestReadyState)

	// Result prop
	// js:"result"
	// jsrewrite:"$_.result"
	Result() (result interface{})

	// Source prop
	// js:"source"
	// jsrewrite:"$_.source"
	Source() (source interface{})

	// Transaction prop
	// js:"transaction"
	// jsrewrite:"$_.transaction"
	Transaction() (transaction *IDBTransaction)
}
