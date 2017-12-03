package mediaencryptedevent

import (
	"github.com/matthewmueller/joy/dom/mediaencryptedeventinit"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*MediaEncryptedEvent)(nil)

// New fn
func New(kind string, eventinitdict *mediaencryptedeventinit.MediaEncryptedEventInit) *MediaEncryptedEvent {
	macro.Rewrite("MediaEncryptedEvent")
	return &MediaEncryptedEvent{}
}

// MediaEncryptedEvent struct
// js:"MediaEncryptedEvent,omit"
type MediaEncryptedEvent struct {
}

// InitEvent fn
// js:"initEvent"
func (*MediaEncryptedEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*MediaEncryptedEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MediaEncryptedEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MediaEncryptedEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// InitData prop
// js:"initData"
func (*MediaEncryptedEvent) InitData() (initData []byte) {
	macro.Rewrite("$_.initData")
	return initData
}

// InitDataType prop
// js:"initDataType"
func (*MediaEncryptedEvent) InitDataType() (initDataType string) {
	macro.Rewrite("$_.initDataType")
	return initDataType
}

// Bubbles prop
// js:"bubbles"
func (*MediaEncryptedEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*MediaEncryptedEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*MediaEncryptedEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*MediaEncryptedEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*MediaEncryptedEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*MediaEncryptedEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*MediaEncryptedEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*MediaEncryptedEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*MediaEncryptedEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*MediaEncryptedEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*MediaEncryptedEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*MediaEncryptedEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*MediaEncryptedEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*MediaEncryptedEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
