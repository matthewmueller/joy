package window

import "github.com/matthewmueller/joy/macro"

var _ Event = (*IDBVersionChangeEvent)(nil)

// IDBVersionChangeEvent struct
// js:"IDBVersionChangeEvent,omit"
type IDBVersionChangeEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*IDBVersionChangeEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*IDBVersionChangeEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*IDBVersionChangeEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*IDBVersionChangeEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// NewVersion prop
// js:"newVersion"
func (*IDBVersionChangeEvent) NewVersion() (newVersion uint64) {
	macro.Rewrite("$_.newVersion")
	return newVersion
}

// OldVersion prop
// js:"oldVersion"
func (*IDBVersionChangeEvent) OldVersion() (oldVersion uint64) {
	macro.Rewrite("$_.oldVersion")
	return oldVersion
}

// Bubbles prop
// js:"bubbles"
func (*IDBVersionChangeEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*IDBVersionChangeEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*IDBVersionChangeEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*IDBVersionChangeEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*IDBVersionChangeEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*IDBVersionChangeEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*IDBVersionChangeEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*IDBVersionChangeEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*IDBVersionChangeEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*IDBVersionChangeEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*IDBVersionChangeEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*IDBVersionChangeEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*IDBVersionChangeEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*IDBVersionChangeEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
