package paymentrequest

import (
	"github.com/matthewmueller/golly/dom/paymentaddress"
	"github.com/matthewmueller/golly/dom/paymentdetails"
	"github.com/matthewmueller/golly/dom/paymentmethoddata"
	"github.com/matthewmueller/golly/dom/paymentoptions"
	"github.com/matthewmueller/golly/dom/paymentresponse"
	"github.com/matthewmueller/golly/dom/paymentshippingtype"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.EventTarget = (*PaymentRequest)(nil)

// New fn
func New(methoddata []*paymentmethoddata.PaymentMethodData, details *paymentdetails.PaymentDetails, options *paymentoptions.PaymentOptions) *PaymentRequest {
	js.Rewrite("PaymentRequest")
	return &PaymentRequest{}
}

// PaymentRequest struct
// js:"PaymentRequest,omit"
type PaymentRequest struct {
}

// Abort fn
// js:"abort"
func (*PaymentRequest) Abort() {
	js.Rewrite("await $_.abort()")
}

// Show fn
// js:"show"
func (*PaymentRequest) Show() (p *paymentresponse.PaymentResponse) {
	js.Rewrite("await $_.show()")
	return p
}

// AddEventListener fn
// js:"addEventListener"
func (*PaymentRequest) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*PaymentRequest) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*PaymentRequest) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onshippingaddresschange prop
// js:"onshippingaddresschange"
func (*PaymentRequest) Onshippingaddresschange() (onshippingaddresschange func(window.Event)) {
	js.Rewrite("$_.onshippingaddresschange")
	return onshippingaddresschange
}

// SetOnshippingaddresschange prop
// js:"onshippingaddresschange"
func (*PaymentRequest) SetOnshippingaddresschange(onshippingaddresschange func(window.Event)) {
	js.Rewrite("$_.onshippingaddresschange = $1", onshippingaddresschange)
}

// Onshippingoptionchange prop
// js:"onshippingoptionchange"
func (*PaymentRequest) Onshippingoptionchange() (onshippingoptionchange func(window.Event)) {
	js.Rewrite("$_.onshippingoptionchange")
	return onshippingoptionchange
}

// SetOnshippingoptionchange prop
// js:"onshippingoptionchange"
func (*PaymentRequest) SetOnshippingoptionchange(onshippingoptionchange func(window.Event)) {
	js.Rewrite("$_.onshippingoptionchange = $1", onshippingoptionchange)
}

// ShippingAddress prop
// js:"shippingAddress"
func (*PaymentRequest) ShippingAddress() (shippingAddress *paymentaddress.PaymentAddress) {
	js.Rewrite("$_.shippingAddress")
	return shippingAddress
}

// ShippingOption prop
// js:"shippingOption"
func (*PaymentRequest) ShippingOption() (shippingOption string) {
	js.Rewrite("$_.shippingOption")
	return shippingOption
}

// ShippingType prop
// js:"shippingType"
func (*PaymentRequest) ShippingType() (shippingType *paymentshippingtype.PaymentShippingType) {
	js.Rewrite("$_.shippingType")
	return shippingType
}
