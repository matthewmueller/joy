package idb

import "github.com/matthewmueller/golly/js"

// DocumentEvent struct
// js:"DocumentEvent,omit"
type DocumentEvent struct {
}

// CreateEvent
func (*DocumentEvent) CreateEvent(eventInterface string) (e Event) {
	js.Rewrite("$<.CreateEvent($1)", eventInterface)
	return e
}
