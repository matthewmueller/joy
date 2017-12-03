package longrunningscriptdetectedevent

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*LongRunningScriptDetectedEvent)(nil)

// LongRunningScriptDetectedEvent struct
// js:"LongRunningScriptDetectedEvent,omit"
type LongRunningScriptDetectedEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*LongRunningScriptDetectedEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*LongRunningScriptDetectedEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*LongRunningScriptDetectedEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*LongRunningScriptDetectedEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// ExecutionTime prop
// js:"executionTime"
func (*LongRunningScriptDetectedEvent) ExecutionTime() (executionTime int) {
	macro.Rewrite("$_.executionTime")
	return executionTime
}

// StopPageScriptExecution prop
// js:"stopPageScriptExecution"
func (*LongRunningScriptDetectedEvent) StopPageScriptExecution() (stopPageScriptExecution bool) {
	macro.Rewrite("$_.stopPageScriptExecution")
	return stopPageScriptExecution
}

// SetStopPageScriptExecution prop
// js:"stopPageScriptExecution"
func (*LongRunningScriptDetectedEvent) SetStopPageScriptExecution(stopPageScriptExecution bool) {
	macro.Rewrite("$_.stopPageScriptExecution = $1", stopPageScriptExecution)
}

// Bubbles prop
// js:"bubbles"
func (*LongRunningScriptDetectedEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*LongRunningScriptDetectedEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*LongRunningScriptDetectedEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*LongRunningScriptDetectedEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*LongRunningScriptDetectedEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*LongRunningScriptDetectedEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*LongRunningScriptDetectedEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*LongRunningScriptDetectedEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*LongRunningScriptDetectedEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*LongRunningScriptDetectedEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*LongRunningScriptDetectedEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*LongRunningScriptDetectedEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*LongRunningScriptDetectedEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*LongRunningScriptDetectedEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
