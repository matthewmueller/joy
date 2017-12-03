package paymentitem

import "github.com/matthewmueller/joy/dom/paymentcurrencyamount"

type PaymentItem struct {
	amount  *paymentcurrencyamount.PaymentCurrencyAmount
	label   string
	pending *bool
}
