package paymentitem

type PaymentItem struct {
	amount  *paymentcurrencyamount.PaymentCurrencyAmount
	label   string
	pending *bool
}
