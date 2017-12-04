package speechsynthesisevent

import (
	"github.com/matthewmueller/joy/dom/speechsynthesiseventinit"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*SpeechSynthesisEvent)(nil)

// New fn
func New(kind string, eventinitdict *speechsynthesiseventinit.SpeechSynthesisEventInit) *SpeechSynthesisEvent {
	macro.Rewrite("new SpeechSynthesisEvent($1, $2)", kind, eventinitdict)
	return &SpeechSynthesisEvent{}
}

// SpeechSynthesisEvent struct
// js:"SpeechSynthesisEvent,omit"
type SpeechSynthesisEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*SpeechSynthesisEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*SpeechSynthesisEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*SpeechSynthesisEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*SpeechSynthesisEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// CharIndex prop
// js:"charIndex"
func (*SpeechSynthesisEvent) CharIndex() (charIndex uint) {
	macro.Rewrite("$_.charIndex")
	return charIndex
}

// ElapsedTime prop
// js:"elapsedTime"
func (*SpeechSynthesisEvent) ElapsedTime() (elapsedTime float32) {
	macro.Rewrite("$_.elapsedTime")
	return elapsedTime
}

// Name prop
// js:"name"
func (*SpeechSynthesisEvent) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// Utterance prop
// js:"utterance"
func (*SpeechSynthesisEvent) Utterance() (utterance *window.SpeechSynthesisUtterance) {
	macro.Rewrite("$_.utterance")
	return utterance
}

// Bubbles prop
// js:"bubbles"
func (*SpeechSynthesisEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*SpeechSynthesisEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*SpeechSynthesisEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*SpeechSynthesisEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*SpeechSynthesisEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*SpeechSynthesisEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*SpeechSynthesisEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*SpeechSynthesisEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*SpeechSynthesisEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*SpeechSynthesisEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*SpeechSynthesisEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*SpeechSynthesisEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*SpeechSynthesisEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*SpeechSynthesisEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
