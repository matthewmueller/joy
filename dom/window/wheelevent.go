package window

import "github.com/matthewmueller/golly/js"

// NewWheelEvent fn
func NewWheelEvent(typearg string, eventinitdict *WheelEventInit) *WheelEvent {
	js.Rewrite("WheelEvent")
	return &WheelEvent{}
}

// WheelEvent struct
// js:"WheelEvent,omit"
type WheelEvent struct {
	MouseEvent
}

// GetCurrentPoint fn
func (*WheelEvent) GetCurrentPoint(element Element) {
	js.Rewrite("$<.getCurrentPoint($1)", element)
}

// InitWheelEvent fn
func (*WheelEvent) InitWheelEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, buttonArg uint8, relatedTargetArg EventTarget, modifiersListArg string, deltaXArg int, deltaYArg int, deltaZArg int, deltaMode uint) {
	js.Rewrite("$<.initWheelEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, buttonArg, relatedTargetArg, modifiersListArg, deltaXArg, deltaYArg, deltaZArg, deltaMode)
}

// DeltaMode prop
func (*WheelEvent) DeltaMode() (deltaMode uint) {
	js.Rewrite("$<.deltaMode")
	return deltaMode
}

// DeltaX prop
func (*WheelEvent) DeltaX() (deltaX int) {
	js.Rewrite("$<.deltaX")
	return deltaX
}

// DeltaY prop
func (*WheelEvent) DeltaY() (deltaY int) {
	js.Rewrite("$<.deltaY")
	return deltaY
}

// DeltaZ prop
func (*WheelEvent) DeltaZ() (deltaZ int) {
	js.Rewrite("$<.deltaZ")
	return deltaZ
}

// WheelDelta prop
func (*WheelEvent) WheelDelta() (wheelDelta int) {
	js.Rewrite("$<.wheelDelta")
	return wheelDelta
}

// WheelDeltaX prop
func (*WheelEvent) WheelDeltaX() (wheelDeltaX int) {
	js.Rewrite("$<.wheelDeltaX")
	return wheelDeltaX
}

// WheelDeltaY prop
func (*WheelEvent) WheelDeltaY() (wheelDeltaY int) {
	js.Rewrite("$<.wheelDeltaY")
	return wheelDeltaY
}
