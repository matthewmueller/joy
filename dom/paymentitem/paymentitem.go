package paymentitem

import "github.com/matthewmueller/golly/dom2/paymentcurrencyamount"

type PaymentItem struct {
	amount  *paymentcurrencyamount.PaymentCurrencyAmount
	label   string
	pending *bool
}
