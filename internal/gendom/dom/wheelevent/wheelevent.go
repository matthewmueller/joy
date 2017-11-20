package wheelevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/mouseevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type WheelEvent struct {
	*mouseevent.MouseEvent
}

func (*WheelEvent) GetCurrentPoint(element *element.Element) {
	js.Rewrite("$<.getCurrentPoint($1)", element)
}

func (*WheelEvent) InitWheelEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, buttonArg uint8, relatedTargetArg *eventtarget.EventTarget, modifiersListArg string, deltaXArg int, deltaYArg int, deltaZArg int, deltaMode uint) {
	js.Rewrite("$<.initWheelEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, buttonArg, relatedTargetArg, modifiersListArg, deltaXArg, deltaYArg, deltaZArg, deltaMode)
}

func (*WheelEvent) GetDeltaMode() (deltaMode uint) {
	js.Rewrite("$<.deltaMode")
	return deltaMode
}

func (*WheelEvent) GetDeltaX() (deltaX int) {
	js.Rewrite("$<.deltaX")
	return deltaX
}

func (*WheelEvent) GetDeltaY() (deltaY int) {
	js.Rewrite("$<.deltaY")
	return deltaY
}

func (*WheelEvent) GetDeltaZ() (deltaZ int) {
	js.Rewrite("$<.deltaZ")
	return deltaZ
}

func (*WheelEvent) GetWheelDelta() (wheelDelta int) {
	js.Rewrite("$<.wheelDelta")
	return wheelDelta
}

func (*WheelEvent) GetWheelDeltaX() (wheelDeltaX int) {
	js.Rewrite("$<.wheelDeltaX")
	return wheelDeltaX
}

func (*WheelEvent) GetWheelDeltaY() (wheelDeltaY int) {
	js.Rewrite("$<.wheelDeltaY")
	return wheelDeltaY
}
