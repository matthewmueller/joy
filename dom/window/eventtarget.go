package window

// EventTarget interface
// js:"EventTarget"
type EventTarget interface {

	// AddEventListener
	// js:"addEventListener"
	// jsrewrite:"$_.addEventListener($1, $2, $3)"
	AddEventListener(kind string, listener func(evt Event), useCapture bool)

	// DispatchEvent
	// js:"dispatchEvent"
	// jsrewrite:"$_.dispatchEvent($1)"
	DispatchEvent(evt Event) (b bool)

	// RemoveEventListener
	// js:"removeEventListener"
	// jsrewrite:"$_.removeEventListener($1, $2, $3)"
	RemoveEventListener(kind string, listener func(evt Event), useCapture bool)
}
