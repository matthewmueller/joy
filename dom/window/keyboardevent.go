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

// GetModifierState fn
func (*KeyboardEvent) GetModifierState(keyArg string) (b bool) {
	js.Rewrite("$<.getModifierState($1)", keyArg)
	return b
}

// InitKeyboardEvent fn
func (*KeyboardEvent) InitKeyboardEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, keyArg string, locationArg uint, modifiersListArg string, repeat bool, locale string) {
	js.Rewrite("$<.initKeyboardEvent($1, $2, $3, $4, $5, $6, $7, $8, $9)", typeArg, canBubbleArg, cancelableArg, viewArg, keyArg, locationArg, modifiersListArg, repeat, locale)
}

// AltKey prop
func (*KeyboardEvent) AltKey() (altKey bool) {
	js.Rewrite("$<.altKey")
	return altKey
}

// Char prop
func (*KeyboardEvent) Char() (char string) {
	js.Rewrite("$<.char")
	return char
}

// CharCode prop
func (*KeyboardEvent) CharCode() (charCode int8) {
	js.Rewrite("$<.charCode")
	return charCode
}

// CtrlKey prop
func (*KeyboardEvent) CtrlKey() (ctrlKey bool) {
	js.Rewrite("$<.ctrlKey")
	return ctrlKey
}

// Key prop
func (*KeyboardEvent) Key() (key string) {
	js.Rewrite("$<.key")
	return key
}

// KeyCode prop
func (*KeyboardEvent) KeyCode() (keyCode int8) {
	js.Rewrite("$<.keyCode")
	return keyCode
}

// Locale prop
func (*KeyboardEvent) Locale() (locale string) {
	js.Rewrite("$<.locale")
	return locale
}

// Location prop
func (*KeyboardEvent) Location() (location uint) {
	js.Rewrite("$<.location")
	return location
}

// MetaKey prop
func (*KeyboardEvent) MetaKey() (metaKey bool) {
	js.Rewrite("$<.metaKey")
	return metaKey
}

// Repeat prop
func (*KeyboardEvent) Repeat() (repeat bool) {
	js.Rewrite("$<.repeat")
	return repeat
}

// ShiftKey prop
func (*KeyboardEvent) ShiftKey() (shiftKey bool) {
	js.Rewrite("$<.shiftKey")
	return shiftKey
}

// Which prop
func (*KeyboardEvent) Which() (which int8) {
	js.Rewrite("$<.which")
	return which
}
