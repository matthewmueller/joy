package scriptnotifyevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type ScriptNotifyEvent struct {
	*event.Event
}

func (*ScriptNotifyEvent) GetCallingURI() (callingUri string) {
	js.Rewrite("$<.callingUri")
	return callingUri
}

func (*ScriptNotifyEvent) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}
