package paymentdetails

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/paymentdetailsmodifier"
	"github.com/matthewmueller/golly/internal/gendom/dom/paymentitem"
	"github.com/matthewmueller/golly/internal/gendom/dom/paymentshippingoption"
)

type PaymentDetails struct {
	displayItems    *[]*paymentitem.PaymentItem
	err             *string
	modifiers       *[]*paymentdetailsmodifier.PaymentDetailsModifier
	shippingOptions *[]*paymentshippingoption.PaymentShippingOption
	total           *paymentitem.PaymentItem
}
