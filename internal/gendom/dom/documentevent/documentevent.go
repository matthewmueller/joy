package documentevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type DocumentEvent struct {
}

func (*DocumentEvent) CreateEvent(eventInterface string) (e *event.Event) {
	js.Rewrite("$<.createEvent($1)", eventInterface)
	return e
}
