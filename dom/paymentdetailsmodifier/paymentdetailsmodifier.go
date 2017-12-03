package paymentdetailsmodifier

import "github.com/matthewmueller/joy/dom/paymentitem"

type PaymentDetailsModifier struct {
	additionalDisplayItems *[]*paymentitem.PaymentItem
	data                   *interface{}
	supportedMethods       []string
	total                  *paymentitem.PaymentItem
}
