package mutationevent

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*MutationEvent)(nil)

// MutationEvent struct
// js:"MutationEvent,omit"
type MutationEvent struct {
}

// InitMutationEvent fn
// js:"initMutationEvent"
func (*MutationEvent) InitMutationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, relatedNodeArg window.Node, prevValueArg string, newValueArg string, attrNameArg string, attrChangeArg uint8) {
	macro.Rewrite("$_.initMutationEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, relatedNodeArg, prevValueArg, newValueArg, attrNameArg, attrChangeArg)
}

// InitEvent fn
// js:"initEvent"
func (*MutationEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MutationEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MutationEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MutationEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// AttrChange prop
// js:"attrChange"
func (*MutationEvent) AttrChange() (attrChange uint8) {
	macro.Rewrite("$_.attrChange")
	return attrChange
}

// AttrName prop
// js:"attrName"
func (*MutationEvent) AttrName() (attrName string) {
	macro.Rewrite("$_.attrName")
	return attrName
}

// NewValue prop
// js:"newValue"
func (*MutationEvent) NewValue() (newValue string) {
	macro.Rewrite("$_.newValue")
	return newValue
}

// PrevValue prop
// js:"prevValue"
func (*MutationEvent) PrevValue() (prevValue string) {
	macro.Rewrite("$_.prevValue")
	return prevValue
}

// RelatedNode prop
// js:"relatedNode"
func (*MutationEvent) RelatedNode() (relatedNode window.Node) {
	macro.Rewrite("$_.relatedNode")
	return relatedNode
}

// Bubbles prop
// js:"bubbles"
func (*MutationEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MutationEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MutationEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MutationEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MutationEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MutationEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MutationEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MutationEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MutationEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MutationEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MutationEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MutationEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MutationEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MutationEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
