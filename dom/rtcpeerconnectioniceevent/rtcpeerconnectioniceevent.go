package rtcpeerconnectioniceevent

import (
	"github.com/matthewmueller/joy/dom/rtcicecandidate"
	"github.com/matthewmueller/joy/dom/rtcpeerconnectioniceeventinit"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.Event = (*RTCPeerConnectionIceEvent)(nil)

// New fn
func New(kind string, eventinitdict *rtcpeerconnectioniceeventinit.RTCPeerConnectionIceEventInit) *RTCPeerConnectionIceEvent {
	js.Rewrite("RTCPeerConnectionIceEvent")
	return &RTCPeerConnectionIceEvent{}
}

// RTCPeerConnectionIceEvent struct
// js:"RTCPeerConnectionIceEvent,omit"
type RTCPeerConnectionIceEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*RTCPeerConnectionIceEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*RTCPeerConnectionIceEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*RTCPeerConnectionIceEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*RTCPeerConnectionIceEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// Candidate prop
// js:"candidate"
func (*RTCPeerConnectionIceEvent) Candidate() (candidate *rtcicecandidate.RTCIceCandidate) {
	js.Rewrite("$_.candidate")
	return candidate
}

// Bubbles prop
// js:"bubbles"
func (*RTCPeerConnectionIceEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*RTCPeerConnectionIceEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*RTCPeerConnectionIceEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*RTCPeerConnectionIceEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*RTCPeerConnectionIceEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*RTCPeerConnectionIceEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*RTCPeerConnectionIceEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*RTCPeerConnectionIceEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*RTCPeerConnectionIceEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*RTCPeerConnectionIceEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*RTCPeerConnectionIceEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*RTCPeerConnectionIceEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*RTCPeerConnectionIceEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*RTCPeerConnectionIceEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
