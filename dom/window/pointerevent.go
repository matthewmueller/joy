package window

import "github.com/matthewmueller/golly/js"

var _ MouseEvent = (*PointerEvent)(nil)
var _ UIEvent = (*PointerEvent)(nil)
var _ Event = (*PointerEvent)(nil)

// PointerEvent struct
// js:"PointerEvent,omit"
type PointerEvent struct {
}

// GetCurrentPoint fn
// js:"getCurrentPoint"
func (*PointerEvent) GetCurrentPoint(element Element) {
	js.Rewrite("$_.getCurrentPoint($1)", element)
}

// GetIntermediatePoints fn
// js:"getIntermediatePoints"
func (*PointerEvent) GetIntermediatePoints(element Element) {
	js.Rewrite("$_.getIntermediatePoints($1)", element)
}

// InitPointerEvent fn
// js:"initPointerEvent"
func (*PointerEvent) InitPointerEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget, offsetXArg float32, offsetYArg float32, widthArg int, heightArg int, pressure float32, rotation int, tiltX int, tiltY int, pointerIdArg int, pointerType interface{}, hwTimestampArg uint64, isPrimary bool) {
	js.Rewrite("$_.initPointerEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, offsetXArg, offsetYArg, widthArg, heightArg, pressure, rotation, tiltX, tiltY, pointerIdArg, pointerType, hwTimestampArg, isPrimary)
}

// GetModifierState fn
// js:"getModifierState"
func (*PointerEvent) GetModifierState(keyArg string) (b bool) {
	js.Rewrite("$_.getModifierState($1)", keyArg)
	return b
}

// InitMouseEvent fn
// js:"initMouseEvent"
func (*PointerEvent) InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget) {
	js.Rewrite("$_.initMouseEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg)
}

// InitUIEvent fn
// js:"initUIEvent"
func (*PointerEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	js.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

// InitEvent fn
// js:"initEvent"
func (*PointerEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*PointerEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*PointerEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*PointerEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// CurrentPoint prop
// js:"currentPoint"
func (*PointerEvent) CurrentPoint() (currentPoint interface{}) {
	js.Rewrite("$_.currentPoint")
	return currentPoint
}

// Height prop
// js:"height"
func (*PointerEvent) Height() (height int) {
	js.Rewrite("$_.height")
	return height
}

// HwTimestamp prop
// js:"hwTimestamp"
func (*PointerEvent) HwTimestamp() (hwTimestamp uint64) {
	js.Rewrite("$_.hwTimestamp")
	return hwTimestamp
}

// IntermediatePoints prop
// js:"intermediatePoints"
func (*PointerEvent) IntermediatePoints() (intermediatePoints interface{}) {
	js.Rewrite("$_.intermediatePoints")
	return intermediatePoints
}

// IsPrimary prop
// js:"isPrimary"
func (*PointerEvent) IsPrimary() (isPrimary bool) {
	js.Rewrite("$_.isPrimary")
	return isPrimary
}

// PointerID prop
// js:"pointerId"
func (*PointerEvent) PointerID() (pointerId int) {
	js.Rewrite("$_.pointerId")
	return pointerId
}

// PointerType prop
// js:"pointerType"
func (*PointerEvent) PointerType() (pointerType interface{}) {
	js.Rewrite("$_.pointerType")
	return pointerType
}

// Pressure prop
// js:"pressure"
func (*PointerEvent) Pressure() (pressure float32) {
	js.Rewrite("$_.pressure")
	return pressure
}

// Rotation prop
// js:"rotation"
func (*PointerEvent) Rotation() (rotation int) {
	js.Rewrite("$_.rotation")
	return rotation
}

// TiltX prop
// js:"tiltX"
func (*PointerEvent) TiltX() (tiltX int) {
	js.Rewrite("$_.tiltX")
	return tiltX
}

// TiltY prop
// js:"tiltY"
func (*PointerEvent) TiltY() (tiltY int) {
	js.Rewrite("$_.tiltY")
	return tiltY
}

// Width prop
// js:"width"
func (*PointerEvent) Width() (width int) {
	js.Rewrite("$_.width")
	return width
}

// AltKey prop
// js:"altKey"
func (*PointerEvent) AltKey() (altKey bool) {
	js.Rewrite("$_.altKey")
	return altKey
}

// Button prop
// js:"button"
func (*PointerEvent) Button() (button uint8) {
	js.Rewrite("$_.button")
	return button
}

// Buttons prop
// js:"buttons"
func (*PointerEvent) Buttons() (buttons uint8) {
	js.Rewrite("$_.buttons")
	return buttons
}

// ClientX prop
// js:"clientX"
func (*PointerEvent) ClientX() (clientX int) {
	js.Rewrite("$_.clientX")
	return clientX
}

// ClientY prop
// js:"clientY"
func (*PointerEvent) ClientY() (clientY int) {
	js.Rewrite("$_.clientY")
	return clientY
}

// CtrlKey prop
// js:"ctrlKey"
func (*PointerEvent) CtrlKey() (ctrlKey bool) {
	js.Rewrite("$_.ctrlKey")
	return ctrlKey
}

// FromElement prop
// js:"fromElement"
func (*PointerEvent) FromElement() (fromElement Element) {
	js.Rewrite("$_.fromElement")
	return fromElement
}

// LayerX prop
// js:"layerX"
func (*PointerEvent) LayerX() (layerX int) {
	js.Rewrite("$_.layerX")
	return layerX
}

// LayerY prop
// js:"layerY"
func (*PointerEvent) LayerY() (layerY int) {
	js.Rewrite("$_.layerY")
	return layerY
}

// MetaKey prop
// js:"metaKey"
func (*PointerEvent) MetaKey() (metaKey bool) {
	js.Rewrite("$_.metaKey")
	return metaKey
}

// MovementX prop
// js:"movementX"
func (*PointerEvent) MovementX() (movementX int) {
	js.Rewrite("$_.movementX")
	return movementX
}

// MovementY prop
// js:"movementY"
func (*PointerEvent) MovementY() (movementY int) {
	js.Rewrite("$_.movementY")
	return movementY
}

// OffsetX prop
// js:"offsetX"
func (*PointerEvent) OffsetX() (offsetX int) {
	js.Rewrite("$_.offsetX")
	return offsetX
}

// OffsetY prop
// js:"offsetY"
func (*PointerEvent) OffsetY() (offsetY int) {
	js.Rewrite("$_.offsetY")
	return offsetY
}

// PageX prop
// js:"pageX"
func (*PointerEvent) PageX() (pageX int) {
	js.Rewrite("$_.pageX")
	return pageX
}

// PageY prop
// js:"pageY"
func (*PointerEvent) PageY() (pageY int) {
	js.Rewrite("$_.pageY")
	return pageY
}

// RelatedTarget prop
// js:"relatedTarget"
func (*PointerEvent) RelatedTarget() (relatedTarget EventTarget) {
	js.Rewrite("$_.relatedTarget")
	return relatedTarget
}

// ScreenX prop
// js:"screenX"
func (*PointerEvent) ScreenX() (screenX int) {
	js.Rewrite("$_.screenX")
	return screenX
}

// ScreenY prop
// js:"screenY"
func (*PointerEvent) ScreenY() (screenY int) {
	js.Rewrite("$_.screenY")
	return screenY
}

// ShiftKey prop
// js:"shiftKey"
func (*PointerEvent) ShiftKey() (shiftKey bool) {
	js.Rewrite("$_.shiftKey")
	return shiftKey
}

// ToElement prop
// js:"toElement"
func (*PointerEvent) ToElement() (toElement Element) {
	js.Rewrite("$_.toElement")
	return toElement
}

// Which prop
// js:"which"
func (*PointerEvent) Which() (which uint8) {
	js.Rewrite("$_.which")
	return which
}

// X prop
// js:"x"
func (*PointerEvent) X() (x int) {
	js.Rewrite("$_.x")
	return x
}

// Y prop
// js:"y"
func (*PointerEvent) Y() (y int) {
	js.Rewrite("$_.y")
	return y
}

// Detail prop
// js:"detail"
func (*PointerEvent) Detail() (detail int) {
	js.Rewrite("$_.detail")
	return detail
}

// View prop
// js:"view"
func (*PointerEvent) View() (view *Window) {
	js.Rewrite("$_.view")
	return view
}

// Bubbles prop
// js:"bubbles"
func (*PointerEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*PointerEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*PointerEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*PointerEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*PointerEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*PointerEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*PointerEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*PointerEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*PointerEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*PointerEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*PointerEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*PointerEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*PointerEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*PointerEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
