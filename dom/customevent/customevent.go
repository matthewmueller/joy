package customevent

import (
	"github.com/matthewmueller/golly/dom/customeventinit"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(typearg string, eventinitdict *customeventinit.CustomEventInit) *CustomEvent {
	js.Rewrite("CustomEvent")
	return &CustomEvent{}
}

// CustomEvent struct
// js:"CustomEvent,omit"
type CustomEvent struct {
	window.Event
}

// InitCustomEvent fn
func (*CustomEvent) InitCustomEvent(typeArg string, canBubbleArg bool, cancelableArg bool, detailArg interface{}) {
	js.Rewrite("$<.initCustomEvent($1, $2, $3, $4)", typeArg, canBubbleArg, cancelableArg, detailArg)
}

// Detail prop
func (*CustomEvent) Detail() (detail interface{}) {
	js.Rewrite("$<.detail")
	return detail
}
