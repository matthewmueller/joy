package textevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// TextEvent struct
// js:"TextEvent,omit"
type TextEvent struct {
	window.UIEvent
}

// InitTextEvent fn
func (*TextEvent) InitTextEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, dataArg string, inputMethod uint, locale string) {
	js.Rewrite("$<.initTextEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, viewArg, dataArg, inputMethod, locale)
}

// Data prop
func (*TextEvent) Data() (data string) {
	js.Rewrite("$<.data")
	return data
}

// InputMethod prop
func (*TextEvent) InputMethod() (inputMethod uint) {
	js.Rewrite("$<.inputMethod")
	return inputMethod
}

// Locale prop
func (*TextEvent) Locale() (locale string) {
	js.Rewrite("$<.locale")
	return locale
}
