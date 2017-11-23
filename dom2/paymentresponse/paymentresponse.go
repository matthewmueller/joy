package paymentresponse

import (
	"github.com/matthewmueller/golly/dom2/paymentaddress"
	"github.com/matthewmueller/golly/js"
)

// js:"PaymentResponse,omit"
type PaymentResponse struct {
}

// Complete
func (*PaymentResponse) Complete(result *paymentcomplete.PaymentComplete) {
	js.Rewrite("await $<.Complete($1)", result)
}

// ToJSON
func (*PaymentResponse) ToJSON() (i interface{}) {
	js.Rewrite("$<.ToJSON()")
	return i
}

// Details
func (*PaymentResponse) Details() (details interface{}) {
	js.Rewrite("$<.Details")
	return details
}

// MethodName
func (*PaymentResponse) MethodName() (methodName string) {
	js.Rewrite("$<.MethodName")
	return methodName
}

// PayerEmail
func (*PaymentResponse) PayerEmail() (payerEmail string) {
	js.Rewrite("$<.PayerEmail")
	return payerEmail
}

// PayerName
func (*PaymentResponse) PayerName() (payerName string) {
	js.Rewrite("$<.PayerName")
	return payerName
}

// PayerPhone
func (*PaymentResponse) PayerPhone() (payerPhone string) {
	js.Rewrite("$<.PayerPhone")
	return payerPhone
}

// ShippingAddress
func (*PaymentResponse) ShippingAddress() (shippingAddress *paymentaddress.PaymentAddress) {
	js.Rewrite("$<.ShippingAddress")
	return shippingAddress
}

// ShippingOption
func (*PaymentResponse) ShippingOption() (shippingOption string) {
	js.Rewrite("$<.ShippingOption")
	return shippingOption
}
