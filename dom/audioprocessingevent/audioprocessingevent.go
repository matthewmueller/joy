package audioprocessingevent

import (
	"github.com/matthewmueller/joy/dom/audiobuffer"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*AudioProcessingEvent)(nil)

// AudioProcessingEvent struct
// js:"AudioProcessingEvent,omit"
type AudioProcessingEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*AudioProcessingEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*AudioProcessingEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*AudioProcessingEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*AudioProcessingEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// InputBuffer prop
// js:"inputBuffer"
func (*AudioProcessingEvent) InputBuffer() (inputBuffer *audiobuffer.AudioBuffer) {
	macro.Rewrite("$_.inputBuffer")
	return inputBuffer
}

// OutputBuffer prop
// js:"outputBuffer"
func (*AudioProcessingEvent) OutputBuffer() (outputBuffer *audiobuffer.AudioBuffer) {
	macro.Rewrite("$_.outputBuffer")
	return outputBuffer
}

// PlaybackTime prop
// js:"playbackTime"
func (*AudioProcessingEvent) PlaybackTime() (playbackTime float32) {
	macro.Rewrite("$_.playbackTime")
	return playbackTime
}

// Bubbles prop
// js:"bubbles"
func (*AudioProcessingEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*AudioProcessingEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*AudioProcessingEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*AudioProcessingEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*AudioProcessingEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*AudioProcessingEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*AudioProcessingEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*AudioProcessingEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*AudioProcessingEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*AudioProcessingEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*AudioProcessingEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*AudioProcessingEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*AudioProcessingEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*AudioProcessingEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
