package window

import "github.com/matthewmueller/golly/js"

// NewKeyboardEvent fn
func NewKeyboardEvent(typearg string, eventinitdict *KeyboardEventInit) *KeyboardEvent {
	js.Rewrite("KeyboardEvent")
	return &KeyboardEvent{}
}

// KeyboardEvent struct
// js:"KeyboardEvent,omit"
type KeyboardEvent struct {
	UIEvent
}

// GetModifierState
func (*KeyboardEvent) GetModifierState(keyArg string) (b bool) {
	js.Rewrite("$<.GetModifierState($1)", keyArg)
	return b
}

// InitKeyboardEvent
func (*KeyboardEvent) InitKeyboardEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, keyArg string, locationArg uint, modifiersListArg string, repeat bool, locale string) {
	js.Rewrite("$<.InitKeyboardEvent($1, $2, $3, $4, $5, $6, $7, $8, $9)", typeArg, canBubbleArg, cancelableArg, viewArg, keyArg, locationArg, modifiersListArg, repeat, locale)
}

// AltKey
func (*KeyboardEvent) AltKey() (altKey bool) {
	js.Rewrite("$<.AltKey")
	return altKey
}

// Char
func (*KeyboardEvent) Char() (char string) {
	js.Rewrite("$<.Char")
	return char
}

// CharCode
func (*KeyboardEvent) CharCode() (charCode int8) {
	js.Rewrite("$<.CharCode")
	return charCode
}

// CtrlKey
func (*KeyboardEvent) CtrlKey() (ctrlKey bool) {
	js.Rewrite("$<.CtrlKey")
	return ctrlKey
}

// Key
func (*KeyboardEvent) Key() (key string) {
	js.Rewrite("$<.Key")
	return key
}

// KeyCode
func (*KeyboardEvent) KeyCode() (keyCode int8) {
	js.Rewrite("$<.KeyCode")
	return keyCode
}

// Locale
func (*KeyboardEvent) Locale() (locale string) {
	js.Rewrite("$<.Locale")
	return locale
}

// Location
func (*KeyboardEvent) Location() (location uint) {
	js.Rewrite("$<.Location")
	return location
}

// MetaKey
func (*KeyboardEvent) MetaKey() (metaKey bool) {
	js.Rewrite("$<.MetaKey")
	return metaKey
}

// Repeat
func (*KeyboardEvent) Repeat() (repeat bool) {
	js.Rewrite("$<.Repeat")
	return repeat
}

// ShiftKey
func (*KeyboardEvent) ShiftKey() (shiftKey bool) {
	js.Rewrite("$<.ShiftKey")
	return shiftKey
}

// Which
func (*KeyboardEvent) Which() (which int8) {
	js.Rewrite("$<.Which")
	return which
}
