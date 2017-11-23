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

// GetCurrentPoint
func (*WheelEvent) GetCurrentPoint(element Element) {
	js.Rewrite("$<.GetCurrentPoint($1)", element)
}

// InitWheelEvent
func (*WheelEvent) InitWheelEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, buttonArg uint8, relatedTargetArg EventTarget, modifiersListArg string, deltaXArg int, deltaYArg int, deltaZArg int, deltaMode uint) {
	js.Rewrite("$<.InitWheelEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, buttonArg, relatedTargetArg, modifiersListArg, deltaXArg, deltaYArg, deltaZArg, deltaMode)
}

// DeltaMode
func (*WheelEvent) DeltaMode() (deltaMode uint) {
	js.Rewrite("$<.DeltaMode")
	return deltaMode
}

// DeltaX
func (*WheelEvent) DeltaX() (deltaX int) {
	js.Rewrite("$<.DeltaX")
	return deltaX
}

// DeltaY
func (*WheelEvent) DeltaY() (deltaY int) {
	js.Rewrite("$<.DeltaY")
	return deltaY
}

// DeltaZ
func (*WheelEvent) DeltaZ() (deltaZ int) {
	js.Rewrite("$<.DeltaZ")
	return deltaZ
}

// WheelDelta
func (*WheelEvent) WheelDelta() (wheelDelta int) {
	js.Rewrite("$<.WheelDelta")
	return wheelDelta
}

// WheelDeltaX
func (*WheelEvent) WheelDeltaX() (wheelDeltaX int) {
	js.Rewrite("$<.WheelDeltaX")
	return wheelDeltaX
}

// WheelDeltaY
func (*WheelEvent) WheelDeltaY() (wheelDeltaY int) {
	js.Rewrite("$<.WheelDeltaY")
	return wheelDeltaY
}
