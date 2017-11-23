package paymentrequestupdateevent

import (
	"github.com/matthewmueller/golly/dom2/paymentdetails"
	"github.com/matthewmueller/golly/dom2/paymentrequestupdateeventinit"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(kind string, eventinitdict *paymentrequestupdateeventinit.PaymentRequestUpdateEventInit) *PaymentRequestUpdateEvent {
	js.Rewrite("PaymentRequestUpdateEvent")
	return &PaymentRequestUpdateEvent{}
}

// PaymentRequestUpdateEvent struct
// js:"PaymentRequestUpdateEvent,omit"
type PaymentRequestUpdateEvent struct {
	window.Event
}

// UpdateWith
func (*PaymentRequestUpdateEvent) UpdateWith(d *paymentdetails.PaymentDetails) {
	js.Rewrite("$<.UpdateWith($1)", d)
}
