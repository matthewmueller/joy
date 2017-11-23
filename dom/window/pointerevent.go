package window

import "github.com/matthewmueller/golly/js"

// NewPointerEvent fn
func NewPointerEvent(typearg string, eventinitdict *PointerEventInit) *PointerEvent {
	js.Rewrite("PointerEvent")
	return &PointerEvent{}
}

// PointerEvent struct
// js:"PointerEvent,omit"
type PointerEvent struct {
	MouseEvent
}

// GetCurrentPoint fn
func (*PointerEvent) GetCurrentPoint(element Element) {
	js.Rewrite("$<.getCurrentPoint($1)", element)
}

// GetIntermediatePoints fn
func (*PointerEvent) GetIntermediatePoints(element Element) {
	js.Rewrite("$<.getIntermediatePoints($1)", element)
}

// InitPointerEvent fn
func (*PointerEvent) InitPointerEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget, offsetXArg float32, offsetYArg float32, widthArg int, heightArg int, pressure float32, rotation int, tiltX int, tiltY int, pointerIdArg int, pointerType interface{}, hwTimestampArg uint64, isPrimary bool) {
	js.Rewrite("$<.initPointerEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, offsetXArg, offsetYArg, widthArg, heightArg, pressure, rotation, tiltX, tiltY, pointerIdArg, pointerType, hwTimestampArg, isPrimary)
}

// CurrentPoint prop
func (*PointerEvent) CurrentPoint() (currentPoint interface{}) {
	js.Rewrite("$<.currentPoint")
	return currentPoint
}

// Height prop
func (*PointerEvent) Height() (height int) {
	js.Rewrite("$<.height")
	return height
}

// HwTimestamp prop
func (*PointerEvent) HwTimestamp() (hwTimestamp uint64) {
	js.Rewrite("$<.hwTimestamp")
	return hwTimestamp
}

// IntermediatePoints prop
func (*PointerEvent) IntermediatePoints() (intermediatePoints interface{}) {
	js.Rewrite("$<.intermediatePoints")
	return intermediatePoints
}

// IsPrimary prop
func (*PointerEvent) IsPrimary() (isPrimary bool) {
	js.Rewrite("$<.isPrimary")
	return isPrimary
}

// PointerID prop
func (*PointerEvent) PointerID() (pointerId int) {
	js.Rewrite("$<.pointerId")
	return pointerId
}

// PointerType prop
func (*PointerEvent) PointerType() (pointerType interface{}) {
	js.Rewrite("$<.pointerType")
	return pointerType
}

// Pressure prop
func (*PointerEvent) Pressure() (pressure float32) {
	js.Rewrite("$<.pressure")
	return pressure
}

// Rotation prop
func (*PointerEvent) Rotation() (rotation int) {
	js.Rewrite("$<.rotation")
	return rotation
}

// TiltX prop
func (*PointerEvent) TiltX() (tiltX int) {
	js.Rewrite("$<.tiltX")
	return tiltX
}

// TiltY prop
func (*PointerEvent) TiltY() (tiltY int) {
	js.Rewrite("$<.tiltY")
	return tiltY
}

// Width prop
func (*PointerEvent) Width() (width int) {
	js.Rewrite("$<.width")
	return width
}
