package longrunningscriptdetectedevent

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.Event = (*LongRunningScriptDetectedEvent)(nil)

// LongRunningScriptDetectedEvent struct
// js:"LongRunningScriptDetectedEvent,omit"
type LongRunningScriptDetectedEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*LongRunningScriptDetectedEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*LongRunningScriptDetectedEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*LongRunningScriptDetectedEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*LongRunningScriptDetectedEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// ExecutionTime prop
// js:"executionTime"
func (*LongRunningScriptDetectedEvent) ExecutionTime() (executionTime int) {
	js.Rewrite("$_.executionTime")
	return executionTime
}

// StopPageScriptExecution prop
// js:"stopPageScriptExecution"
func (*LongRunningScriptDetectedEvent) StopPageScriptExecution() (stopPageScriptExecution bool) {
	js.Rewrite("$_.stopPageScriptExecution")
	return stopPageScriptExecution
}

// SetStopPageScriptExecution prop
// js:"stopPageScriptExecution"
func (*LongRunningScriptDetectedEvent) SetStopPageScriptExecution(stopPageScriptExecution bool) {
	js.Rewrite("$_.stopPageScriptExecution = $1", stopPageScriptExecution)
}

// Bubbles prop
// js:"bubbles"
func (*LongRunningScriptDetectedEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*LongRunningScriptDetectedEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*LongRunningScriptDetectedEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*LongRunningScriptDetectedEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*LongRunningScriptDetectedEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*LongRunningScriptDetectedEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*LongRunningScriptDetectedEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*LongRunningScriptDetectedEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*LongRunningScriptDetectedEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*LongRunningScriptDetectedEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*LongRunningScriptDetectedEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*LongRunningScriptDetectedEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*LongRunningScriptDetectedEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*LongRunningScriptDetectedEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
