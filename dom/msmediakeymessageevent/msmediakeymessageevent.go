package msmediakeymessageevent

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.Event = (*MSMediaKeyMessageEvent)(nil)

// MSMediaKeyMessageEvent struct
// js:"MSMediaKeyMessageEvent,omit"
type MSMediaKeyMessageEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*MSMediaKeyMessageEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MSMediaKeyMessageEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MSMediaKeyMessageEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MSMediaKeyMessageEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// DestinationURL prop
// js:"destinationURL"
func (*MSMediaKeyMessageEvent) DestinationURL() (destinationURL string) {
	js.Rewrite("$_.destinationURL")
	return destinationURL
}

// Message prop
// js:"message"
func (*MSMediaKeyMessageEvent) Message() (message []uint8) {
	js.Rewrite("$_.message")
	return message
}

// Bubbles prop
// js:"bubbles"
func (*MSMediaKeyMessageEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MSMediaKeyMessageEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MSMediaKeyMessageEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MSMediaKeyMessageEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MSMediaKeyMessageEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MSMediaKeyMessageEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MSMediaKeyMessageEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MSMediaKeyMessageEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MSMediaKeyMessageEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MSMediaKeyMessageEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MSMediaKeyMessageEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MSMediaKeyMessageEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MSMediaKeyMessageEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MSMediaKeyMessageEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
