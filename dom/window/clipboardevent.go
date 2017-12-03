package window

import (
	"github.com/matthewmueller/joy/dom/datatransfer"
	"github.com/matthewmueller/joy/macro"
)

var _ Event = (*ClipboardEvent)(nil)

// ClipboardEvent struct
// js:"ClipboardEvent,omit"
type ClipboardEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*ClipboardEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*ClipboardEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*ClipboardEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*ClipboardEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// ClipboardData prop
// js:"clipboardData"
func (*ClipboardEvent) ClipboardData() (clipboardData *datatransfer.DataTransfer) {
	macro.Rewrite("$_.clipboardData")
	return clipboardData
}

// Bubbles prop
// js:"bubbles"
func (*ClipboardEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*ClipboardEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*ClipboardEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*ClipboardEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*ClipboardEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*ClipboardEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*ClipboardEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*ClipboardEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*ClipboardEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*ClipboardEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*ClipboardEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*ClipboardEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*ClipboardEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*ClipboardEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
