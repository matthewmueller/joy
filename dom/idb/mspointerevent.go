package idb

import "github.com/matthewmueller/golly/js"

// NewMSPointerEvent fn
func NewMSPointerEvent(typearg string, eventinitdict *PointerEventInit) *MSPointerEvent {
	js.Rewrite("MSPointerEvent")
	return &MSPointerEvent{}
}

// MSPointerEvent struct
// js:"MSPointerEvent,omit"
type MSPointerEvent struct {
	MouseEvent
}

// GetCurrentPoint
func (*MSPointerEvent) GetCurrentPoint(element Element) {
	js.Rewrite("$<.GetCurrentPoint($1)", element)
}

// GetIntermediatePoints
func (*MSPointerEvent) GetIntermediatePoints(element Element) {
	js.Rewrite("$<.GetIntermediatePoints($1)", element)
}

// InitPointerEvent
func (*MSPointerEvent) InitPointerEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg EventTarget, offsetXArg float32, offsetYArg float32, widthArg int, heightArg int, pressure float32, rotation int, tiltX int, tiltY int, pointerIdArg int, pointerType interface{}, hwTimestampArg uint64, isPrimary bool) {
	js.Rewrite("$<.InitPointerEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, offsetXArg, offsetYArg, widthArg, heightArg, pressure, rotation, tiltX, tiltY, pointerIdArg, pointerType, hwTimestampArg, isPrimary)
}

// CurrentPoint
func (*MSPointerEvent) CurrentPoint() (currentPoint interface{}) {
	js.Rewrite("$<.CurrentPoint")
	return currentPoint
}

// Height
func (*MSPointerEvent) Height() (height int) {
	js.Rewrite("$<.Height")
	return height
}

// HwTimestamp
func (*MSPointerEvent) HwTimestamp() (hwTimestamp uint64) {
	js.Rewrite("$<.HwTimestamp")
	return hwTimestamp
}

// IntermediatePoints
func (*MSPointerEvent) IntermediatePoints() (intermediatePoints interface{}) {
	js.Rewrite("$<.IntermediatePoints")
	return intermediatePoints
}

// IsPrimary
func (*MSPointerEvent) IsPrimary() (isPrimary bool) {
	js.Rewrite("$<.IsPrimary")
	return isPrimary
}

// PointerID
func (*MSPointerEvent) PointerID() (pointerId int) {
	js.Rewrite("$<.PointerID")
	return pointerId
}

// PointerType
func (*MSPointerEvent) PointerType() (pointerType interface{}) {
	js.Rewrite("$<.PointerType")
	return pointerType
}

// Pressure
func (*MSPointerEvent) Pressure() (pressure float32) {
	js.Rewrite("$<.Pressure")
	return pressure
}

// Rotation
func (*MSPointerEvent) Rotation() (rotation int) {
	js.Rewrite("$<.Rotation")
	return rotation
}

// TiltX
func (*MSPointerEvent) TiltX() (tiltX int) {
	js.Rewrite("$<.TiltX")
	return tiltX
}

// TiltY
func (*MSPointerEvent) TiltY() (tiltY int) {
	js.Rewrite("$<.TiltY")
	return tiltY
}

// Width
func (*MSPointerEvent) Width() (width int) {
	js.Rewrite("$<.Width")
	return width
}
