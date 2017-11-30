package focusnavigationevent

import (
	"github.com/matthewmueller/golly/dom/navigationreason"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.Event = (*FocusNavigationEvent)(nil)

// FocusNavigationEvent struct
// js:"FocusNavigationEvent,omit"
type FocusNavigationEvent struct {
}

// RequestFocus fn
// js:"requestFocus"
func (*FocusNavigationEvent) RequestFocus() {
	js.Rewrite("$_.requestFocus()")
}

// InitEvent fn
// js:"initEvent"
func (*FocusNavigationEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*FocusNavigationEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*FocusNavigationEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*FocusNavigationEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// NavigationReason prop
// js:"navigationReason"
func (*FocusNavigationEvent) NavigationReason() (navigationReason *navigationreason.NavigationReason) {
	js.Rewrite("$_.navigationReason")
	return navigationReason
}

// OriginHeight prop
// js:"originHeight"
func (*FocusNavigationEvent) OriginHeight() (originHeight float32) {
	js.Rewrite("$_.originHeight")
	return originHeight
}

// OriginLeft prop
// js:"originLeft"
func (*FocusNavigationEvent) OriginLeft() (originLeft float32) {
	js.Rewrite("$_.originLeft")
	return originLeft
}

// OriginTop prop
// js:"originTop"
func (*FocusNavigationEvent) OriginTop() (originTop float32) {
	js.Rewrite("$_.originTop")
	return originTop
}

// OriginWidth prop
// js:"originWidth"
func (*FocusNavigationEvent) OriginWidth() (originWidth float32) {
	js.Rewrite("$_.originWidth")
	return originWidth
}

// Bubbles prop
// js:"bubbles"
func (*FocusNavigationEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*FocusNavigationEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*FocusNavigationEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*FocusNavigationEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*FocusNavigationEvent) CurrentTarget() (currentTarget window.EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*FocusNavigationEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*FocusNavigationEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*FocusNavigationEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*FocusNavigationEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*FocusNavigationEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*FocusNavigationEvent) SrcElement() (srcElement window.Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*FocusNavigationEvent) Target() (target window.EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*FocusNavigationEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*FocusNavigationEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
