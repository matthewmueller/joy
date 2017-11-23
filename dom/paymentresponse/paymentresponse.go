package paymentresponse

import (
	"github.com/matthewmueller/golly/dom2/paymentaddress"
	"github.com/matthewmueller/golly/dom2/paymentcomplete"
	"github.com/matthewmueller/golly/js"
)

// PaymentResponse struct
// js:"PaymentResponse,omit"
type PaymentResponse struct {
}

// Complete fn
func (*PaymentResponse) Complete(result *paymentcomplete.PaymentComplete) {
	js.Rewrite("await $<.complete($1)", result)
}

// ToJSON fn
func (*PaymentResponse) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

// Details prop
func (*PaymentResponse) Details() (details interface{}) {
	js.Rewrite("$<.details")
	return details
}

// MethodName prop
func (*PaymentResponse) MethodName() (methodName string) {
	js.Rewrite("$<.methodName")
	return methodName
}

// PayerEmail prop
func (*PaymentResponse) PayerEmail() (payerEmail string) {
	js.Rewrite("$<.payerEmail")
	return payerEmail
}

// PayerName prop
func (*PaymentResponse) PayerName() (payerName string) {
	js.Rewrite("$<.payerName")
	return payerName
}

// PayerPhone prop
func (*PaymentResponse) PayerPhone() (payerPhone string) {
	js.Rewrite("$<.payerPhone")
	return payerPhone
}

// ShippingAddress prop
func (*PaymentResponse) ShippingAddress() (shippingAddress *paymentaddress.PaymentAddress) {
	js.Rewrite("$<.shippingAddress")
	return shippingAddress
}

// ShippingOption prop
func (*PaymentResponse) ShippingOption() (shippingOption string) {
	js.Rewrite("$<.shippingOption")
	return shippingOption
}
