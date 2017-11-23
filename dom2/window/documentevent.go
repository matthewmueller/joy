package window

// js:"DocumentEvent,omit"
type DocumentEvent interface {

	// CreateEvent
	CreateEvent(eventInterface string) (e Event)
}
