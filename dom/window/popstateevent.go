package window

import (
	"github.com/matthewmueller/joy/dom/popstateeventinit"
	"github.com/matthewmueller/joy/macro"
)

var _ Event = (*PopStateEvent)(nil)

// NewPopStateEvent fn
func NewPopStateEvent(typearg string, eventinitdict *popstateeventinit.PopStateEventInit) *PopStateEvent {
	macro.Rewrite("PopStateEvent")
	return &PopStateEvent{}
}

// PopStateEvent struct
// js:"PopStateEvent,omit"
type PopStateEvent struct {
}

// InitPopStateEvent fn
// js:"initPopStateEvent"
func (*PopStateEvent) InitPopStateEvent(typeArg string, canBubbleArg bool, cancelableArg bool, stateArg interface{}) {
	macro.Rewrite("$_.initPopStateEvent($1, $2, $3, $4)", typeArg, canBubbleArg, cancelableArg, stateArg)
}

// InitEvent fn
// js:"initEvent"
func (*PopStateEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*PopStateEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*PopStateEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*PopStateEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// State prop
// js:"state"
func (*PopStateEvent) State() (state interface{}) {
	macro.Rewrite("$_.state")
	return state
}

// Bubbles prop
// js:"bubbles"
func (*PopStateEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*PopStateEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*PopStateEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*PopStateEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*PopStateEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*PopStateEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*PopStateEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*PopStateEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*PopStateEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*PopStateEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*PopStateEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*PopStateEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*PopStateEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*PopStateEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
