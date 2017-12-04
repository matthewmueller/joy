package errorevent

import (
	"github.com/matthewmueller/joy/dom/erroreventinit"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*ErrorEvent)(nil)

// New fn
func New(typearg string, eventinitdict *erroreventinit.ErrorEventInit) *ErrorEvent {
	macro.Rewrite("new ErrorEvent($1, $2)", typearg, eventinitdict)
	return &ErrorEvent{}
}

// ErrorEvent struct
// js:"ErrorEvent,omit"
type ErrorEvent struct {
}

// InitErrorEvent fn
// js:"initErrorEvent"
func (*ErrorEvent) InitErrorEvent(typeArg string, canBubbleArg bool, cancelableArg bool, messageArg string, filenameArg string, linenoArg uint) {
	macro.Rewrite("$_.initErrorEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, messageArg, filenameArg, linenoArg)
}

// InitEvent fn
// js:"initEvent"
func (*ErrorEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*ErrorEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*ErrorEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*ErrorEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// Colno prop
// js:"colno"
func (*ErrorEvent) Colno() (colno uint) {
	macro.Rewrite("$_.colno")
	return colno
}

// Error prop
// js:"error"
func (*ErrorEvent) Error() (err interface{}) {
	macro.Rewrite("$_.error")
	return err
}

// Filename prop
// js:"filename"
func (*ErrorEvent) Filename() (filename string) {
	macro.Rewrite("$_.filename")
	return filename
}

// Lineno prop
// js:"lineno"
func (*ErrorEvent) Lineno() (lineno uint) {
	macro.Rewrite("$_.lineno")
	return lineno
}

// Message prop
// js:"message"
func (*ErrorEvent) Message() (message string) {
	macro.Rewrite("$_.message")
	return message
}

// Bubbles prop
// js:"bubbles"
func (*ErrorEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*ErrorEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*ErrorEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*ErrorEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*ErrorEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*ErrorEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*ErrorEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*ErrorEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*ErrorEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*ErrorEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*ErrorEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*ErrorEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*ErrorEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*ErrorEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
