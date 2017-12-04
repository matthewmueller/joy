package paymentrequestupdateevent

import (
	"github.com/matthewmueller/joy/dom/paymentdetails"
	"github.com/matthewmueller/joy/dom/paymentrequestupdateeventinit"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.Event = (*PaymentRequestUpdateEvent)(nil)

// New fn
func New(kind string, eventinitdict *paymentrequestupdateeventinit.PaymentRequestUpdateEventInit) *PaymentRequestUpdateEvent {
	macro.Rewrite("new PaymentRequestUpdateEvent($1, $2)", kind, eventinitdict)
	return &PaymentRequestUpdateEvent{}
}

// PaymentRequestUpdateEvent struct
// js:"PaymentRequestUpdateEvent,omit"
type PaymentRequestUpdateEvent struct {
}

// UpdateWith fn
// js:"updateWith"
func (*PaymentRequestUpdateEvent) UpdateWith(d *paymentdetails.PaymentDetails) {
	macro.Rewrite("$_.updateWith($1)", d)
}

// InitEvent fn
// js:"initEvent"
func (*PaymentRequestUpdateEvent) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	macro.Rewrite("$_.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

// PreventDefault fn
// js:"preventDefault"
func (*PaymentRequestUpdateEvent) PreventDefault() {
	macro.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*PaymentRequestUpdateEvent) StopImmediatePropagation() {
	macro.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*PaymentRequestUpdateEvent) StopPropagation() {
	macro.Rewrite("$_.stopPropagation()")
}

// Bubbles prop
// js:"bubbles"
func (*PaymentRequestUpdateEvent) Bubbles() (bubbles bool) {
	macro.Rewrite("$_.bubbles")
	return bubbles
}

// Cancelable prop
// js:"cancelable"
func (*PaymentRequestUpdateEvent) Cancelable() (cancelable bool) {
	macro.Rewrite("$_.cancelable")
	return cancelable
}

// CancelBubble prop
// js:"cancelBubble"
func (*PaymentRequestUpdateEvent) CancelBubble() (cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble")
	return cancelBubble
}

// SetCancelBubble prop
// js:"cancelBubble"
func (*PaymentRequestUpdateEvent) SetCancelBubble(cancelBubble bool) {
	macro.Rewrite("$_.cancelBubble = $1", cancelBubble)
}

// CurrentTarget prop
// js:"currentTarget"
func (*PaymentRequestUpdateEvent) CurrentTarget() (currentTarget window.EventTarget) {
	macro.Rewrite("$_.currentTarget")
	return currentTarget
}

// DefaultPrevented prop
// js:"defaultPrevented"
func (*PaymentRequestUpdateEvent) DefaultPrevented() (defaultPrevented bool) {
	macro.Rewrite("$_.defaultPrevented")
	return defaultPrevented
}

// EventPhase prop
// js:"eventPhase"
func (*PaymentRequestUpdateEvent) EventPhase() (eventPhase uint8) {
	macro.Rewrite("$_.eventPhase")
	return eventPhase
}

// IsTrusted prop
// js:"isTrusted"
func (*PaymentRequestUpdateEvent) IsTrusted() (isTrusted bool) {
	macro.Rewrite("$_.isTrusted")
	return isTrusted
}

// ReturnValue prop
// js:"returnValue"
func (*PaymentRequestUpdateEvent) ReturnValue() (returnValue bool) {
	macro.Rewrite("$_.returnValue")
	return returnValue
}

// SetReturnValue prop
// js:"returnValue"
func (*PaymentRequestUpdateEvent) SetReturnValue(returnValue bool) {
	macro.Rewrite("$_.returnValue = $1", returnValue)
}

// SrcElement prop
// js:"srcElement"
func (*PaymentRequestUpdateEvent) SrcElement() (srcElement window.Element) {
	macro.Rewrite("$_.srcElement")
	return srcElement
}

// Target prop
// js:"target"
func (*PaymentRequestUpdateEvent) Target() (target window.EventTarget) {
	macro.Rewrite("$_.target")
	return target
}

// TimeStamp prop
// js:"timeStamp"
func (*PaymentRequestUpdateEvent) TimeStamp() (timeStamp uint64) {
	macro.Rewrite("$_.timeStamp")
	return timeStamp
}

// Type prop
// js:"type"
func (*PaymentRequestUpdateEvent) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}
