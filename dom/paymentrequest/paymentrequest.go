package paymentrequest

import (
	"github.com/matthewmueller/golly/dom2/paymentaddress"
	"github.com/matthewmueller/golly/dom2/paymentdetails"
	"github.com/matthewmueller/golly/dom2/paymentmethoddata"
	"github.com/matthewmueller/golly/dom2/paymentoptions"
	"github.com/matthewmueller/golly/dom2/paymentresponse"
	"github.com/matthewmueller/golly/dom2/paymentshippingtype"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(methoddata []*paymentmethoddata.PaymentMethodData, details *paymentdetails.PaymentDetails, options *paymentoptions.PaymentOptions) *PaymentRequest {
	js.Rewrite("PaymentRequest")
	return &PaymentRequest{}
}

// PaymentRequest struct
// js:"PaymentRequest,omit"
type PaymentRequest struct {
	window.EventTarget
}

// Abort fn
func (*PaymentRequest) Abort() {
	js.Rewrite("await $<.abort()")
}

// Show fn
func (*PaymentRequest) Show() (p *paymentresponse.PaymentResponse) {
	js.Rewrite("await $<.show()")
	return p
}

// Onshippingaddresschange prop
func (*PaymentRequest) Onshippingaddresschange() (onshippingaddresschange func(window.Event)) {
	js.Rewrite("$<.onshippingaddresschange")
	return onshippingaddresschange
}

// Onshippingaddresschange prop
func (*PaymentRequest) SetOnshippingaddresschange(onshippingaddresschange func(window.Event)) {
	js.Rewrite("$<.onshippingaddresschange = $1", onshippingaddresschange)
}

// Onshippingoptionchange prop
func (*PaymentRequest) Onshippingoptionchange() (onshippingoptionchange func(window.Event)) {
	js.Rewrite("$<.onshippingoptionchange")
	return onshippingoptionchange
}

// Onshippingoptionchange prop
func (*PaymentRequest) SetOnshippingoptionchange(onshippingoptionchange func(window.Event)) {
	js.Rewrite("$<.onshippingoptionchange = $1", onshippingoptionchange)
}

// ShippingAddress prop
func (*PaymentRequest) ShippingAddress() (shippingAddress *paymentaddress.PaymentAddress) {
	js.Rewrite("$<.shippingAddress")
	return shippingAddress
}

// ShippingOption prop
func (*PaymentRequest) ShippingOption() (shippingOption string) {
	js.Rewrite("$<.shippingOption")
	return shippingOption
}

// ShippingType prop
func (*PaymentRequest) ShippingType() (shippingType *paymentshippingtype.PaymentShippingType) {
	js.Rewrite("$<.shippingType")
	return shippingType
}
