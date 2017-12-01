package window

import "github.com/matthewmueller/golly/js"

var _ MouseEvent = (*WheelEvent)(nil)
var _ UIEvent = (*WheelEvent)(nil)
var _ Event = (*WheelEvent)(nil)

// NewWheelEvent fn
func NewWheelEvent(typearg string, eventinitdict *WheelEventInit) *WheelEvent {
	js.Rewrite("WheelEvent")
	return &WheelEvent{}
}

// WheelEvent struct
// js:"WheelEvent,omit"
type WheelEvent struct {
}

// GetCurrentPoint fn
// js:"getCurrentPoint"
func (*WheelEvent) GetCurrentPoint(element Element) {
	js.Rewrite("$_.getCurrentPoint($1)", element)
}

// InitWheelEvent fn
// js:"initWheelEvent"
func (*WheelEvent) InitWheelEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, buttonArg uint8, relatedTargetArg EventTarget, modifiersListArg string, deltaXArg int, deltaYArg int, deltaZArg int, deltaMode uint) {
	js.Rewrite("$_.initWheelEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, buttonArg, relatedTargetArg, modifiersListArg, deltaXArg, deltaYArg, deltaZArg, deltaMode)
}

// GetModifierState fn
// js:"getModifierState"
func (*WheelEvent) GetModifierState(keyArg string) (b bool) {
	js.Rewrite("$_.getModifierState($1)", keyArg)
	return b
}

// InitMouseEvent fn
// js:"initMouseEvent"
func (*WheelEvent) InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget) {
	js.Rewrite("$_.initMouseEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg)
}

// InitUIEvent fn
// js:"initUIEvent"
func (*WheelEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	js.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

// InitEvent fn
// js:"initEvent"
func (*WheelEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*WheelEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*WheelEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*WheelEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// DeltaMode prop
// js:"deltaMode"
func (*WheelEvent) DeltaMode() (deltaMode uint) {
	js.Rewrite("$_.deltaMode")
	return deltaMode
}

// DeltaX prop
// js:"deltaX"
func (*WheelEvent) DeltaX() (deltaX int) {
	js.Rewrite("$_.deltaX")
	return deltaX
}

// DeltaY prop
// js:"deltaY"
func (*WheelEvent) DeltaY() (deltaY int) {
	js.Rewrite("$_.deltaY")
	return deltaY
}

// DeltaZ prop
// js:"deltaZ"
func (*WheelEvent) DeltaZ() (deltaZ int) {
	js.Rewrite("$_.deltaZ")
	return deltaZ
}

// WheelDelta prop
// js:"wheelDelta"
func (*WheelEvent) WheelDelta() (wheelDelta int) {
	js.Rewrite("$_.wheelDelta")
	return wheelDelta
}

// WheelDeltaX prop
// js:"wheelDeltaX"
func (*WheelEvent) WheelDeltaX() (wheelDeltaX int) {
	js.Rewrite("$_.wheelDeltaX")
	return wheelDeltaX
}

// WheelDeltaY prop
// js:"wheelDeltaY"
func (*WheelEvent) WheelDeltaY() (wheelDeltaY int) {
	js.Rewrite("$_.wheelDeltaY")
	return wheelDeltaY
}

// AltKey prop
// js:"altKey"
func (*WheelEvent) AltKey() (altKey bool) {
	js.Rewrite("$_.altKey")
	return altKey
}

// Button prop
// js:"button"
func (*WheelEvent) Button() (button uint8) {
	js.Rewrite("$_.button")
	return button
}

// Buttons prop
// js:"buttons"
func (*WheelEvent) Buttons() (buttons uint8) {
	js.Rewrite("$_.buttons")
	return buttons
}

// ClientX prop
// js:"clientX"
func (*WheelEvent) ClientX() (clientX int) {
	js.Rewrite("$_.clientX")
	return clientX
}

// ClientY prop
// js:"clientY"
func (*WheelEvent) ClientY() (clientY int) {
	js.Rewrite("$_.clientY")
	return clientY
}

// CtrlKey prop
// js:"ctrlKey"
func (*WheelEvent) CtrlKey() (ctrlKey bool) {
	js.Rewrite("$_.ctrlKey")
	return ctrlKey
}

// FromElement prop
// js:"fromElement"
func (*WheelEvent) FromElement() (fromElement Element) {
	js.Rewrite("$_.fromElement")
	return fromElement
}

// LayerX prop
// js:"layerX"
func (*WheelEvent) LayerX() (layerX int) {
	js.Rewrite("$_.layerX")
	return layerX
}

// LayerY prop
// js:"layerY"
func (*WheelEvent) LayerY() (layerY int) {
	js.Rewrite("$_.layerY")
	return layerY
}

// MetaKey prop
// js:"metaKey"
func (*WheelEvent) MetaKey() (metaKey bool) {
	js.Rewrite("$_.metaKey")
	return metaKey
}

// MovementX prop
// js:"movementX"
func (*WheelEvent) MovementX() (movementX int) {
	js.Rewrite("$_.movementX")
	return movementX
}

// MovementY prop
// js:"movementY"
func (*WheelEvent) MovementY() (movementY int) {
	js.Rewrite("$_.movementY")
	return movementY
}

// OffsetX prop
// js:"offsetX"
func (*WheelEvent) OffsetX() (offsetX int) {
	js.Rewrite("$_.offsetX")
	return offsetX
}

// OffsetY prop
// js:"offsetY"
func (*WheelEvent) OffsetY() (offsetY int) {
	js.Rewrite("$_.offsetY")
	return offsetY
}

// PageX prop
// js:"pageX"
func (*WheelEvent) PageX() (pageX int) {
	js.Rewrite("$_.pageX")
	return pageX
}

// PageY prop
// js:"pageY"
func (*WheelEvent) PageY() (pageY int) {
	js.Rewrite("$_.pageY")
	return pageY
}

// RelatedTarget prop
// js:"relatedTarget"
func (*WheelEvent) RelatedTarget() (relatedTarget EventTarget) {
	js.Rewrite("$_.relatedTarget")
	return relatedTarget
}

// ScreenX prop
// js:"screenX"
func (*WheelEvent) ScreenX() (screenX int) {
	js.Rewrite("$_.screenX")
	return screenX
}

// ScreenY prop
// js:"screenY"
func (*WheelEvent) ScreenY() (screenY int) {
	js.Rewrite("$_.screenY")
	return screenY
}

// ShiftKey prop
// js:"shiftKey"
func (*WheelEvent) ShiftKey() (shiftKey bool) {
	js.Rewrite("$_.shiftKey")
	return shiftKey
}

// ToElement prop
// js:"toElement"
func (*WheelEvent) ToElement() (toElement Element) {
	js.Rewrite("$_.toElement")
	return toElement
}

// Which prop
// js:"which"
func (*WheelEvent) Which() (which uint8) {
	js.Rewrite("$_.which")
	return which
}

// X prop
// js:"x"
func (*WheelEvent) X() (x int) {
	js.Rewrite("$_.x")
	return x
}

// Y prop
// js:"y"
func (*WheelEvent) Y() (y int) {
	js.Rewrite("$_.y")
	return y
}

// Detail prop
// js:"detail"
func (*WheelEvent) Detail() (detail int) {
	js.Rewrite("$_.detail")
	return detail
}

// View prop
// js:"view"
func (*WheelEvent) View() (view *Window) {
	js.Rewrite("$_.view")
	return view
}

// Bubbles prop
// js:"bubbles"
func (*WheelEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*WheelEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*WheelEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*WheelEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*WheelEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*WheelEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*WheelEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*WheelEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*WheelEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*WheelEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*WheelEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*WheelEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*WheelEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*WheelEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
