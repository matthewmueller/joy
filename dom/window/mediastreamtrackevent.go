package window

import "github.com/matthewmueller/joy/macro"

var _ Event = (*MediaStreamTrackEvent)(nil)

// NewMediaStreamTrackEvent fn
func NewMediaStreamTrackEvent(typearg string, eventinitdict *MediaStreamTrackEventInit) *MediaStreamTrackEvent {
	macro.Rewrite("MediaStreamTrackEvent")
	return &MediaStreamTrackEvent{}
}

// MediaStreamTrackEvent struct
// js:"MediaStreamTrackEvent,omit"
type MediaStreamTrackEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*MediaStreamTrackEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MediaStreamTrackEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MediaStreamTrackEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MediaStreamTrackEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// Track prop
// js:"track"
func (*MediaStreamTrackEvent) Track() (track *MediaStreamTrack) {
	macro.Rewrite("$_.track")
	return track
}

// Bubbles prop
// js:"bubbles"
func (*MediaStreamTrackEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MediaStreamTrackEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MediaStreamTrackEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MediaStreamTrackEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MediaStreamTrackEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MediaStreamTrackEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MediaStreamTrackEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MediaStreamTrackEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MediaStreamTrackEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MediaStreamTrackEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MediaStreamTrackEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MediaStreamTrackEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MediaStreamTrackEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MediaStreamTrackEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
