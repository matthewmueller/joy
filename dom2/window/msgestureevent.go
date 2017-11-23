package window

import "github.com/matthewmueller/golly/js"

// MSGestureEvent struct
// js:"MSGestureEvent,omit"
type MSGestureEvent struct {
	UIEvent
}

// InitGestureEvent
func (*MSGestureEvent) InitGestureEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, offsetXArg float32, offsetYArg float32, translationXArg float32, translationYArg float32, scaleArg float32, expansionArg float32, rotationArg float32, velocityXArg float32, velocityYArg float32, velocityExpansionArg float32, velocityAngularArg float32, hwTimestampArg uint64) {
	js.Rewrite("$<.InitGestureEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, offsetXArg, offsetYArg, translationXArg, translationYArg, scaleArg, expansionArg, rotationArg, velocityXArg, velocityYArg, velocityExpansionArg, velocityAngularArg, hwTimestampArg)
}

// ClientX
func (*MSGestureEvent) ClientX() (clientX float32) {
	js.Rewrite("$<.ClientX")
	return clientX
}

// ClientY
func (*MSGestureEvent) ClientY() (clientY float32) {
	js.Rewrite("$<.ClientY")
	return clientY
}

// Expansion
func (*MSGestureEvent) Expansion() (expansion float32) {
	js.Rewrite("$<.Expansion")
	return expansion
}

// GestureObject
func (*MSGestureEvent) GestureObject() (gestureObject interface{}) {
	js.Rewrite("$<.GestureObject")
	return gestureObject
}

// HwTimestamp
func (*MSGestureEvent) HwTimestamp() (hwTimestamp uint64) {
	js.Rewrite("$<.HwTimestamp")
	return hwTimestamp
}

// OffsetX
func (*MSGestureEvent) OffsetX() (offsetX float32) {
	js.Rewrite("$<.OffsetX")
	return offsetX
}

// OffsetY
func (*MSGestureEvent) OffsetY() (offsetY float32) {
	js.Rewrite("$<.OffsetY")
	return offsetY
}

// Rotation
func (*MSGestureEvent) Rotation() (rotation float32) {
	js.Rewrite("$<.Rotation")
	return rotation
}

// Scale
func (*MSGestureEvent) Scale() (scale float32) {
	js.Rewrite("$<.Scale")
	return scale
}

// ScreenX
func (*MSGestureEvent) ScreenX() (screenX int) {
	js.Rewrite("$<.ScreenX")
	return screenX
}

// ScreenY
func (*MSGestureEvent) ScreenY() (screenY int) {
	js.Rewrite("$<.ScreenY")
	return screenY
}

// TranslationX
func (*MSGestureEvent) TranslationX() (translationX float32) {
	js.Rewrite("$<.TranslationX")
	return translationX
}

// TranslationY
func (*MSGestureEvent) TranslationY() (translationY float32) {
	js.Rewrite("$<.TranslationY")
	return translationY
}

// VelocityAngular
func (*MSGestureEvent) VelocityAngular() (velocityAngular float32) {
	js.Rewrite("$<.VelocityAngular")
	return velocityAngular
}

// VelocityExpansion
func (*MSGestureEvent) VelocityExpansion() (velocityExpansion float32) {
	js.Rewrite("$<.VelocityExpansion")
	return velocityExpansion
}

// VelocityX
func (*MSGestureEvent) VelocityX() (velocityX float32) {
	js.Rewrite("$<.VelocityX")
	return velocityX
}

// VelocityY
func (*MSGestureEvent) VelocityY() (velocityY float32) {
	js.Rewrite("$<.VelocityY")
	return velocityY
}
