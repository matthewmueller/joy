package window

import "github.com/matthewmueller/golly/js"

// TouchEvent struct
// js:"TouchEvent,omit"
type TouchEvent struct {
	UIEvent
}

// AltKey prop
func (*TouchEvent) AltKey() (altKey bool) {
	js.Rewrite("$<.altKey")
	return altKey
}

// ChangedTouches prop
func (*TouchEvent) ChangedTouches() (changedTouches *TouchList) {
	js.Rewrite("$<.changedTouches")
	return changedTouches
}

// CharCode prop
func (*TouchEvent) CharCode() (charCode int8) {
	js.Rewrite("$<.charCode")
	return charCode
}

// CtrlKey prop
func (*TouchEvent) CtrlKey() (ctrlKey bool) {
	js.Rewrite("$<.ctrlKey")
	return ctrlKey
}

// KeyCode prop
func (*TouchEvent) KeyCode() (keyCode int8) {
	js.Rewrite("$<.keyCode")
	return keyCode
}

// MetaKey prop
func (*TouchEvent) MetaKey() (metaKey bool) {
	js.Rewrite("$<.metaKey")
	return metaKey
}

// ShiftKey prop
func (*TouchEvent) ShiftKey() (shiftKey bool) {
	js.Rewrite("$<.shiftKey")
	return shiftKey
}

// TargetTouches prop
func (*TouchEvent) TargetTouches() (targetTouches *TouchList) {
	js.Rewrite("$<.targetTouches")
	return targetTouches
}

// Touches prop
func (*TouchEvent) Touches() (touches *TouchList) {
	js.Rewrite("$<.touches")
	return touches
}

// Which prop
func (*TouchEvent) Which() (which uint8) {
	js.Rewrite("$<.which")
	return which
}
