package window

import "github.com/matthewmueller/golly/js"

// js:"EventTarget,omit"
type EventTarget interface {

	// AddEventListener
	// js:"addEventListener,rewrite=addeventlistener"
	AddEventListener(kind string, listener func(evt Event), useCapture bool)

	// DispatchEvent
	// js:"dispatchEvent,rewrite=dispatchevent"
	DispatchEvent(evt Event) (b bool)

	// RemoveEventListener
	// js:"removeEventListener,rewrite=removeeventlistener"
	RemoveEventListener(kind string, listener func(evt Event), useCapture bool)
}

// addeventlistener fn
func addeventlistener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$<.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// dispatchevent fn
func dispatchevent(evt Event) (b bool) {
	js.Rewrite("$<.dispatchEvent($1)", evt)
	return b
}

// removeeventlistener fn
func removeeventlistener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$<.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}
