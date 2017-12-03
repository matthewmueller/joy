package window

import "github.com/matthewmueller/joy/js"

var _ UIEvent = (*KeyboardEvent)(nil)
var _ Event = (*KeyboardEvent)(nil)

// NewKeyboardEvent fn
func NewKeyboardEvent(typearg string, eventinitdict *KeyboardEventInit) *KeyboardEvent {
	js.Rewrite("KeyboardEvent")
	return &KeyboardEvent{}
}

// KeyboardEvent struct
// js:"KeyboardEvent,omit"
type KeyboardEvent struct {
}

// GetModifierState fn
// js:"getModifierState"
func (*KeyboardEvent) GetModifierState(keyArg string) (b bool) {
	js.Rewrite("$_.getModifierState($1)", keyArg)
	return b
}

// InitKeyboardEvent fn
// js:"initKeyboardEvent"
func (*KeyboardEvent) InitKeyboardEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, keyArg string, locationArg uint, modifiersListArg string, repeat bool, locale string) {
	js.Rewrite("$_.initKeyboardEvent($1, $2, $3, $4, $5, $6, $7, $8, $9)", typeArg, canBubbleArg, cancelableArg, viewArg, keyArg, locationArg, modifiersListArg, repeat, locale)
}

// InitUIEvent fn
// js:"initUIEvent"
func (*KeyboardEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	js.Rewrite("$_.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

// InitEvent fn
// js:"initEvent"
func (*KeyboardEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*KeyboardEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*KeyboardEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*KeyboardEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// AltKey prop
// js:"altKey"
func (*KeyboardEvent) AltKey() (altKey bool) {
	js.Rewrite("$_.altKey")
	return altKey
}

// Char prop
// js:"char"
func (*KeyboardEvent) Char() (char string) {
	js.Rewrite("$_.char")
	return char
}

// CharCode prop
// js:"charCode"
func (*KeyboardEvent) CharCode() (charCode int8) {
	js.Rewrite("$_.charCode")
	return charCode
}

// CtrlKey prop
// js:"ctrlKey"
func (*KeyboardEvent) CtrlKey() (ctrlKey bool) {
	js.Rewrite("$_.ctrlKey")
	return ctrlKey
}

// Key prop
// js:"key"
func (*KeyboardEvent) Key() (key string) {
	js.Rewrite("$_.key")
	return key
}

// KeyCode prop
// js:"keyCode"
func (*KeyboardEvent) KeyCode() (keyCode int8) {
	js.Rewrite("$_.keyCode")
	return keyCode
}

// Locale prop
// js:"locale"
func (*KeyboardEvent) Locale() (locale string) {
	js.Rewrite("$_.locale")
	return locale
}

// Location prop
// js:"location"
func (*KeyboardEvent) Location() (location uint) {
	js.Rewrite("$_.location")
	return location
}

// MetaKey prop
// js:"metaKey"
func (*KeyboardEvent) MetaKey() (metaKey bool) {
	js.Rewrite("$_.metaKey")
	return metaKey
}

// Repeat prop
// js:"repeat"
func (*KeyboardEvent) Repeat() (repeat bool) {
	js.Rewrite("$_.repeat")
	return repeat
}

// ShiftKey prop
// js:"shiftKey"
func (*KeyboardEvent) ShiftKey() (shiftKey bool) {
	js.Rewrite("$_.shiftKey")
	return shiftKey
}

// Which prop
// js:"which"
func (*KeyboardEvent) Which() (which int8) {
	js.Rewrite("$_.which")
	return which
}

// Detail prop
// js:"detail"
func (*KeyboardEvent) Detail() (detail int) {
	js.Rewrite("$_.detail")
	return detail
}

// View prop
// js:"view"
func (*KeyboardEvent) View() (view *Window) {
	js.Rewrite("$_.view")
	return view
}

// Bubbles prop
// js:"bubbles"
func (*KeyboardEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*KeyboardEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*KeyboardEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*KeyboardEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*KeyboardEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*KeyboardEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*KeyboardEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*KeyboardEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*KeyboardEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*KeyboardEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*KeyboardEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*KeyboardEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*KeyboardEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*KeyboardEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
