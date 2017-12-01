package paymentshippingoption

import "github.com/matthewmueller/golly/dom/paymentcurrencyamount"

type PaymentShippingOption struct {
	amount   *paymentcurrencyamount.PaymentCurrencyAmount
	id       string
	label    string
	selected *bool
}
