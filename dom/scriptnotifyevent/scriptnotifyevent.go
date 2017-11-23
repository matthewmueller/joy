package scriptnotifyevent

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// ScriptNotifyEvent struct
// js:"ScriptNotifyEvent,omit"
type ScriptNotifyEvent struct {
	window.Event
}

// CallingURI prop
func (*ScriptNotifyEvent) CallingURI() (callingUri string) {
	js.Rewrite("$<.callingUri")
	return callingUri
}

// Value prop
func (*ScriptNotifyEvent) Value() (value string) {
	js.Rewrite("$<.value")
	return value
}
