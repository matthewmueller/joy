package beforeunloadevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type BeforeUnloadEvent struct {
	*event.Event
}

func (*BeforeUnloadEvent) GetReturnValue() (returnValue string) {
	js.Rewrite("$<.returnValue")
	return returnValue
}

func (*BeforeUnloadEvent) SetReturnValue(returnValue string) {
	js.Rewrite("$<.returnValue = $1", returnValue)
}
