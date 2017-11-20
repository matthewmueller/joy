package devicelightevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type DeviceLightEvent struct {
	*event.Event
}

func (*DeviceLightEvent) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}
