package window

import (
	"github.com/matthewmueller/joy/dom/storage"
	"github.com/matthewmueller/joy/js"
)

var _ Event = (*StorageEvent)(nil)

// StorageEvent struct
// js:"StorageEvent,omit"
type StorageEvent struct {
}

// InitStorageEvent fn
// js:"initStorageEvent"
func (*StorageEvent) InitStorageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, keyArg string, oldValueArg interface{}, newValueArg interface{}, urlArg string, storageAreaArg *storage.Storage) {
	js.Rewrite("$_.initStorageEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, keyArg, oldValueArg, newValueArg, urlArg, storageAreaArg)
}

// InitEvent fn
// js:"initEvent"
func (*StorageEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*StorageEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*StorageEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*StorageEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// Key prop
// js:"key"
func (*StorageEvent) Key() (key string) {
	js.Rewrite("$_.key")
	return key
}

// NewValue prop
// js:"newValue"
func (*StorageEvent) NewValue() (newValue interface{}) {
	js.Rewrite("$_.newValue")
	return newValue
}

// OldValue prop
// js:"oldValue"
func (*StorageEvent) OldValue() (oldValue interface{}) {
	js.Rewrite("$_.oldValue")
	return oldValue
}

// StorageArea prop
// js:"storageArea"
func (*StorageEvent) StorageArea() (storageArea *storage.Storage) {
	js.Rewrite("$_.storageArea")
	return storageArea
}

// URL prop
// js:"url"
func (*StorageEvent) URL() (url string) {
	js.Rewrite("$_.url")
	return url
}

// Bubbles prop
// js:"bubbles"
func (*StorageEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*StorageEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*StorageEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*StorageEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*StorageEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*StorageEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*StorageEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*StorageEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*StorageEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*StorageEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*StorageEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*StorageEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*StorageEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*StorageEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
