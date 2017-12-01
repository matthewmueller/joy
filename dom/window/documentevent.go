package window

// DocumentEvent interface
// js:"DocumentEvent"
type DocumentEvent interface {

	// CreateEvent
	// js:"createEvent"
	// jsrewrite:"$_.createEvent($1)"
	CreateEvent(eventInterface string) (e Event)
}
