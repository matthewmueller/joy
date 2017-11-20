package compositionevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type CompositionEvent struct {
	*uievent.UIEvent
}

func (*CompositionEvent) InitCompositionEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, dataArg string, locale string) {
	js.Rewrite("$<.initCompositionEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, viewArg, dataArg, locale)
}

func (*CompositionEvent) GetData() (data string) {
	js.Rewrite("$<.data")
	return data
}

func (*CompositionEvent) GetLocale() (locale string) {
	js.Rewrite("$<.locale")
	return locale
}
