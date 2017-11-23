package textevent

import "github.com/matthewmueller/golly/js"

// js:"TextEvent,omit"
type TextEvent struct {
	window.UIEvent
}

// InitTextEvent
func (*TextEvent) InitTextEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, dataArg string, inputMethod uint, locale string) {
	js.Rewrite("$<.InitTextEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, viewArg, dataArg, inputMethod, locale)
}

// Data
func (*TextEvent) Data() (data string) {
	js.Rewrite("$<.Data")
	return data
}

// InputMethod
func (*TextEvent) InputMethod() (inputMethod uint) {
	js.Rewrite("$<.InputMethod")
	return inputMethod
}

// Locale
func (*TextEvent) Locale() (locale string) {
	js.Rewrite("$<.Locale")
	return locale
}
