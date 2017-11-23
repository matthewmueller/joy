package window

import "github.com/matthewmueller/golly/js"

// MSGestureEvent struct
// js:"MSGestureEvent,omit"
type MSGestureEvent struct {
	UIEvent
}

// InitGestureEvent fn
func (*MSGestureEvent) InitGestureEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, offsetXArg float32, offsetYArg float32, translationXArg float32, translationYArg float32, scaleArg float32, expansionArg float32, rotationArg float32, velocityXArg float32, velocityYArg float32, velocityExpansionArg float32, velocityAngularArg float32, hwTimestampArg uint64) {
	js.Rewrite("$<.initGestureEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, offsetXArg, offsetYArg, translationXArg, translationYArg, scaleArg, expansionArg, rotationArg, velocityXArg, velocityYArg, velocityExpansionArg, velocityAngularArg, hwTimestampArg)
}

// ClientX prop
func (*MSGestureEvent) ClientX() (clientX float32) {
	js.Rewrite("$<.clientX")
	return clientX
}

// ClientY prop
func (*MSGestureEvent) ClientY() (clientY float32) {
	js.Rewrite("$<.clientY")
	return clientY
}

// Expansion prop
func (*MSGestureEvent) Expansion() (expansion float32) {
	js.Rewrite("$<.expansion")
	return expansion
}

// GestureObject prop
func (*MSGestureEvent) GestureObject() (gestureObject interface{}) {
	js.Rewrite("$<.gestureObject")
	return gestureObject
}

// HwTimestamp prop
func (*MSGestureEvent) HwTimestamp() (hwTimestamp uint64) {
	js.Rewrite("$<.hwTimestamp")
	return hwTimestamp
}

// OffsetX prop
func (*MSGestureEvent) OffsetX() (offsetX float32) {
	js.Rewrite("$<.offsetX")
	return offsetX
}

// OffsetY prop
func (*MSGestureEvent) OffsetY() (offsetY float32) {
	js.Rewrite("$<.offsetY")
	return offsetY
}

// Rotation prop
func (*MSGestureEvent) Rotation() (rotation float32) {
	js.Rewrite("$<.rotation")
	return rotation
}

// Scale prop
func (*MSGestureEvent) Scale() (scale float32) {
	js.Rewrite("$<.scale")
	return scale
}

// ScreenX prop
func (*MSGestureEvent) ScreenX() (screenX int) {
	js.Rewrite("$<.screenX")
	return screenX
}

// ScreenY prop
func (*MSGestureEvent) ScreenY() (screenY int) {
	js.Rewrite("$<.screenY")
	return screenY
}

// TranslationX prop
func (*MSGestureEvent) TranslationX() (translationX float32) {
	js.Rewrite("$<.translationX")
	return translationX
}

// TranslationY prop
func (*MSGestureEvent) TranslationY() (translationY float32) {
	js.Rewrite("$<.translationY")
	return translationY
}

// VelocityAngular prop
func (*MSGestureEvent) VelocityAngular() (velocityAngular float32) {
	js.Rewrite("$<.velocityAngular")
	return velocityAngular
}

// VelocityExpansion prop
func (*MSGestureEvent) VelocityExpansion() (velocityExpansion float32) {
	js.Rewrite("$<.velocityExpansion")
	return velocityExpansion
}

// VelocityX prop
func (*MSGestureEvent) VelocityX() (velocityX float32) {
	js.Rewrite("$<.velocityX")
	return velocityX
}

// VelocityY prop
func (*MSGestureEvent) VelocityY() (velocityY float32) {
	js.Rewrite("$<.velocityY")
	return velocityY
}
