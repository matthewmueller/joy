package window

// EventTarget interface
type EventTarget interface {
	AddEventListener(kind string, listener func(evt Event), useCapture bool)
	RemoveEventListener(kind string, listener func(evt Event), useCapture bool)
	DispatchEvent(evt Event)
}
