package mediakeymessageevent

import (
	"github.com/matthewmueller/joy/dom/mediakeymessageeventinit"
	"github.com/matthewmueller/joy/dom/mediakeymessagetype"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*MediaKeyMessageEvent)(nil)

// New fn
func New(kind string, eventinitdict *mediakeymessageeventinit.MediaKeyMessageEventInit) *MediaKeyMessageEvent {
	macro.Rewrite("MediaKeyMessageEvent")
	return &MediaKeyMessageEvent{}
}

// MediaKeyMessageEvent struct
// js:"MediaKeyMessageEvent,omit"
type MediaKeyMessageEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*MediaKeyMessageEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MediaKeyMessageEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MediaKeyMessageEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MediaKeyMessageEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// Message prop
// js:"message"
func (*MediaKeyMessageEvent) Message() (message []byte) {
	macro.Rewrite("$_.message")
	return message
}

// MessageType prop
// js:"messageType"
func (*MediaKeyMessageEvent) MessageType() (messageType *mediakeymessagetype.MediaKeyMessageType) {
	macro.Rewrite("$_.messageType")
	return messageType
}

// Bubbles prop
// js:"bubbles"
func (*MediaKeyMessageEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MediaKeyMessageEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MediaKeyMessageEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MediaKeyMessageEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MediaKeyMessageEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MediaKeyMessageEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MediaKeyMessageEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MediaKeyMessageEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MediaKeyMessageEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MediaKeyMessageEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MediaKeyMessageEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MediaKeyMessageEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MediaKeyMessageEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MediaKeyMessageEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
