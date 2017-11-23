package paymentrequest

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"PaymentRequest,omit"
type PaymentRequest struct {
	window.EventTarget
}

// Abort
func (*PaymentRequest) Abort() {
	js.Rewrite("await $<.Abort()")
}

// Show
func (*PaymentRequest) Show() (p *paymentresponse.PaymentResponse) {
	js.Rewrite("await $<.Show()")
	return p
}

// Onshippingaddresschange
func (*PaymentRequest) Onshippingaddresschange() (onshippingaddresschange func(window.Event)) {
	js.Rewrite("$<.Onshippingaddresschange")
	return onshippingaddresschange
}

// Onshippingaddresschange
func (*PaymentRequest) SetOnshippingaddresschange(onshippingaddresschange func(window.Event)) {
	js.Rewrite("$<.Onshippingaddresschange = $1", onshippingaddresschange)
}

// Onshippingoptionchange
func (*PaymentRequest) Onshippingoptionchange() (onshippingoptionchange func(window.Event)) {
	js.Rewrite("$<.Onshippingoptionchange")
	return onshippingoptionchange
}

// Onshippingoptionchange
func (*PaymentRequest) SetOnshippingoptionchange(onshippingoptionchange func(window.Event)) {
	js.Rewrite("$<.Onshippingoptionchange = $1", onshippingoptionchange)
}

// ShippingAddress
func (*PaymentRequest) ShippingAddress() (shippingAddress *paymentaddress.PaymentAddress) {
	js.Rewrite("$<.ShippingAddress")
	return shippingAddress
}

// ShippingOption
func (*PaymentRequest) ShippingOption() (shippingOption string) {
	js.Rewrite("$<.ShippingOption")
	return shippingOption
}

// ShippingType
func (*PaymentRequest) ShippingType() (shippingType *paymentshippingtype.PaymentShippingType) {
	js.Rewrite("$<.ShippingType")
	return shippingType
}
