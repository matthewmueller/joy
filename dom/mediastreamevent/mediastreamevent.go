package mediastreamevent

import (
	"github.com/matthewmueller/joy/dom/mediastreameventinit"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.Event = (*MediaStreamEvent)(nil)

// New fn
func New(kind string, eventinitdict *mediastreameventinit.MediaStreamEventInit) *MediaStreamEvent {
	js.Rewrite("MediaStreamEvent")
	return &MediaStreamEvent{}
}

// MediaStreamEvent struct
// js:"MediaStreamEvent,omit"
type MediaStreamEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*MediaStreamEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MediaStreamEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MediaStreamEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MediaStreamEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// Stream prop
// js:"stream"
func (*MediaStreamEvent) Stream() (stream *window.MediaStream) {
	js.Rewrite("$_.stream")
	return stream
}

// Bubbles prop
// js:"bubbles"
func (*MediaStreamEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MediaStreamEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MediaStreamEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MediaStreamEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MediaStreamEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MediaStreamEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MediaStreamEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MediaStreamEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MediaStreamEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MediaStreamEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MediaStreamEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MediaStreamEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MediaStreamEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MediaStreamEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
