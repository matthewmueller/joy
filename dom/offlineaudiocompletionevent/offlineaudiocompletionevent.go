package offlineaudiocompletionevent

import (
	"github.com/matthewmueller/joy/dom/audiobuffer"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.Event = (*OfflineAudioCompletionEvent)(nil)

// OfflineAudioCompletionEvent struct
// js:"OfflineAudioCompletionEvent,omit"
type OfflineAudioCompletionEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*OfflineAudioCompletionEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*OfflineAudioCompletionEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*OfflineAudioCompletionEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*OfflineAudioCompletionEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// RenderedBuffer prop
// js:"renderedBuffer"
func (*OfflineAudioCompletionEvent) RenderedBuffer() (renderedBuffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$_.renderedBuffer")
	return renderedBuffer
}

// Bubbles prop
// js:"bubbles"
func (*OfflineAudioCompletionEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*OfflineAudioCompletionEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*OfflineAudioCompletionEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*OfflineAudioCompletionEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*OfflineAudioCompletionEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*OfflineAudioCompletionEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*OfflineAudioCompletionEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*OfflineAudioCompletionEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*OfflineAudioCompletionEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*OfflineAudioCompletionEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*OfflineAudioCompletionEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*OfflineAudioCompletionEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*OfflineAudioCompletionEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*OfflineAudioCompletionEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
