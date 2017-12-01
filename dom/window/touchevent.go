package window

import "github.com/matthewmueller/golly/js"

var _ UIEvent = (*TouchEvent)(nil)
var _ Event = (*TouchEvent)(nil)

// TouchEvent struct
// js:"TouchEvent,omit"
type TouchEvent struct {
}

// InitUIEvent fn
// js:"initUIEvent"
func (*TouchEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	js.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

// InitEvent fn
// js:"initEvent"
func (*TouchEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*TouchEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*TouchEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*TouchEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// AltKey prop
// js:"altKey"
func (*TouchEvent) AltKey() (altKey bool) {
	js.Rewrite("$_.altKey")
	return altKey
}

// ChangedTouches prop
// js:"changedTouches"
func (*TouchEvent) ChangedTouches() (changedTouches *TouchList) {
	js.Rewrite("$_.changedTouches")
	return changedTouches
}

// CharCode prop
// js:"charCode"
func (*TouchEvent) CharCode() (charCode int8) {
	js.Rewrite("$_.charCode")
	return charCode
}

// CtrlKey prop
// js:"ctrlKey"
func (*TouchEvent) CtrlKey() (ctrlKey bool) {
	js.Rewrite("$_.ctrlKey")
	return ctrlKey
}

// KeyCode prop
// js:"keyCode"
func (*TouchEvent) KeyCode() (keyCode int8) {
	js.Rewrite("$_.keyCode")
	return keyCode
}

// MetaKey prop
// js:"metaKey"
func (*TouchEvent) MetaKey() (metaKey bool) {
	js.Rewrite("$_.metaKey")
	return metaKey
}

// ShiftKey prop
// js:"shiftKey"
func (*TouchEvent) ShiftKey() (shiftKey bool) {
	js.Rewrite("$_.shiftKey")
	return shiftKey
}

// TargetTouches prop
// js:"targetTouches"
func (*TouchEvent) TargetTouches() (targetTouches *TouchList) {
	js.Rewrite("$_.targetTouches")
	return targetTouches
}

// Touches prop
// js:"touches"
func (*TouchEvent) Touches() (touches *TouchList) {
	js.Rewrite("$_.touches")
	return touches
}

// Which prop
// js:"which"
func (*TouchEvent) Which() (which uint8) {
	js.Rewrite("$_.which")
	return which
}

// Detail prop
// js:"detail"
func (*TouchEvent) Detail() (detail int) {
	js.Rewrite("$_.detail")
	return detail
}

// View prop
// js:"view"
func (*TouchEvent) View() (view *Window) {
	js.Rewrite("$_.view")
	return view
}

// Bubbles prop
// js:"bubbles"
func (*TouchEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*TouchEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*TouchEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*TouchEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*TouchEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*TouchEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*TouchEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*TouchEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*TouchEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*TouchEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*TouchEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*TouchEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*TouchEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*TouchEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
