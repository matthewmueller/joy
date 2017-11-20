package paymentresponse

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/paymentaddress"
	"github.com/matthewmueller/golly/internal/gendom/dom/paymentcomplete"
	"github.com/matthewmueller/golly/js"
)

type PaymentResponse struct {
}

func (*PaymentResponse) Complete(result *paymentcomplete.PaymentComplete) {
	js.Rewrite("await $<.complete($1)", result)
}

func (*PaymentResponse) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*PaymentResponse) GetDetails() (details interface{}) {
	js.Rewrite("$<.details")
	return details
}

func (*PaymentResponse) GetMethodName() (methodName string) {
	js.Rewrite("$<.methodName")
	return methodName
}

func (*PaymentResponse) GetPayerEmail() (payerEmail string) {
	js.Rewrite("$<.payerEmail")
	return payerEmail
}

func (*PaymentResponse) GetPayerName() (payerName string) {
	js.Rewrite("$<.payerName")
	return payerName
}

func (*PaymentResponse) GetPayerPhone() (payerPhone string) {
	js.Rewrite("$<.payerPhone")
	return payerPhone
}

func (*PaymentResponse) GetShippingAddress() (shippingAddress *paymentaddress.PaymentAddress) {
	js.Rewrite("$<.shippingAddress")
	return shippingAddress
}

func (*PaymentResponse) GetShippingOption() (shippingOption string) {
	js.Rewrite("$<.shippingOption")
	return shippingOption
}
