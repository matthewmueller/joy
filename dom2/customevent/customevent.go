package customevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"CustomEvent,omit"
type CustomEvent struct {
	window.Event
}

// InitCustomEvent
func (*CustomEvent) InitCustomEvent(typeArg string, canBubbleArg bool, cancelableArg bool, detailArg interface{}) {
	js.Rewrite("$<.InitCustomEvent($1, $2, $3, $4)", typeArg, canBubbleArg, cancelableArg, detailArg)
}

// Detail
func (*CustomEvent) Detail() (detail interface{}) {
	js.Rewrite("$<.Detail")
	return detail
}
