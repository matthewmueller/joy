package rtcdtmftonechangeevent

import (
	"github.com/matthewmueller/joy/dom/rtcdtmftonechangeeventinit"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*RTCDTMFToneChangeEvent)(nil)

// New fn
func New(typearg string, eventinitdict *rtcdtmftonechangeeventinit.RTCDTMFToneChangeEventInit) *RTCDTMFToneChangeEvent {
	macro.Rewrite("RTCDTMFToneChangeEvent")
	return &RTCDTMFToneChangeEvent{}
}

// RTCDTMFToneChangeEvent struct
// js:"RTCDTMFToneChangeEvent,omit"
type RTCDTMFToneChangeEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*RTCDTMFToneChangeEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*RTCDTMFToneChangeEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*RTCDTMFToneChangeEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*RTCDTMFToneChangeEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// Tone prop
// js:"tone"
func (*RTCDTMFToneChangeEvent) Tone() (tone string) {
	macro.Rewrite("$_.tone")
	return tone
}

// Bubbles prop
// js:"bubbles"
func (*RTCDTMFToneChangeEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*RTCDTMFToneChangeEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*RTCDTMFToneChangeEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*RTCDTMFToneChangeEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*RTCDTMFToneChangeEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*RTCDTMFToneChangeEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*RTCDTMFToneChangeEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*RTCDTMFToneChangeEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*RTCDTMFToneChangeEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*RTCDTMFToneChangeEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*RTCDTMFToneChangeEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*RTCDTMFToneChangeEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*RTCDTMFToneChangeEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*RTCDTMFToneChangeEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
