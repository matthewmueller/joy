package paymentrequestupdateevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/paymentdetails"
	"github.com/matthewmueller/golly/js"
)

type PaymentRequestUpdateEvent struct {
	*event.Event
}

func (*PaymentRequestUpdateEvent) UpdateWith(d *paymentdetails.PaymentDetails) {
	js.Rewrite("$<.updateWith($1)", d)
}
