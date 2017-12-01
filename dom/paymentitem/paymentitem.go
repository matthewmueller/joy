package paymentitem

import "github.com/matthewmueller/golly/dom/paymentcurrencyamount"

type PaymentItem struct {
	amount  *paymentcurrencyamount.PaymentCurrencyAmount
	label   string
	pending *bool
}
