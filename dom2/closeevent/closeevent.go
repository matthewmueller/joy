package closeevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"CloseEvent,omit"
type CloseEvent struct {
	window.Event
}

// InitCloseEvent
func (*CloseEvent) InitCloseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, wasCleanArg bool, codeArg uint8, reasonArg string) {
	js.Rewrite("$<.InitCloseEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, wasCleanArg, codeArg, reasonArg)
}

// Code
func (*CloseEvent) Code() (code uint8) {
	js.Rewrite("$<.Code")
	return code
}

// Reason
func (*CloseEvent) Reason() (reason string) {
	js.Rewrite("$<.Reason")
	return reason
}

// WasClean
func (*CloseEvent) WasClean() (wasClean bool) {
	js.Rewrite("$<.WasClean")
	return wasClean
}
