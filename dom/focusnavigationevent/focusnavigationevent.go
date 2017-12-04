package focusnavigationevent

import (
	"github.com/matthewmueller/joy/dom/focusnavigationeventinit"
	"github.com/matthewmueller/joy/dom/navigationreason"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*FocusNavigationEvent)(nil)

// New fn
func New(kind string, eventinitdict *focusnavigationeventinit.FocusNavigationEventInit) *FocusNavigationEvent {
	macro.Rewrite("new FocusNavigationEvent($1, $2)", kind, eventinitdict)
	return &FocusNavigationEvent{}
}

// FocusNavigationEvent struct
// js:"FocusNavigationEvent,omit"
type FocusNavigationEvent struct {
}

// RequestFocus fn
// js:"requestFocus"
func (*FocusNavigationEvent) RequestFocus() {
	macro.Rewrite("$_.requestFocus()")
}

// InitEvent fn
// js:"initEvent"
func (*FocusNavigationEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*FocusNavigationEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*FocusNavigationEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*FocusNavigationEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// NavigationReason prop
// js:"navigationReason"
func (*FocusNavigationEvent) NavigationReason() (navigationReason *navigationreason.NavigationReason) {
	macro.Rewrite("$_.navigationReason")
	return navigationReason
}

// OriginHeight prop
// js:"originHeight"
func (*FocusNavigationEvent) OriginHeight() (originHeight float32) {
	macro.Rewrite("$_.originHeight")
	return originHeight
}

// OriginLeft prop
// js:"originLeft"
func (*FocusNavigationEvent) OriginLeft() (originLeft float32) {
	macro.Rewrite("$_.originLeft")
	return originLeft
}

// OriginTop prop
// js:"originTop"
func (*FocusNavigationEvent) OriginTop() (originTop float32) {
	macro.Rewrite("$_.originTop")
	return originTop
}

// OriginWidth prop
// js:"originWidth"
func (*FocusNavigationEvent) OriginWidth() (originWidth float32) {
	macro.Rewrite("$_.originWidth")
	return originWidth
}

// Bubbles prop
// js:"bubbles"
func (*FocusNavigationEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*FocusNavigationEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*FocusNavigationEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*FocusNavigationEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*FocusNavigationEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*FocusNavigationEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*FocusNavigationEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*FocusNavigationEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*FocusNavigationEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*FocusNavigationEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*FocusNavigationEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*FocusNavigationEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*FocusNavigationEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*FocusNavigationEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
