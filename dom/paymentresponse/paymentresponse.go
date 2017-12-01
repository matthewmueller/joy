package paymentresponse

import (
	"github.com/matthewmueller/golly/dom/paymentaddress"
	"github.com/matthewmueller/golly/dom/paymentcomplete"
	"github.com/matthewmueller/golly/js"
)

// PaymentResponse struct
// js:"PaymentResponse,omit"
type PaymentResponse struct {
}

// Complete fn
// js:"complete"
func (*PaymentResponse) Complete(result *paymentcomplete.PaymentComplete) {
	js.Rewrite("await $_.complete($1)", result)
}

// ToJSON fn
// js:"toJSON"
func (*PaymentResponse) ToJSON() (i interface{}) {
	js.Rewrite("$_.toJSON()")
	return i
}

// Details prop
// js:"details"
func (*PaymentResponse) Details() (details interface{}) {
	js.Rewrite("$_.details")
	return details
}

// MethodName prop
// js:"methodName"
func (*PaymentResponse) MethodName() (methodName string) {
	js.Rewrite("$_.methodName")
	return methodName
}

// PayerEmail prop
// js:"payerEmail"
func (*PaymentResponse) PayerEmail() (payerEmail string) {
	js.Rewrite("$_.payerEmail")
	return payerEmail
}

// PayerName prop
// js:"payerName"
func (*PaymentResponse) PayerName() (payerName string) {
	js.Rewrite("$_.payerName")
	return payerName
}

// PayerPhone prop
// js:"payerPhone"
func (*PaymentResponse) PayerPhone() (payerPhone string) {
	js.Rewrite("$_.payerPhone")
	return payerPhone
}

// ShippingAddress prop
// js:"shippingAddress"
func (*PaymentResponse) ShippingAddress() (shippingAddress *paymentaddress.PaymentAddress) {
	js.Rewrite("$_.shippingAddress")
	return shippingAddress
}

// ShippingOption prop
// js:"shippingOption"
func (*PaymentResponse) ShippingOption() (shippingOption string) {
	js.Rewrite("$_.shippingOption")
	return shippingOption
}
