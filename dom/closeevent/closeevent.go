package closeevent

import (
	"github.com/matthewmueller/golly/dom/closeeventinit"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.Event = (*CloseEvent)(nil)

// New fn
func New(typearg string, eventinitdict *closeeventinit.CloseEventInit) *CloseEvent {
	js.Rewrite("CloseEvent")
	return &CloseEvent{}
}

// CloseEvent struct
// js:"CloseEvent,omit"
type CloseEvent struct {
}

// InitCloseEvent fn
// js:"initCloseEvent"
func (*CloseEvent) InitCloseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, wasCleanArg bool, codeArg uint8, reasonArg string) {
	js.Rewrite("$_.initCloseEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, wasCleanArg, codeArg, reasonArg)
}

// InitEvent fn
// js:"initEvent"
func (*CloseEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*CloseEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*CloseEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*CloseEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// Code prop
// js:"code"
func (*CloseEvent) Code() (code uint8) {
	js.Rewrite("$_.code")
	return code
}

// Reason prop
// js:"reason"
func (*CloseEvent) Reason() (reason string) {
	js.Rewrite("$_.reason")
	return reason
}

// WasClean prop
// js:"wasClean"
func (*CloseEvent) WasClean() (wasClean bool) {
	js.Rewrite("$_.wasClean")
	return wasClean
}

// Bubbles prop
// js:"bubbles"
func (*CloseEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*CloseEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*CloseEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*CloseEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*CloseEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*CloseEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*CloseEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*CloseEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*CloseEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*CloseEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*CloseEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*CloseEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*CloseEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*CloseEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
