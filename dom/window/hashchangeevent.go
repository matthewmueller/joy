package window

import (
	"github.com/matthewmueller/joy/dom/hashchangeeventinit"
	"github.com/matthewmueller/joy/macro"
)

var _ Event = (*HashChangeEvent)(nil)

// NewHashChangeEvent fn
func NewHashChangeEvent(typearg string, eventinitdict *hashchangeeventinit.HashChangeEventInit) *HashChangeEvent {
	macro.Rewrite("new HashChangeEvent($1, $2)", typearg, eventinitdict)
	return &HashChangeEvent{}
}

// HashChangeEvent struct
// js:"HashChangeEvent,omit"
type HashChangeEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*HashChangeEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*HashChangeEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*HashChangeEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*HashChangeEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// NewURL prop
// js:"newURL"
func (*HashChangeEvent) NewURL() (newURL string) {
	macro.Rewrite("$_.newURL")
	return newURL
}

// OldURL prop
// js:"oldURL"
func (*HashChangeEvent) OldURL() (oldURL string) {
	macro.Rewrite("$_.oldURL")
	return oldURL
}

// Bubbles prop
// js:"bubbles"
func (*HashChangeEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*HashChangeEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*HashChangeEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*HashChangeEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*HashChangeEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*HashChangeEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*HashChangeEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*HashChangeEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*HashChangeEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*HashChangeEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*HashChangeEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*HashChangeEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*HashChangeEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*HashChangeEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
