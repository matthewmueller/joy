package closeevent

import (
	"github.com/matthewmueller/joy/dom/closeeventinit"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*CloseEvent)(nil)

// New fn
func New(typearg string, eventinitdict *closeeventinit.CloseEventInit) *CloseEvent {
	macro.Rewrite("CloseEvent")
	return &CloseEvent{}
}

// CloseEvent struct
// js:"CloseEvent,omit"
type CloseEvent struct {
}

// InitCloseEvent fn
// js:"initCloseEvent"
func (*CloseEvent) InitCloseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, wasCleanArg bool, codeArg uint8, reasonArg string) {
	macro.Rewrite("$_.initCloseEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, wasCleanArg, codeArg, reasonArg)
}

// InitEvent fn
// js:"initEvent"
func (*CloseEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*CloseEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*CloseEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*CloseEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// Code prop
// js:"code"
func (*CloseEvent) Code() (code uint8) {
	macro.Rewrite("$_.code")
	return code
}

// Reason prop
// js:"reason"
func (*CloseEvent) Reason() (reason string) {
	macro.Rewrite("$_.reason")
	return reason
}

// WasClean prop
// js:"wasClean"
func (*CloseEvent) WasClean() (wasClean bool) {
	macro.Rewrite("$_.wasClean")
	return wasClean
}

// Bubbles prop
// js:"bubbles"
func (*CloseEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*CloseEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*CloseEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*CloseEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*CloseEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*CloseEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*CloseEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*CloseEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*CloseEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*CloseEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*CloseEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*CloseEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*CloseEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*CloseEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
