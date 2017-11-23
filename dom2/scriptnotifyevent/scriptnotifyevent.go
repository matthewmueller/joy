package scriptnotifyevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// ScriptNotifyEvent struct
// js:"ScriptNotifyEvent,omit"
type ScriptNotifyEvent struct {
	window.Event
}

// CallingURI
func (*ScriptNotifyEvent) CallingURI() (callingUri string) {
	js.Rewrite("$<.CallingURI")
	return callingUri
}

// Value
func (*ScriptNotifyEvent) Value() (value string) {
	js.Rewrite("$<.Value")
	return value
}
