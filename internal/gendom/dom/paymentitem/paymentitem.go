package paymentitem

import "github.com/matthewmueller/golly/internal/gendom/dom/paymentcurrencyamount"

type PaymentItem struct {
	amount  *paymentcurrencyamount.PaymentCurrencyAmount
	label   string
	pending *bool
}
