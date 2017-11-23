package compositionevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"CompositionEvent,omit"
type CompositionEvent struct {
	window.UIEvent
}

// InitCompositionEvent
func (*CompositionEvent) InitCompositionEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, dataArg string, locale string) {
	js.Rewrite("$<.InitCompositionEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, viewArg, dataArg, locale)
}

// Data
func (*CompositionEvent) Data() (data string) {
	js.Rewrite("$<.Data")
	return data
}

// Locale
func (*CompositionEvent) Locale() (locale string) {
	js.Rewrite("$<.Locale")
	return locale
}
