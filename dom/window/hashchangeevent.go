package window

import (
	"github.com/matthewmueller/golly/dom/hashchangeeventinit"
	"github.com/matthewmueller/golly/js"
)

var _ Event = (*HashChangeEvent)(nil)

// NewHashChangeEvent fn
func NewHashChangeEvent(typearg string, eventinitdict *hashchangeeventinit.HashChangeEventInit) *HashChangeEvent {
	js.Rewrite("HashChangeEvent")
	return &HashChangeEvent{}
}

// HashChangeEvent struct
// js:"HashChangeEvent,omit"
type HashChangeEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*HashChangeEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*HashChangeEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*HashChangeEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*HashChangeEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// NewURL prop
// js:"newURL"
func (*HashChangeEvent) NewURL() (newURL string) {
	js.Rewrite("$_.newURL")
	return newURL
}

// OldURL prop
// js:"oldURL"
func (*HashChangeEvent) OldURL() (oldURL string) {
	js.Rewrite("$_.oldURL")
	return oldURL
}

// Bubbles prop
// js:"bubbles"
func (*HashChangeEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*HashChangeEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*HashChangeEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*HashChangeEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*HashChangeEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*HashChangeEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*HashChangeEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*HashChangeEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*HashChangeEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*HashChangeEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*HashChangeEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*HashChangeEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*HashChangeEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*HashChangeEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
