package compositionevent

import (
	"github.com/matthewmueller/golly/dom/compositioneventinit"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(typearg string, eventinitdict *compositioneventinit.CompositionEventInit) *CompositionEvent {
	js.Rewrite("CompositionEvent")
	return &CompositionEvent{}
}

// CompositionEvent struct
// js:"CompositionEvent,omit"
type CompositionEvent struct {
	window.UIEvent
}

// InitCompositionEvent fn
func (*CompositionEvent) InitCompositionEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, dataArg string, locale string) {
	js.Rewrite("$<.initCompositionEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, viewArg, dataArg, locale)
}

// Data prop
func (*CompositionEvent) Data() (data string) {
	js.Rewrite("$<.data")
	return data
}

// Locale prop
func (*CompositionEvent) Locale() (locale string) {
	js.Rewrite("$<.locale")
	return locale
}
