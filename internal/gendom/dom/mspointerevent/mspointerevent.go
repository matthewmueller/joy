package mspointerevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/mouseevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type MSPointerEvent struct {
	*mouseevent.MouseEvent
}

func (*MSPointerEvent) GetCurrentPoint(element *element.Element) {
	js.Rewrite("$<.getCurrentPoint($1)", element)
}

func (*MSPointerEvent) GetIntermediatePoints(element *element.Element) {
	js.Rewrite("$<.getIntermediatePoints($1)", element)
}

func (*MSPointerEvent) InitPointerEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg *eventtarget.EventTarget, offsetXArg float32, offsetYArg float32, widthArg int, heightArg int, pressure float32, rotation int, tiltX int, tiltY int, pointerIdArg int, pointerType interface{}, hwTimestampArg uint64, isPrimary bool) {
	js.Rewrite("$<.initPointerEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, offsetXArg, offsetYArg, widthArg, heightArg, pressure, rotation, tiltX, tiltY, pointerIdArg, pointerType, hwTimestampArg, isPrimary)
}

func (*MSPointerEvent) GetHeight() (height int) {
	js.Rewrite("$<.height")
	return height
}

func (*MSPointerEvent) GetHwTimestamp() (hwTimestamp uint64) {
	js.Rewrite("$<.hwTimestamp")
	return hwTimestamp
}

func (*MSPointerEvent) GetIsPrimary() (isPrimary bool) {
	js.Rewrite("$<.isPrimary")
	return isPrimary
}

func (*MSPointerEvent) GetPointerID() (pointerId int) {
	js.Rewrite("$<.pointerId")
	return pointerId
}

func (*MSPointerEvent) GetPointerType() (pointerType interface{}) {
	js.Rewrite("$<.pointerType")
	return pointerType
}

func (*MSPointerEvent) GetPressure() (pressure float32) {
	js.Rewrite("$<.pressure")
	return pressure
}

func (*MSPointerEvent) GetRotation() (rotation int) {
	js.Rewrite("$<.rotation")
	return rotation
}

func (*MSPointerEvent) GetTiltX() (tiltX int) {
	js.Rewrite("$<.tiltX")
	return tiltX
}

func (*MSPointerEvent) GetTiltY() (tiltY int) {
	js.Rewrite("$<.tiltY")
	return tiltY
}

func (*MSPointerEvent) GetWidth() (width int) {
	js.Rewrite("$<.width")
	return width
}
