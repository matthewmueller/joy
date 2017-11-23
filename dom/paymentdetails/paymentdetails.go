package paymentdetails

import (
	"github.com/matthewmueller/golly/dom/paymentdetailsmodifier"
	"github.com/matthewmueller/golly/dom/paymentitem"
	"github.com/matthewmueller/golly/dom/paymentshippingoption"
)

type PaymentDetails struct {
	displayItems    *[]*paymentitem.PaymentItem
	err             *string
	modifiers       *[]*paymentdetailsmodifier.PaymentDetailsModifier
	shippingOptions *[]*paymentshippingoption.PaymentShippingOption
	total           *paymentitem.PaymentItem
}
