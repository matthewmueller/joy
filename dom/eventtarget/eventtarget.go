package eventtarget

import "github.com/matthewmueller/golly/dom/event"

// EventTarget interface
// js:"EventTarget"
type EventTarget interface {

	// AddEventListener
	// js:"addEventListener"
	AddEventListener(kind string, listener func(evt event.Event), useCapture bool)

	// DispatchEvent
	// js:"dispatchEvent"
	DispatchEvent(evt event.Event) (b bool)
}
