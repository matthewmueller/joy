package window

import (
	"github.com/matthewmueller/joy/dom/progresseventinit"
	"github.com/matthewmueller/joy/js"
)

var _ Event = (*ProgressEvent)(nil)

// NewProgressEvent fn
func NewProgressEvent(typearg string, eventinitdict *progresseventinit.ProgressEventInit) *ProgressEvent {
	js.Rewrite("ProgressEvent")
	return &ProgressEvent{}
}

// ProgressEvent struct
// js:"ProgressEvent,omit"
type ProgressEvent struct {
}

// InitProgressEvent fn
// js:"initProgressEvent"
func (*ProgressEvent) InitProgressEvent(typeArg string, canBubbleArg bool, cancelableArg bool, lengthComputableArg bool, loadedArg uint64, totalArg uint64) {
	js.Rewrite("$_.initProgressEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, lengthComputableArg, loadedArg, totalArg)
}

// InitEvent fn
// js:"initEvent"
func (*ProgressEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*ProgressEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*ProgressEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*ProgressEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// LengthComputable prop
// js:"lengthComputable"
func (*ProgressEvent) LengthComputable() (lengthComputable bool) {
	js.Rewrite("$_.lengthComputable")
	return lengthComputable
}

// Loaded prop
// js:"loaded"
func (*ProgressEvent) Loaded() (loaded uint64) {
	js.Rewrite("$_.loaded")
	return loaded
}

// Total prop
// js:"total"
func (*ProgressEvent) Total() (total uint64) {
	js.Rewrite("$_.total")
	return total
}

// Bubbles prop
// js:"bubbles"
func (*ProgressEvent) Bubbles() (bubbles bool) {
	js.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*ProgressEvent) Cancelable() (cancelable bool) {
	js.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*ProgressEvent) CancelBubble() (cancelBubble bool) {
	js.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*ProgressEvent) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*ProgressEvent) CurrentTarget() (currentTarget EventTarget) {
	js.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*ProgressEvent) DefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*ProgressEvent) EventPhase() (eventPhase uint8) {
	js.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*ProgressEvent) IsTrusted() (isTrusted bool) {
	js.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*ProgressEvent) ReturnValue() (returnValue bool) {
	js.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*ProgressEvent) SetReturnValue(returnValue bool) {
	js.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*ProgressEvent) SrcElement() (srcElement Element) {
	js.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*ProgressEvent) Target() (target EventTarget) {
	js.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*ProgressEvent) TimeStamp() (timeStamp uint64) {
	js.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*ProgressEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
