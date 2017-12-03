package window

import "github.com/matthewmueller/joy/macro"

var _ MouseEvent = (*MSPointerEvent)(nil)
var _ UIEvent = (*MSPointerEvent)(nil)
var _ Event = (*MSPointerEvent)(nil)

// NewMSPointerEvent fn
func NewMSPointerEvent(typearg string, eventinitdict *PointerEventInit) *MSPointerEvent {
	macro.Rewrite("MSPointerEvent")
	return &MSPointerEvent{}
}

// MSPointerEvent struct
// js:"MSPointerEvent,omit"
type MSPointerEvent struct {
}

// GetCurrentPoint fn
// js:"getCurrentPoint"
func (*MSPointerEvent) GetCurrentPoint(element Element) {
	macro.Rewrite("$_.getCurrentPoint($1)", element)
}

// GetIntermediatePoints fn
// js:"getIntermediatePoints"
func (*MSPointerEvent) GetIntermediatePoints(element Element) {
	macro.Rewrite("$_.getIntermediatePoints($1)", element)
}

// InitPointerEvent fn
// js:"initPointerEvent"
func (*MSPointerEvent) InitPointerEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget, offsetXArg float32, offsetYArg float32, widthArg int, heightArg int, pressure float32, rotation int, tiltX int, tiltY int, pointerIdArg int, pointerType interface{}, hwTimestampArg uint64, isPrimary bool) {
	macro.Rewrite("$_.initPointerEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, offsetXArg, offsetYArg, widthArg, heightArg, pressure, rotation, tiltX, tiltY, pointerIdArg, pointerType, hwTimestampArg, isPrimary)
}

// GetModifierState fn
// js:"getModifierState"
func (*MSPointerEvent) GetModifierState(keyArg string) (b bool) {
	macro.Rewrite("$_.getModifierState($1)", keyArg)
	return b
}

// InitMouseEvent fn
// js:"initMouseEvent"
func (*MSPointerEvent) InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget) {
	macro.Rewrite("$_.initMouseEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg)
}

// InitUIEvent fn
// js:"initUIEvent"
func (*MSPointerEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	macro.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

// InitEvent fn
// js:"initEvent"
func (*MSPointerEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MSPointerEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MSPointerEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MSPointerEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// CurrentPoint prop
// js:"currentPoint"
func (*MSPointerEvent) CurrentPoint() (currentPoint interface{}) {
	macro.Rewrite("$_.currentPoint")
	return currentPoint
}

// Height prop
// js:"height"
func (*MSPointerEvent) Height() (height int) {
	macro.Rewrite("$_.height")
	return height
}

// HwTimestamp prop
// js:"hwTimestamp"
func (*MSPointerEvent) HwTimestamp() (hwTimestamp uint64) {
	macro.Rewrite("$_.hwTimestamp")
	return hwTimestamp
}

// IntermediatePoints prop
// js:"intermediatePoints"
func (*MSPointerEvent) IntermediatePoints() (intermediatePoints interface{}) {
	macro.Rewrite("$_.intermediatePoints")
	return intermediatePoints
}

// IsPrimary prop
// js:"isPrimary"
func (*MSPointerEvent) IsPrimary() (isPrimary bool) {
	macro.Rewrite("$_.isPrimary")
	return isPrimary
}

// PointerID prop
// js:"pointerId"
func (*MSPointerEvent) PointerID() (pointerId int) {
	macro.Rewrite("$_.pointerId")
	return pointerId
}

// PointerType prop
// js:"pointerType"
func (*MSPointerEvent) PointerType() (pointerType interface{}) {
	macro.Rewrite("$_.pointerType")
	return pointerType
}

// Pressure prop
// js:"pressure"
func (*MSPointerEvent) Pressure() (pressure float32) {
	macro.Rewrite("$_.pressure")
	return pressure
}

// Rotation prop
// js:"rotation"
func (*MSPointerEvent) Rotation() (rotation int) {
	macro.Rewrite("$_.rotation")
	return rotation
}

// TiltX prop
// js:"tiltX"
func (*MSPointerEvent) TiltX() (tiltX int) {
	macro.Rewrite("$_.tiltX")
	return tiltX
}

// TiltY prop
// js:"tiltY"
func (*MSPointerEvent) TiltY() (tiltY int) {
	macro.Rewrite("$_.tiltY")
	return tiltY
}

// Width prop
// js:"width"
func (*MSPointerEvent) Width() (width int) {
	macro.Rewrite("$_.width")
	return width
}

// AltKey prop
// js:"altKey"
func (*MSPointerEvent) AltKey() (altKey bool) {
	macro.Rewrite("$_.altKey")
	return altKey
}

// Button prop
// js:"button"
func (*MSPointerEvent) Button() (button uint8) {
	macro.Rewrite("$_.button")
	return button
}

// Buttons prop
// js:"buttons"
func (*MSPointerEvent) Buttons() (buttons uint8) {
	macro.Rewrite("$_.buttons")
	return buttons
}

// ClientX prop
// js:"clientX"
func (*MSPointerEvent) ClientX() (clientX int) {
	macro.Rewrite("$_.clientX")
	return clientX
}

// ClientY prop
// js:"clientY"
func (*MSPointerEvent) ClientY() (clientY int) {
	macro.Rewrite("$_.clientY")
	return clientY
}

// CtrlKey prop
// js:"ctrlKey"
func (*MSPointerEvent) CtrlKey() (ctrlKey bool) {
	macro.Rewrite("$_.ctrlKey")
	return ctrlKey
}

// FromElement prop
// js:"fromElement"
func (*MSPointerEvent) FromElement() (fromElement Element) {
	macro.Rewrite("$_.fromElement")
	return fromElement
}

// LayerX prop
// js:"layerX"
func (*MSPointerEvent) LayerX() (layerX int) {
	macro.Rewrite("$_.layerX")
	return layerX
}

// LayerY prop
// js:"layerY"
func (*MSPointerEvent) LayerY() (layerY int) {
	macro.Rewrite("$_.layerY")
	return layerY
}

// MetaKey prop
// js:"metaKey"
func (*MSPointerEvent) MetaKey() (metaKey bool) {
	macro.Rewrite("$_.metaKey")
	return metaKey
}

// MovementX prop
// js:"movementX"
func (*MSPointerEvent) MovementX() (movementX int) {
	macro.Rewrite("$_.movementX")
	return movementX
}

// MovementY prop
// js:"movementY"
func (*MSPointerEvent) MovementY() (movementY int) {
	macro.Rewrite("$_.movementY")
	return movementY
}

// OffsetX prop
// js:"offsetX"
func (*MSPointerEvent) OffsetX() (offsetX int) {
	macro.Rewrite("$_.offsetX")
	return offsetX
}

// OffsetY prop
// js:"offsetY"
func (*MSPointerEvent) OffsetY() (offsetY int) {
	macro.Rewrite("$_.offsetY")
	return offsetY
}

// PageX prop
// js:"pageX"
func (*MSPointerEvent) PageX() (pageX int) {
	macro.Rewrite("$_.pageX")
	return pageX
}

// PageY prop
// js:"pageY"
func (*MSPointerEvent) PageY() (pageY int) {
	macro.Rewrite("$_.pageY")
	return pageY
}

// RelatedTarget prop
// js:"relatedTarget"
func (*MSPointerEvent) RelatedTarget() (relatedTarget EventTarget) {
	macro.Rewrite("$_.relatedTarget")
	return relatedTarget
}

// ScreenX prop
// js:"screenX"
func (*MSPointerEvent) ScreenX() (screenX int) {
	macro.Rewrite("$_.screenX")
	return screenX
}

// ScreenY prop
// js:"screenY"
func (*MSPointerEvent) ScreenY() (screenY int) {
	macro.Rewrite("$_.screenY")
	return screenY
}

// ShiftKey prop
// js:"shiftKey"
func (*MSPointerEvent) ShiftKey() (shiftKey bool) {
	macro.Rewrite("$_.shiftKey")
	return shiftKey
}

// ToElement prop
// js:"toElement"
func (*MSPointerEvent) ToElement() (toElement Element) {
	macro.Rewrite("$_.toElement")
	return toElement
}

// Which prop
// js:"which"
func (*MSPointerEvent) Which() (which uint8) {
	macro.Rewrite("$_.which")
	return which
}

// X prop
// js:"x"
func (*MSPointerEvent) X() (x int) {
	macro.Rewrite("$_.x")
	return x
}

// Y prop
// js:"y"
func (*MSPointerEvent) Y() (y int) {
	macro.Rewrite("$_.y")
	return y
}

// Detail prop
// js:"detail"
func (*MSPointerEvent) Detail() (detail int) {
	macro.Rewrite("$_.detail")
	return detail
}

// View prop
// js:"view"
func (*MSPointerEvent) View() (view *Window) {
	macro.Rewrite("$_.view")
	return view
}

// Bubbles prop
// js:"bubbles"
func (*MSPointerEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MSPointerEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MSPointerEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MSPointerEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MSPointerEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MSPointerEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MSPointerEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MSPointerEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MSPointerEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MSPointerEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MSPointerEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MSPointerEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MSPointerEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MSPointerEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
