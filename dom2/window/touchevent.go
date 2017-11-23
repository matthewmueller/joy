package window

import "github.com/matthewmueller/golly/js"

// js:"TouchEvent,omit"
type TouchEvent struct {
	UIEvent
}

// AltKey
func (*TouchEvent) AltKey() (altKey bool) {
	js.Rewrite("$<.AltKey")
	return altKey
}

// ChangedTouches
func (*TouchEvent) ChangedTouches() (changedTouches *TouchList) {
	js.Rewrite("$<.ChangedTouches")
	return changedTouches
}

// CharCode
func (*TouchEvent) CharCode() (charCode int8) {
	js.Rewrite("$<.CharCode")
	return charCode
}

// CtrlKey
func (*TouchEvent) CtrlKey() (ctrlKey bool) {
	js.Rewrite("$<.CtrlKey")
	return ctrlKey
}

// KeyCode
func (*TouchEvent) KeyCode() (keyCode int8) {
	js.Rewrite("$<.KeyCode")
	return keyCode
}

// MetaKey
func (*TouchEvent) MetaKey() (metaKey bool) {
	js.Rewrite("$<.MetaKey")
	return metaKey
}

// ShiftKey
func (*TouchEvent) ShiftKey() (shiftKey bool) {
	js.Rewrite("$<.ShiftKey")
	return shiftKey
}

// TargetTouches
func (*TouchEvent) TargetTouches() (targetTouches *TouchList) {
	js.Rewrite("$<.TargetTouches")
	return targetTouches
}

// Touches
func (*TouchEvent) Touches() (touches *TouchList) {
	js.Rewrite("$<.Touches")
	return touches
}

// Which
func (*TouchEvent) Which() (which uint8) {
	js.Rewrite("$<.Which")
	return which
}
