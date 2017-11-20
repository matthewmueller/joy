package textevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type TextEvent struct {
	*uievent.UIEvent
}

func (*TextEvent) InitTextEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, dataArg string, inputMethod uint, locale string) {
	js.Rewrite("$<.initTextEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, viewArg, dataArg, inputMethod, locale)
}

func (*TextEvent) GetData() (data string) {
	js.Rewrite("$<.data")
	return data
}

func (*TextEvent) GetInputMethod() (inputMethod uint) {
	js.Rewrite("$<.inputMethod")
	return inputMethod
}

func (*TextEvent) GetLocale() (locale string) {
	js.Rewrite("$<.locale")
	return locale
}
