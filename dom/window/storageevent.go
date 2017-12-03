package window

import (
	"github.com/matthewmueller/joy/dom/storage"
	"github.com/matthewmueller/joy/macro"
)

var _ Event = (*StorageEvent)(nil)

// StorageEvent struct
// js:"StorageEvent,omit"
type StorageEvent struct {
}

// InitStorageEvent fn
// js:"initStorageEvent"
func (*StorageEvent) InitStorageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, keyArg string, oldValueArg interface{}, newValueArg interface{}, urlArg string, storageAreaArg *storage.Storage) {
	macro.Rewrite("$_.initStorageEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, keyArg, oldValueArg, newValueArg, urlArg, storageAreaArg)
}

// InitEvent fn
// js:"initEvent"
func (*StorageEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*StorageEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*StorageEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*StorageEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// Key prop
// js:"key"
func (*StorageEvent) Key() (key string) {
	macro.Rewrite("$_.key")
	return key
}

// NewValue prop
// js:"newValue"
func (*StorageEvent) NewValue() (newValue interface{}) {
	macro.Rewrite("$_.newValue")
	return newValue
}

// OldValue prop
// js:"oldValue"
func (*StorageEvent) OldValue() (oldValue interface{}) {
	macro.Rewrite("$_.oldValue")
	return oldValue
}

// StorageArea prop
// js:"storageArea"
func (*StorageEvent) StorageArea() (storageArea *storage.Storage) {
	macro.Rewrite("$_.storageArea")
	return storageArea
}

// URL prop
// js:"url"
func (*StorageEvent) URL() (url string) {
	macro.Rewrite("$_.url")
	return url
}

// Bubbles prop
// js:"bubbles"
func (*StorageEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*StorageEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*StorageEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*StorageEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*StorageEvent) CurrentTarget() (currentTarget EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*StorageEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*StorageEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*StorageEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*StorageEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*StorageEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*StorageEvent) SrcElement() (srcElement Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*StorageEvent) Target() (target EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*StorageEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*StorageEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
