package window

import "github.com/matthewmueller/joy/macro"

var _ UIEvent = (*TouchEvent)(nil)
var _ Event = (*TouchEvent)(nil)

// TouchEvent struct
// js:"TouchEvent,omit"
type TouchEvent struct {
}

// InitUIEvent fn
// js:"initUIEvent"
func (*TouchEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

// InitEvent fn
// js:"initEvent"
func (*TouchEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*TouchEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*TouchEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*TouchEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// AltKey prop
// js:"altKey"
func (*TouchEvent) AltKey() (altKey bool) {
	macro.Rewrite("$_.altKey")
	return altKey
}

// ChangedTouches prop
// js:"changedTouches"
func (*TouchEvent) ChangedTouches() (changedTouches *TouchList) {
	macro.Rewrite("$_.changedTouches")
	return changedTouches
}

// CharCode prop
// js:"charCode"
func (*TouchEvent) CharCode() (charCode int8) {
	macro.Rewrite("$_.charCode")
	return charCode
}

// CtrlKey prop
// js:"ctrlKey"
func (*TouchEvent) CtrlKey() (ctrlKey bool) {
	macro.Rewrite("$_.ctrlKey")
	return ctrlKey
}

// KeyCode prop
// js:"keyCode"
func (*TouchEvent) KeyCode() (keyCode int8) {
	macro.Rewrite("$_.keyCode")
	return keyCode
}

// MetaKey prop
// js:"metaKey"
func (*TouchEvent) MetaKey() (metaKey bool) {
	macro.Rewrite("$_.metaKey")
	return metaKey
}

// ShiftKey prop
// js:"shiftKey"
func (*TouchEvent) ShiftKey() (shiftKey bool) {
	macro.Rewrite("$_.shiftKey")
	return shiftKey
}

// TargetTouches prop
// js:"targetTouches"
func (*TouchEvent) TargetTouches() (targetTouches *TouchList) {
	macro.Rewrite("$_.targetTouches")
	return targetTouches
}

// Touches prop
// js:"touches"
func (*TouchEvent) Touches() (touches *TouchList) {
	macro.Rewrite("$_.touches")
	return touches
}

// Which prop
// js:"which"
func (*TouchEvent) Which() (which uint8) {
	macro.Rewrite("$_.which")
	return which
}

// Detail prop
// js:"detail"
func (*TouchEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

// View prop
// js:"view"
func (*TouchEvent) View() (view *Window) {
	macro.Rewrite("$_.view")
	return view
}

// Bubbles prop
// js:"bubbles"
func (*TouchEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*TouchEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*TouchEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*TouchEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*TouchEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*TouchEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*TouchEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*TouchEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*TouchEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*TouchEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*TouchEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*TouchEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*TouchEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*TouchEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
