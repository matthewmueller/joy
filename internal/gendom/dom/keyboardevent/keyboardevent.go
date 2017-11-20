package keyboardevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/uievent"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type KeyboardEvent struct {
	*uievent.UIEvent
}

func (*KeyboardEvent) GetModifierState(keyArg string) (b bool) {
	js.Rewrite("$<.getModifierState($1)", keyArg)
	return b
}

func (*KeyboardEvent) InitKeyboardEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *window.Window, keyArg string, locationArg uint, modifiersListArg string, repeat bool, locale string) {
	js.Rewrite("$<.initKeyboardEvent($1, $2, $3, $4, $5, $6, $7, $8, $9)", typeArg, canBubbleArg, cancelableArg, viewArg, keyArg, locationArg, modifiersListArg, repeat, locale)
}

func (*KeyboardEvent) GetAltKey() (altKey bool) {
	js.Rewrite("$<.altKey")
	return altKey
}

func (*KeyboardEvent) GetChar() (char string) {
	js.Rewrite("$<.char")
	return char
}

func (*KeyboardEvent) GetCharCode() (charCode int8) {
	js.Rewrite("$<.charCode")
	return charCode
}

func (*KeyboardEvent) GetCtrlKey() (ctrlKey bool) {
	js.Rewrite("$<.ctrlKey")
	return ctrlKey
}

func (*KeyboardEvent) GetKey() (key string) {
	js.Rewrite("$<.key")
	return key
}

func (*KeyboardEvent) GetKeyCode() (keyCode int8) {
	js.Rewrite("$<.keyCode")
	return keyCode
}

func (*KeyboardEvent) GetLocale() (locale string) {
	js.Rewrite("$<.locale")
	return locale
}

func (*KeyboardEvent) GetLocation() (location uint) {
	js.Rewrite("$<.location")
	return location
}

func (*KeyboardEvent) GetMetaKey() (metaKey bool) {
	js.Rewrite("$<.metaKey")
	return metaKey
}

func (*KeyboardEvent) GetRepeat() (repeat bool) {
	js.Rewrite("$<.repeat")
	return repeat
}

func (*KeyboardEvent) GetShiftKey() (shiftKey bool) {
	js.Rewrite("$<.shiftKey")
	return shiftKey
}

func (*KeyboardEvent) GetWhich() (which int8) {
	js.Rewrite("$<.which")
	return which
}
