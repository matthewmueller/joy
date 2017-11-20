package pointerevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/mouseevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type PointerEvent struct {
	*mouseevent.MouseEvent
}

func (*PointerEvent) GetCurrentPoint(element *element.Element) {
	js.Rewrite("$<.getCurrentPoint($1)", element)
}

func (*PointerEvent) GetIntermediatePoints(element *element.Element) {
	js.Rewrite("$<.getIntermediatePoints($1)", element)
}

func (*PointerEvent) InitPointerEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg *eventtarget.EventTarget, offsetXArg float32, offsetYArg float32, widthArg int, heightArg int, pressure float32, rotation int, tiltX int, tiltY int, pointerIdArg int, pointerType interface{}, hwTimestampArg uint64, isPrimary bool) {
	js.Rewrite("$<.initPointerEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, offsetXArg, offsetYArg, widthArg, heightArg, pressure, rotation, tiltX, tiltY, pointerIdArg, pointerType, hwTimestampArg, isPrimary)
}

func (*PointerEvent) GetHeight() (height int) {
	js.Rewrite("$<.height")
	return height
}

func (*PointerEvent) GetHwTimestamp() (hwTimestamp uint64) {
	js.Rewrite("$<.hwTimestamp")
	return hwTimestamp
}

func (*PointerEvent) GetIsPrimary() (isPrimary bool) {
	js.Rewrite("$<.isPrimary")
	return isPrimary
}

func (*PointerEvent) GetPointerID() (pointerId int) {
	js.Rewrite("$<.pointerId")
	return pointerId
}

func (*PointerEvent) GetPointerType() (pointerType interface{}) {
	js.Rewrite("$<.pointerType")
	return pointerType
}

func (*PointerEvent) GetPressure() (pressure float32) {
	js.Rewrite("$<.pressure")
	return pressure
}

func (*PointerEvent) GetRotation() (rotation int) {
	js.Rewrite("$<.rotation")
	return rotation
}

func (*PointerEvent) GetTiltX() (tiltX int) {
	js.Rewrite("$<.tiltX")
	return tiltX
}

func (*PointerEvent) GetTiltY() (tiltY int) {
	js.Rewrite("$<.tiltY")
	return tiltY
}

func (*PointerEvent) GetWidth() (width int) {
	js.Rewrite("$<.width")
	return width
}
