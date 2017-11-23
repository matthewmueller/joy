package idb

// js:"EventTarget,omit"
type EventTarget interface {

	// AddEventListener
	AddEventListener(kind string, listener func(evt Event), useCapture bool)

	// DispatchEvent
	DispatchEvent(evt Event) (b bool)

	// RemoveEventListener
	RemoveEventListener(kind string, listener func(evt Event), useCapture bool)
}
