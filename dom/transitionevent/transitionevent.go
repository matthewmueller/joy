package transitionevent

import (
	"github.com/matthewmueller/golly/dom/transitioneventinit"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.Event = (*TransitionEvent)(nil)

// New fn
func New(typearg string, eventinitdict *transitioneventinit.TransitionEventInit) *TransitionEvent {
	js.Rewrite("TransitionEvent")
	return &TransitionEvent{}
}

// TransitionEvent struct
// js:"TransitionEvent,omit"
type TransitionEvent struct {
}

// InitTransitionEvent fn
// js:"initTransitionEvent"
func (*TransitionEvent) InitTransitionEvent(typeArg string, canBubbleArg bool, cancelableArg bool, propertyNameArg string, elapsedTimeArg float32) {
	js.Rewrite("$_.initTransitionEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, propertyNameArg, elapsedTimeArg)
}

// InitEvent fn
// js:"initEvent"
func (*TransitionEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*TransitionEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*TransitionEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*TransitionEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// ElapsedTime prop
// js:"elapsedTime"
func (*TransitionEvent) ElapsedTime() (elapsedTime float32) {
	js.Rewrite("$_.elapsedTime")
	return elapsedTime
}

// PropertyName prop
// js:"propertyName"
func (*TransitionEvent) PropertyName() (propertyName string) {
	js.Rewrite("$_.propertyName")
	return propertyName
}

// Bubbles prop
// js:"bubbles"
func (*TransitionEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*TransitionEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*TransitionEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*TransitionEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*TransitionEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*TransitionEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*TransitionEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*TransitionEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*TransitionEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*TransitionEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*TransitionEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*TransitionEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*TransitionEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*TransitionEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
