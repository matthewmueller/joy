package paymentrequestupdateevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"PaymentRequestUpdateEvent,omit"
type PaymentRequestUpdateEvent struct {
	window.Event
}

// UpdateWith
func (*PaymentRequestUpdateEvent) UpdateWith(d *paymentdetails.PaymentDetails) {
	js.Rewrite("$<.UpdateWith($1)", d)
}
