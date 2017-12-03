package audioprocessingevent

import (
	"github.com/matthewmueller/joy/dom/audiobuffer"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.Event = (*AudioProcessingEvent)(nil)

// AudioProcessingEvent struct
// js:"AudioProcessingEvent,omit"
type AudioProcessingEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*AudioProcessingEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*AudioProcessingEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*AudioProcessingEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*AudioProcessingEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// InputBuffer prop
// js:"inputBuffer"
func (*AudioProcessingEvent) InputBuffer() (inputBuffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$_.inputBuffer")
	return inputBuffer
}

// OutputBuffer prop
// js:"outputBuffer"
func (*AudioProcessingEvent) OutputBuffer() (outputBuffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$_.outputBuffer")
	return outputBuffer
}

// PlaybackTime prop
// js:"playbackTime"
func (*AudioProcessingEvent) PlaybackTime() (playbackTime float32) {
	js.Rewrite("$_.playbackTime")
	return playbackTime
}

// Bubbles prop
// js:"bubbles"
func (*AudioProcessingEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*AudioProcessingEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*AudioProcessingEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*AudioProcessingEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*AudioProcessingEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*AudioProcessingEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*AudioProcessingEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*AudioProcessingEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*AudioProcessingEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*AudioProcessingEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*AudioProcessingEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*AudioProcessingEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*AudioProcessingEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*AudioProcessingEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
