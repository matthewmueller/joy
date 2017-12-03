package transitionevent

import (
	"github.com/matthewmueller/joy/dom/transitioneventinit"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*TransitionEvent)(nil)

// New fn
func New(typearg string, eventinitdict *transitioneventinit.TransitionEventInit) *TransitionEvent {
	macro.Rewrite("TransitionEvent")
	return &TransitionEvent{}
}

// TransitionEvent struct
// js:"TransitionEvent,omit"
type TransitionEvent struct {
}

// InitTransitionEvent fn
// js:"initTransitionEvent"
func (*TransitionEvent) InitTransitionEvent(typeArg string, canBubbleArg bool, cancelableArg bool, propertyNameArg string, elapsedTimeArg float32) {
	macro.Rewrite("$_.initTransitionEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, propertyNameArg, elapsedTimeArg)
}

// InitEvent fn
// js:"initEvent"
func (*TransitionEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*TransitionEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*TransitionEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*TransitionEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// ElapsedTime prop
// js:"elapsedTime"
func (*TransitionEvent) ElapsedTime() (elapsedTime float32) {
	macro.Rewrite("$_.elapsedTime")
	return elapsedTime
}

// PropertyName prop
// js:"propertyName"
func (*TransitionEvent) PropertyName() (propertyName string) {
	macro.Rewrite("$_.propertyName")
	return propertyName
}

// Bubbles prop
// js:"bubbles"
func (*TransitionEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*TransitionEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*TransitionEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*TransitionEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*TransitionEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*TransitionEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*TransitionEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*TransitionEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*TransitionEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*TransitionEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*TransitionEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*TransitionEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*TransitionEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*TransitionEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
