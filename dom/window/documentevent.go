package window

// DocumentEvent interface
// js:"DocumentEvent"
type DocumentEvent interface {

	// CreateEvent
	// js:"createEvent"
	CreateEvent(eventInterface string) (e Event)
}
