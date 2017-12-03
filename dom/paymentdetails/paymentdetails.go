package paymentdetails

import (
	"github.com/matthewmueller/joy/dom/paymentdetailsmodifier"
	"github.com/matthewmueller/joy/dom/paymentitem"
	"github.com/matthewmueller/joy/dom/paymentshippingoption"
)

type PaymentDetails struct {
	displayItems    *[]*paymentitem.PaymentItem
	err             *string
	modifiers       *[]*paymentdetailsmodifier.PaymentDetailsModifier
	shippingOptions *[]*paymentshippingoption.PaymentShippingOption
	total           *paymentitem.PaymentItem
}
