package idb

import "github.com/matthewmueller/golly/js"

// BeforeUnloadEvent struct
// js:"BeforeUnloadEvent,omit"
type BeforeUnloadEvent struct {
	Event
}

// ReturnValue
func (*BeforeUnloadEvent) ReturnValue() (returnValue string) {
	js.Rewrite("$<.ReturnValue")
	return returnValue
}

// ReturnValue
func (*BeforeUnloadEvent) SetReturnValue(returnValue string) {
	js.Rewrite("$<.ReturnValue = $1", returnValue)
}
