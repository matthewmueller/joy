package rtcdtlstransportstatechangedevent

import (
	"github.com/matthewmueller/joy/dom/rtcdtlstransportstate"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.Event = (*RTCDtlsTransportStateChangedEvent)(nil)

// RTCDtlsTransportStateChangedEvent struct
// js:"RTCDtlsTransportStateChangedEvent,omit"
type RTCDtlsTransportStateChangedEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*RTCDtlsTransportStateChangedEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*RTCDtlsTransportStateChangedEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*RTCDtlsTransportStateChangedEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*RTCDtlsTransportStateChangedEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// State prop
// js:"state"
func (*RTCDtlsTransportStateChangedEvent) State() (state *rtcdtlstransportstate.RTCDtlsTransportState) {
	js.Rewrite("$_.state")
	return state
}

// Bubbles prop
// js:"bubbles"
func (*RTCDtlsTransportStateChangedEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*RTCDtlsTransportStateChangedEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*RTCDtlsTransportStateChangedEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*RTCDtlsTransportStateChangedEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*RTCDtlsTransportStateChangedEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*RTCDtlsTransportStateChangedEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*RTCDtlsTransportStateChangedEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*RTCDtlsTransportStateChangedEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*RTCDtlsTransportStateChangedEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*RTCDtlsTransportStateChangedEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*RTCDtlsTransportStateChangedEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*RTCDtlsTransportStateChangedEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*RTCDtlsTransportStateChangedEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*RTCDtlsTransportStateChangedEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
