package paymentdetails

import "github.com/matthewmueller/golly/dom2/paymentitem"

type PaymentDetails struct {
	displayItems    *[]*paymentitem.PaymentItem
	err             *string
	modifiers       *[]*paymentdetailsmodifier.PaymentDetailsModifier
	shippingOptions *[]*paymentshippingoption.PaymentShippingOption
	total           *paymentitem.PaymentItem
}
