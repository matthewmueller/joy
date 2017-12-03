package window

import (
	"github.com/matthewmueller/joy/dom/progresseventinit"
	"github.com/matthewmueller/joy/macro"
)

var _ Event = (*ProgressEvent)(nil)

// NewProgressEvent fn
func NewProgressEvent(typearg string, eventinitdict *progresseventinit.ProgressEventInit) *ProgressEvent {
	macro.Rewrite("ProgressEvent")
	return &ProgressEvent{}
}

// ProgressEvent struct
// js:"ProgressEvent,omit"
type ProgressEvent struct {
}

// InitProgressEvent fn
// js:"initProgressEvent"
func (*ProgressEvent) InitProgressEvent(typeArg string, canBubbleArg bool, cancelableArg bool, lengthComputableArg bool, loadedArg uint64, totalArg uint64) {
	macro.Rewrite("$_.initProgressEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, lengthComputableArg, loadedArg, totalArg)
}

// InitEvent fn
// js:"initEvent"
func (*ProgressEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*ProgressEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*ProgressEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*ProgressEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// LengthComputable prop
// js:"lengthComputable"
func (*ProgressEvent) LengthComputable() (lengthComputable bool) {
	macro.Rewrite("$_.lengthComputable")
	return lengthComputable
}

// Loaded prop
// js:"loaded"
func (*ProgressEvent) Loaded() (loaded uint64) {
	macro.Rewrite("$_.loaded")
	return loaded
}

// Total prop
// js:"total"
func (*ProgressEvent) Total() (total uint64) {
	macro.Rewrite("$_.total")
	return total
}

// Bubbles prop
// js:"bubbles"
func (*ProgressEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*ProgressEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*ProgressEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*ProgressEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*ProgressEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*ProgressEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*ProgressEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*ProgressEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*ProgressEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*ProgressEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*ProgressEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*ProgressEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*ProgressEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*ProgressEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
