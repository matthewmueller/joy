package paymentdetails

import (
	"github.com/matthewmueller/golly/dom2/paymentdetailsmodifier"
	"github.com/matthewmueller/golly/dom2/paymentitem"
	"github.com/matthewmueller/golly/dom2/paymentshippingoption"
)

type PaymentDetails struct {
	displayItems    *[]*paymentitem.PaymentItem
	err             *string
	modifiers       *[]*paymentdetailsmodifier.PaymentDetailsModifier
	shippingOptions *[]*paymentshippingoption.PaymentShippingOption
	total           *paymentitem.PaymentItem
}
