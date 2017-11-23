package closeevent

import (
	"github.com/matthewmueller/golly/dom2/closeeventinit"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(typearg string, eventinitdict *closeeventinit.CloseEventInit) *CloseEvent {
	js.Rewrite("CloseEvent")
	return &CloseEvent{}
}

// CloseEvent struct
// js:"CloseEvent,omit"
type CloseEvent struct {
	window.Event
}

// InitCloseEvent fn
func (*CloseEvent) InitCloseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, wasCleanArg bool, codeArg uint8, reasonArg string) {
	js.Rewrite("$<.initCloseEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, wasCleanArg, codeArg, reasonArg)
}

// Code prop
func (*CloseEvent) Code() (code uint8) {
	js.Rewrite("$<.code")
	return code
}

// Reason prop
func (*CloseEvent) Reason() (reason string) {
	js.Rewrite("$<.reason")
	return reason
}

// WasClean prop
func (*CloseEvent) WasClean() (wasClean bool) {
	js.Rewrite("$<.wasClean")
	return wasClean
}
