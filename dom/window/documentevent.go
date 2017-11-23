package window

import "github.com/matthewmueller/golly/js"

// DocumentEvent struct
// js:"DocumentEvent,omit"
type DocumentEvent struct {
}

// CreateEvent fn
func (*DocumentEvent) CreateEvent(eventInterface string) (e Event) {
	js.Rewrite("$<.createEvent($1)", eventInterface)
	return e
}
