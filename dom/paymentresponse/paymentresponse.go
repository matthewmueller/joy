package paymentresponse

import (
	"github.com/matthewmueller/joy/dom/paymentaddress"
	"github.com/matthewmueller/joy/dom/paymentcomplete"
	"github.com/matthewmueller/joy/macro"
)

// PaymentResponse struct
// js:"PaymentResponse,omit"
type PaymentResponse struct {
}

// Complete fn
// js:"complete"
func (*PaymentResponse) Complete(result *paymentcomplete.PaymentComplete) {
	macro.Rewrite("await $_.complete($1)", result)
}

// ToJSON fn
// js:"toJSON"
func (*PaymentResponse) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

// Details prop
// js:"details"
func (*PaymentResponse) Details() (details interface{}) {
	macro.Rewrite("$_.details")
	return details
}

// MethodName prop
// js:"methodName"
func (*PaymentResponse) MethodName() (methodName string) {
	macro.Rewrite("$_.methodName")
	return methodName
}

// PayerEmail prop
// js:"payerEmail"
func (*PaymentResponse) PayerEmail() (payerEmail string) {
	macro.Rewrite("$_.payerEmail")
	return payerEmail
}

// PayerName prop
// js:"payerName"
func (*PaymentResponse) PayerName() (payerName string) {
	macro.Rewrite("$_.payerName")
	return payerName
}

// PayerPhone prop
// js:"payerPhone"
func (*PaymentResponse) PayerPhone() (payerPhone string) {
	macro.Rewrite("$_.payerPhone")
	return payerPhone
}

// ShippingAddress prop
// js:"shippingAddress"
func (*PaymentResponse) ShippingAddress() (shippingAddress *paymentaddress.PaymentAddress) {
	macro.Rewrite("$_.shippingAddress")
	return shippingAddress
}

// ShippingOption prop
// js:"shippingOption"
func (*PaymentResponse) ShippingOption() (shippingOption string) {
	macro.Rewrite("$_.shippingOption")
	return shippingOption
}
