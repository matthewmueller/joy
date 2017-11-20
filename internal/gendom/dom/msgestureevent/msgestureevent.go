package msgestureevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type MSGestureEvent struct {
	*uievent.UIEvent
}

func (*MSGestureEvent) InitGestureEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, offsetXArg float32, offsetYArg float32, translationXArg float32, translationYArg float32, scaleArg float32, expansionArg float32, rotationArg float32, velocityXArg float32, velocityYArg float32, velocityExpansionArg float32, velocityAngularArg float32, hwTimestampArg uint64) {
	js.Rewrite("$<.initGestureEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, offsetXArg, offsetYArg, translationXArg, translationYArg, scaleArg, expansionArg, rotationArg, velocityXArg, velocityYArg, velocityExpansionArg, velocityAngularArg, hwTimestampArg)
}

func (*MSGestureEvent) GetClientX() (clientX float32) {
	js.Rewrite("$<.clientX")
	return clientX
}

func (*MSGestureEvent) GetClientY() (clientY float32) {
	js.Rewrite("$<.clientY")
	return clientY
}

func (*MSGestureEvent) GetExpansion() (expansion float32) {
	js.Rewrite("$<.expansion")
	return expansion
}

func (*MSGestureEvent) GetGestureObject() (gestureObject interface{}) {
	js.Rewrite("$<.gestureObject")
	return gestureObject
}

func (*MSGestureEvent) GetHwTimestamp() (hwTimestamp uint64) {
	js.Rewrite("$<.hwTimestamp")
	return hwTimestamp
}

func (*MSGestureEvent) GetOffsetX() (offsetX float32) {
	js.Rewrite("$<.offsetX")
	return offsetX
}

func (*MSGestureEvent) GetOffsetY() (offsetY float32) {
	js.Rewrite("$<.offsetY")
	return offsetY
}

func (*MSGestureEvent) GetRotation() (rotation float32) {
	js.Rewrite("$<.rotation")
	return rotation
}

func (*MSGestureEvent) GetScale() (scale float32) {
	js.Rewrite("$<.scale")
	return scale
}

func (*MSGestureEvent) GetScreenX() (screenX int) {
	js.Rewrite("$<.screenX")
	return screenX
}

func (*MSGestureEvent) GetScreenY() (screenY int) {
	js.Rewrite("$<.screenY")
	return screenY
}

func (*MSGestureEvent) GetTranslationX() (translationX float32) {
	js.Rewrite("$<.translationX")
	return translationX
}

func (*MSGestureEvent) GetTranslationY() (translationY float32) {
	js.Rewrite("$<.translationY")
	return translationY
}

func (*MSGestureEvent) GetVelocityAngular() (velocityAngular float32) {
	js.Rewrite("$<.velocityAngular")
	return velocityAngular
}

func (*MSGestureEvent) GetVelocityExpansion() (velocityExpansion float32) {
	js.Rewrite("$<.velocityExpansion")
	return velocityExpansion
}

func (*MSGestureEvent) GetVelocityX() (velocityX float32) {
	js.Rewrite("$<.velocityX")
	return velocityX
}

func (*MSGestureEvent) GetVelocityY() (velocityY float32) {
	js.Rewrite("$<.velocityY")
	return velocityY
}
