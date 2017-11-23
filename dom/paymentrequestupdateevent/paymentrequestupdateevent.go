package paymentrequestupdateevent

import (
	"github.com/matthewmueller/golly/dom/paymentdetails"
	"github.com/matthewmueller/golly/dom/paymentrequestupdateeventinit"
	"github.com/matthewmueller/golly/dom/window"
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

// UpdateWith fn
func (*PaymentRequestUpdateEvent) UpdateWith(d *paymentdetails.PaymentDetails) {
	js.Rewrite("$<.updateWith($1)", d)
}
