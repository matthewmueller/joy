package paymentshippingoption

import "github.com/matthewmueller/golly/dom2/paymentcurrencyamount"

type PaymentShippingOption struct {
	amount   *paymentcurrencyamount.PaymentCurrencyAmount
	id       string
	label    string
	selected *bool
}
