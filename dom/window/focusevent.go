package window

import "github.com/matthewmueller/joy/macro"

var _ UIEvent = (*FocusEvent)(nil)
var _ Event = (*FocusEvent)(nil)

// NewFocusEvent fn
func NewFocusEvent(typearg string, eventinitdict *FocusEventInit) *FocusEvent {
	macro.Rewrite("FocusEvent")
	return &FocusEvent{}
}

// FocusEvent struct
// js:"FocusEvent,omit"
type FocusEvent struct {
}

// InitFocusEvent fn
// js:"initFocusEvent"
func (*FocusEvent) InitFocusEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, relatedTargetArg EventTarget) {
	macro.Rewrite("$_.initFocusEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, relatedTargetArg)
}

// InitUIEvent fn
// js:"initUIEvent"
func (*FocusEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

// InitEvent fn
// js:"initEvent"
func (*FocusEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*FocusEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*FocusEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*FocusEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// RelatedTarget prop
// js:"relatedTarget"
func (*FocusEvent) RelatedTarget() (relatedTarget EventTarget) {
	macro.Rewrite("$_.relatedTarget")
	return relatedTarget
}

// Detail prop
// js:"detail"
func (*FocusEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

// View prop
// js:"view"
func (*FocusEvent) View() (view *Window) {
	macro.Rewrite("$_.view")
	return view
}

// Bubbles prop
// js:"bubbles"
func (*FocusEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*FocusEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*FocusEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*FocusEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*FocusEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*FocusEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*FocusEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*FocusEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*FocusEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*FocusEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*FocusEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*FocusEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*FocusEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*FocusEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
