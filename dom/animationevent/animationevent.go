package animationevent

import (
	"github.com/matthewmueller/joy/dom/animationeventinit"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.Event = (*AnimationEvent)(nil)

// New fn
func New(typearg string, eventinitdict *animationeventinit.AnimationEventInit) *AnimationEvent {
	js.Rewrite("AnimationEvent")
	return &AnimationEvent{}
}

// AnimationEvent struct
// js:"AnimationEvent,omit"
type AnimationEvent struct {
}

// InitAnimationEvent fn
// js:"initAnimationEvent"
func (*AnimationEvent) InitAnimationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, animationNameArg string, elapsedTimeArg float32) {
	js.Rewrite("$_.initAnimationEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, animationNameArg, elapsedTimeArg)
}

// InitEvent fn
// js:"initEvent"
func (*AnimationEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*AnimationEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*AnimationEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*AnimationEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// AnimationName prop
// js:"animationName"
func (*AnimationEvent) AnimationName() (animationName string) {
	js.Rewrite("$_.animationName")
	return animationName
}

// ElapsedTime prop
// js:"elapsedTime"
func (*AnimationEvent) ElapsedTime() (elapsedTime float32) {
	js.Rewrite("$_.elapsedTime")
	return elapsedTime
}

// Bubbles prop
// js:"bubbles"
func (*AnimationEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*AnimationEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*AnimationEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*AnimationEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*AnimationEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*AnimationEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*AnimationEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*AnimationEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*AnimationEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*AnimationEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*AnimationEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*AnimationEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*AnimationEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*AnimationEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
