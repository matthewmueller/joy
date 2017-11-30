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
	Error() (err *domerror.DOMError)

	// Onerror prop
	// js:"onerror"
	Onerror() (onerror func(Event))

	// SetOnerror prop
	// js:"onerror"
	SetOnerror(onerror func(Event))

	// Onsuccess prop
	// js:"onsuccess"
	Onsuccess() (onsuccess func(Event))

	// SetOnsuccess prop
	// js:"onsuccess"
	SetOnsuccess(onsuccess func(Event))

	// ReadyState prop
	// js:"readyState"
	ReadyState() (readyState *idbrequestreadystate.IDBRequestReadyState)

	// Result prop
	// js:"result"
	Result() (result interface{})

	// Source prop
	// js:"source"
	Source() (source interface{})

	// Transaction prop
	// js:"transaction"
	Transaction() (transaction *IDBTransaction)
}
