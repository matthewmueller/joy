package window

import "github.com/matthewmueller/golly/js"

// BeforeUnloadEvent struct
// js:"BeforeUnloadEvent,omit"
type BeforeUnloadEvent struct {
	Event
}

// ReturnValue prop
func (*BeforeUnloadEvent) ReturnValue() (returnValue string) {
	js.Rewrite("$<.returnValue")
	return returnValue
}

// ReturnValue prop
func (*BeforeUnloadEvent) SetReturnValue(returnValue string) {
	js.Rewrite("$<.returnValue = $1", returnValue)
}
