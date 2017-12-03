package window

import "github.com/matthewmueller/joy/macro"

var _ UIEvent = (*MSGestureEvent)(nil)
var _ Event = (*MSGestureEvent)(nil)

// MSGestureEvent struct
// js:"MSGestureEvent,omit"
type MSGestureEvent struct {
}

// InitGestureEvent fn
// js:"initGestureEvent"
func (*MSGestureEvent) InitGestureEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, offsetXArg float32, offsetYArg float32, translationXArg float32, translationYArg float32, scaleArg float32, expansionArg float32, rotationArg float32, velocityXArg float32, velocityYArg float32, velocityExpansionArg float32, velocityAngularArg float32, hwTimestampArg uint64) {
	macro.Rewrite("$_.initGestureEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, offsetXArg, offsetYArg, translationXArg, translationYArg, scaleArg, expansionArg, rotationArg, velocityXArg, velocityYArg, velocityExpansionArg, velocityAngularArg, hwTimestampArg)
}

// InitUIEvent fn
// js:"initUIEvent"
func (*MSGestureEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

// InitEvent fn
// js:"initEvent"
func (*MSGestureEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MSGestureEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MSGestureEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MSGestureEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// ClientX prop
// js:"clientX"
func (*MSGestureEvent) ClientX() (clientX float32) {
	macro.Rewrite("$_.clientX")
	return clientX
}

// ClientY prop
// js:"clientY"
func (*MSGestureEvent) ClientY() (clientY float32) {
	macro.Rewrite("$_.clientY")
	return clientY
}

// Expansion prop
// js:"expansion"
func (*MSGestureEvent) Expansion() (expansion float32) {
	macro.Rewrite("$_.expansion")
	return expansion
}

// GestureObject prop
// js:"gestureObject"
func (*MSGestureEvent) GestureObject() (gestureObject interface{}) {
	macro.Rewrite("$_.gestureObject")
	return gestureObject
}

// HwTimestamp prop
// js:"hwTimestamp"
func (*MSGestureEvent) HwTimestamp() (hwTimestamp uint64) {
	macro.Rewrite("$_.hwTimestamp")
	return hwTimestamp
}

// OffsetX prop
// js:"offsetX"
func (*MSGestureEvent) OffsetX() (offsetX float32) {
	macro.Rewrite("$_.offsetX")
	return offsetX
}

// OffsetY prop
// js:"offsetY"
func (*MSGestureEvent) OffsetY() (offsetY float32) {
	macro.Rewrite("$_.offsetY")
	return offsetY
}

// Rotation prop
// js:"rotation"
func (*MSGestureEvent) Rotation() (rotation float32) {
	macro.Rewrite("$_.rotation")
	return rotation
}

// Scale prop
// js:"scale"
func (*MSGestureEvent) Scale() (scale float32) {
	macro.Rewrite("$_.scale")
	return scale
}

// ScreenX prop
// js:"screenX"
func (*MSGestureEvent) ScreenX() (screenX int) {
	macro.Rewrite("$_.screenX")
	return screenX
}

// ScreenY prop
// js:"screenY"
func (*MSGestureEvent) ScreenY() (screenY int) {
	macro.Rewrite("$_.screenY")
	return screenY
}

// TranslationX prop
// js:"translationX"
func (*MSGestureEvent) TranslationX() (translationX float32) {
	macro.Rewrite("$_.translationX")
	return translationX
}

// TranslationY prop
// js:"translationY"
func (*MSGestureEvent) TranslationY() (translationY float32) {
	macro.Rewrite("$_.translationY")
	return translationY
}

// VelocityAngular prop
// js:"velocityAngular"
func (*MSGestureEvent) VelocityAngular() (velocityAngular float32) {
	macro.Rewrite("$_.velocityAngular")
	return velocityAngular
}

// VelocityExpansion prop
// js:"velocityExpansion"
func (*MSGestureEvent) VelocityExpansion() (velocityExpansion float32) {
	macro.Rewrite("$_.velocityExpansion")
	return velocityExpansion
}

// VelocityX prop
// js:"velocityX"
func (*MSGestureEvent) VelocityX() (velocityX float32) {
	macro.Rewrite("$_.velocityX")
	return velocityX
}

// VelocityY prop
// js:"velocityY"
func (*MSGestureEvent) VelocityY() (velocityY float32) {
	macro.Rewrite("$_.velocityY")
	return velocityY
}

// Detail prop
// js:"detail"
func (*MSGestureEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

// View prop
// js:"view"
func (*MSGestureEvent) View() (view *Window) {
	macro.Rewrite("$_.view")
	return view
}

// Bubbles prop
// js:"bubbles"
func (*MSGestureEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MSGestureEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MSGestureEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MSGestureEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MSGestureEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MSGestureEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MSGestureEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MSGestureEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MSGestureEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MSGestureEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MSGestureEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MSGestureEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MSGestureEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MSGestureEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
