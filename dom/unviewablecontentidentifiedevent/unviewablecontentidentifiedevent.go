package unviewablecontentidentifiedevent

import (
	"github.com/matthewmueller/golly/dom/navigationevent"
	"github.com/matthewmueller/golly/dom/navigationeventwithreferrer"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ navigationeventwithreferrer.NavigationEventWithReferrer = (*UnviewableContentIdentifiedEvent)(nil)
var _ navigationevent.NavigationEvent = (*UnviewableContentIdentifiedEvent)(nil)
var _ window.Event = (*UnviewableContentIdentifiedEvent)(nil)

// UnviewableContentIdentifiedEvent struct
// js:"UnviewableContentIdentifiedEvent,omit"
type UnviewableContentIdentifiedEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*UnviewableContentIdentifiedEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*UnviewableContentIdentifiedEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*UnviewableContentIdentifiedEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*UnviewableContentIdentifiedEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// MediaType prop
// js:"mediaType"
func (*UnviewableContentIdentifiedEvent) MediaType() (mediaType string) {
	js.Rewrite("$_.mediaType")
	return mediaType
}

// Referer prop
// js:"referer"
func (*UnviewableContentIdentifiedEvent) Referer() (referer string) {
	js.Rewrite("$_.referer")
	return referer
}

// URI prop
// js:"uri"
func (*UnviewableContentIdentifiedEvent) URI() (uri string) {
	js.Rewrite("$_.uri")
	return uri
}

// Bubbles prop
// js:"bubbles"
func (*UnviewableContentIdentifiedEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*UnviewableContentIdentifiedEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*UnviewableContentIdentifiedEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*UnviewableContentIdentifiedEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*UnviewableContentIdentifiedEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*UnviewableContentIdentifiedEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*UnviewableContentIdentifiedEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*UnviewableContentIdentifiedEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*UnviewableContentIdentifiedEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*UnviewableContentIdentifiedEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*UnviewableContentIdentifiedEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*UnviewableContentIdentifiedEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*UnviewableContentIdentifiedEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*UnviewableContentIdentifiedEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
