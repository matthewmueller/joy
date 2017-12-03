package mutationevent

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.Event = (*MutationEvent)(nil)

// MutationEvent struct
// js:"MutationEvent,omit"
type MutationEvent struct {
}

// InitMutationEvent fn
// js:"initMutationEvent"
func (*MutationEvent) InitMutationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, relatedNodeArg window.Node, prevValueArg string, newValueArg string, attrNameArg string, attrChangeArg uint8) {
	js.Rewrite("$_.initMutationEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, relatedNodeArg, prevValueArg, newValueArg, attrNameArg, attrChangeArg)
}

// InitEvent fn
// js:"initEvent"
func (*MutationEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MutationEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MutationEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MutationEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// AttrChange prop
// js:"attrChange"
func (*MutationEvent) AttrChange() (attrChange uint8) {
	js.Rewrite("$_.attrChange")
	return attrChange
}

// AttrName prop
// js:"attrName"
func (*MutationEvent) AttrName() (attrName string) {
	js.Rewrite("$_.attrName")
	return attrName
}

// NewValue prop
// js:"newValue"
func (*MutationEvent) NewValue() (newValue string) {
	js.Rewrite("$_.newValue")
	return newValue
}

// PrevValue prop
// js:"prevValue"
func (*MutationEvent) PrevValue() (prevValue string) {
	js.Rewrite("$_.prevValue")
	return prevValue
}

// RelatedNode prop
// js:"relatedNode"
func (*MutationEvent) RelatedNode() (relatedNode window.Node) {
	js.Rewrite("$_.relatedNode")
	return relatedNode
}

// Bubbles prop
// js:"bubbles"
func (*MutationEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MutationEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MutationEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MutationEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MutationEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MutationEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MutationEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MutationEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MutationEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MutationEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MutationEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MutationEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MutationEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MutationEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
