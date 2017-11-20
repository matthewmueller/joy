package eventtarget

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type EventTarget struct {
}

func (*EventTarget) AddEventListener(kind string, listener func(evt *event.Event), useCapture bool) {
	js.Rewrite("$<.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*EventTarget) DispatchEvent(evt *event.Event) (b bool) {
	js.Rewrite("$<.dispatchEvent($1)", evt)
	return b
}

func (*EventTarget) RemoveEventListener(kind string, listener func(evt *event.Event), useCapture bool) {
	js.Rewrite("$<.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}
