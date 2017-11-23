package paymentdetailsmodifier

import "github.com/matthewmueller/golly/dom2/paymentitem"

type PaymentDetailsModifier struct {
	additionalDisplayItems *[]*paymentitem.PaymentItem
	data                   *interface{}
	supportedMethods       []string
	total                  *paymentitem.PaymentItem
}
