package navigationcompletedevent

import (
	"github.com/matthewmueller/joy/dom/navigationevent"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ navigationevent.NavigationEvent = (*NavigationCompletedEvent)(nil)
var _ window.Event = (*NavigationCompletedEvent)(nil)

// NavigationCompletedEvent struct
// js:"NavigationCompletedEvent,omit"
type NavigationCompletedEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*NavigationCompletedEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*NavigationCompletedEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*NavigationCompletedEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*NavigationCompletedEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// IsSuccess prop
// js:"isSuccess"
func (*NavigationCompletedEvent) IsSuccess() (isSuccess bool) {
	macro.Rewrite("$_.isSuccess")
	return isSuccess
}

// WebErrorStatus prop
// js:"webErrorStatus"
func (*NavigationCompletedEvent) WebErrorStatus() (webErrorStatus uint64) {
	macro.Rewrite("$_.webErrorStatus")
	return webErrorStatus
}

// URI prop
// js:"uri"
func (*NavigationCompletedEvent) URI() (uri string) {
	macro.Rewrite("$_.uri")
	return uri
}

// Bubbles prop
// js:"bubbles"
func (*NavigationCompletedEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*NavigationCompletedEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*NavigationCompletedEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*NavigationCompletedEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*NavigationCompletedEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*NavigationCompletedEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*NavigationCompletedEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*NavigationCompletedEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*NavigationCompletedEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*NavigationCompletedEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*NavigationCompletedEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*NavigationCompletedEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*NavigationCompletedEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*NavigationCompletedEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
