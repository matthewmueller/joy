package window

import "github.com/matthewmueller/golly/js"

var _ Event = (*MediaStreamTrackEvent)(nil)

// NewMediaStreamTrackEvent fn
func NewMediaStreamTrackEvent(typearg string, eventinitdict *MediaStreamTrackEventInit) *MediaStreamTrackEvent {
	js.Rewrite("MediaStreamTrackEvent")
	return &MediaStreamTrackEvent{}
}

// MediaStreamTrackEvent struct
// js:"MediaStreamTrackEvent,omit"
type MediaStreamTrackEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*MediaStreamTrackEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MediaStreamTrackEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MediaStreamTrackEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MediaStreamTrackEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// Track prop
// js:"track"
func (*MediaStreamTrackEvent) Track() (track *MediaStreamTrack) {
	js.Rewrite("$_.track")
	return track
}

// Bubbles prop
// js:"bubbles"
func (*MediaStreamTrackEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MediaStreamTrackEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MediaStreamTrackEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MediaStreamTrackEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MediaStreamTrackEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MediaStreamTrackEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MediaStreamTrackEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MediaStreamTrackEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MediaStreamTrackEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MediaStreamTrackEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MediaStreamTrackEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MediaStreamTrackEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MediaStreamTrackEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MediaStreamTrackEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
