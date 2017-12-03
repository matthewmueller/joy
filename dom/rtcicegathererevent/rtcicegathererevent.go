package rtcicegathererevent

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*RTCIceGathererEvent)(nil)

// RTCIceGathererEvent struct
// js:"RTCIceGathererEvent,omit"
type RTCIceGathererEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*RTCIceGathererEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*RTCIceGathererEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*RTCIceGathererEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*RTCIceGathererEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// Candidate prop
// js:"candidate"
func (*RTCIceGathererEvent) Candidate() (candidate interface{}) {
	macro.Rewrite("$_.candidate")
	return candidate
}

// Bubbles prop
// js:"bubbles"
func (*RTCIceGathererEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*RTCIceGathererEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*RTCIceGathererEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*RTCIceGathererEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*RTCIceGathererEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*RTCIceGathererEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*RTCIceGathererEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*RTCIceGathererEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*RTCIceGathererEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*RTCIceGathererEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*RTCIceGathererEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*RTCIceGathererEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*RTCIceGathererEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*RTCIceGathererEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
