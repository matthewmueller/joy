package mouseevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type MouseEvent struct {
	*uievent.UIEvent
}

func (*MouseEvent) GetModifierState(keyArg string) (b bool) {
	js.Rewrite("$<.getModifierState($1)", keyArg)
	return b
}

func (*MouseEvent) InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg *eventtarget.EventTarget) {
	js.Rewrite("$<.initMouseEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg)
}

func (*MouseEvent) GetAltKey() (altKey bool) {
	js.Rewrite("$<.altKey")
	return altKey
}

func (*MouseEvent) GetButton() (button uint8) {
	js.Rewrite("$<.button")
	return button
}

func (*MouseEvent) GetButtons() (buttons uint8) {
	js.Rewrite("$<.buttons")
	return buttons
}

func (*MouseEvent) GetClientX() (clientX int) {
	js.Rewrite("$<.clientX")
	return clientX
}

func (*MouseEvent) GetClientY() (clientY int) {
	js.Rewrite("$<.clientY")
	return clientY
}

func (*MouseEvent) GetCtrlKey() (ctrlKey bool) {
	js.Rewrite("$<.ctrlKey")
	return ctrlKey
}

func (*MouseEvent) GetFromElement() (fromElement *element.Element) {
	js.Rewrite("$<.fromElement")
	return fromElement
}

func (*MouseEvent) GetLayerX() (layerX int) {
	js.Rewrite("$<.layerX")
	return layerX
}

func (*MouseEvent) GetLayerY() (layerY int) {
	js.Rewrite("$<.layerY")
	return layerY
}

func (*MouseEvent) GetMetaKey() (metaKey bool) {
	js.Rewrite("$<.metaKey")
	return metaKey
}

func (*MouseEvent) GetMovementX() (movementX int) {
	js.Rewrite("$<.movementX")
	return movementX
}

func (*MouseEvent) GetMovementY() (movementY int) {
	js.Rewrite("$<.movementY")
	return movementY
}

func (*MouseEvent) GetOffsetX() (offsetX int) {
	js.Rewrite("$<.offsetX")
	return offsetX
}

func (*MouseEvent) GetOffsetY() (offsetY int) {
	js.Rewrite("$<.offsetY")
	return offsetY
}

func (*MouseEvent) GetPageX() (pageX int) {
	js.Rewrite("$<.pageX")
	return pageX
}

func (*MouseEvent) GetPageY() (pageY int) {
	js.Rewrite("$<.pageY")
	return pageY
}

func (*MouseEvent) GetRelatedTarget() (relatedTarget *eventtarget.EventTarget) {
	js.Rewrite("$<.relatedTarget")
	return relatedTarget
}

func (*MouseEvent) GetScreenX() (screenX int) {
	js.Rewrite("$<.screenX")
	return screenX
}

func (*MouseEvent) GetScreenY() (screenY int) {
	js.Rewrite("$<.screenY")
	return screenY
}

func (*MouseEvent) GetShiftKey() (shiftKey bool) {
	js.Rewrite("$<.shiftKey")
	return shiftKey
}

func (*MouseEvent) GetToElement() (toElement *element.Element) {
	js.Rewrite("$<.toElement")
	return toElement
}

func (*MouseEvent) GetWhich() (which uint8) {
	js.Rewrite("$<.which")
	return which
}

func (*MouseEvent) GetX() (x int) {
	js.Rewrite("$<.x")
	return x
}

func (*MouseEvent) GetY() (y int) {
	js.Rewrite("$<.y")
	return y
}
