package window

import (
	"github.com/matthewmueller/joy/dom/mediastreamerror"
	"github.com/matthewmueller/joy/dom/mediastreamerroreventinit"
	"github.com/matthewmueller/joy/js"
)

var _ Event = (*MediaStreamErrorEvent)(nil)

// NewMediaStreamErrorEvent fn
func NewMediaStreamErrorEvent(typearg string, eventinitdict *mediastreamerroreventinit.MediaStreamErrorEventInit) *MediaStreamErrorEvent {
	js.Rewrite("MediaStreamErrorEvent")
	return &MediaStreamErrorEvent{}
}

// MediaStreamErrorEvent struct
// js:"MediaStreamErrorEvent,omit"
type MediaStreamErrorEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*MediaStreamErrorEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MediaStreamErrorEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MediaStreamErrorEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MediaStreamErrorEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// Error prop
// js:"error"
func (*MediaStreamErrorEvent) Error() (err *mediastreamerror.MediaStreamError) {
	js.Rewrite("$_.error")
	return err
}

// Bubbles prop
// js:"bubbles"
func (*MediaStreamErrorEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MediaStreamErrorEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MediaStreamErrorEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MediaStreamErrorEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MediaStreamErrorEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MediaStreamErrorEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MediaStreamErrorEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MediaStreamErrorEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MediaStreamErrorEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MediaStreamErrorEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MediaStreamErrorEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MediaStreamErrorEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MediaStreamErrorEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MediaStreamErrorEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
