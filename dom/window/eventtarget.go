package window

// EventTarget interface
// js:"EventTarget"
type EventTarget interface {

	// AddEventListener
	// js:"addEventListener"
	AddEventListener(kind string, listener func(evt Event), useCapture bool)

	// DispatchEvent
	// js:"dispatchEvent"
	DispatchEvent(evt Event) (b bool)

	// RemoveEventListener
	// js:"removeEventListener"
	RemoveEventListener(kind string, listener func(evt Event), useCapture bool)
}
