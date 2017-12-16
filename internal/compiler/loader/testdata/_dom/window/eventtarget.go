package window

// EventTarget
type EventTarget interface {
	DispatchEvent() Event
}
