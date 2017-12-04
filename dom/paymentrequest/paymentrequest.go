package paymentrequest

import (
	"github.com/matthewmueller/joy/dom/paymentaddress"
	"github.com/matthewmueller/joy/dom/paymentdetails"
	"github.com/matthewmueller/joy/dom/paymentmethoddata"
	"github.com/matthewmueller/joy/dom/paymentoptions"
	"github.com/matthewmueller/joy/dom/paymentresponse"
	"github.com/matthewmueller/joy/dom/paymentshippingtype"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.EventTarget = (*PaymentRequest)(nil)

// New fn
func New(methoddata []*paymentmethoddata.PaymentMethodData, details *paymentdetails.PaymentDetails, options *paymentoptions.PaymentOptions) *PaymentRequest {
	macro.Rewrite("new PaymentRequest($1, $2, $3)", methoddata, details, options)
	return &PaymentRequest{}
}

// PaymentRequest struct
// js:"PaymentRequest,omit"
type PaymentRequest struct {
}

// Abort fn
// js:"abort"
func (*PaymentRequest) Abort() {
	macro.Rewrite("await $_.abort()")
}

// Show fn
// js:"show"
func (*PaymentRequest) Show() (p *paymentresponse.PaymentResponse) {
	macro.Rewrite("await $_.show()")
	return p
}

// AddEventListener fn
// js:"addEventListener"
func (*PaymentRequest) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*PaymentRequest) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*PaymentRequest) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onshippingaddresschange prop
// js:"onshippingaddresschange"
func (*PaymentRequest) Onshippingaddresschange() (onshippingaddresschange func(window.Event)) {
	macro.Rewrite("$_.onshippingaddresschange")
	return onshippingaddresschange
}

// SetOnshippingaddresschange prop
// js:"onshippingaddresschange"
func (*PaymentRequest) SetOnshippingaddresschange(onshippingaddresschange func(window.Event)) {
	macro.Rewrite("$_.onshippingaddresschange = $1", onshippingaddresschange)
}

// Onshippingoptionchange prop
// js:"onshippingoptionchange"
func (*PaymentRequest) Onshippingoptionchange() (onshippingoptionchange func(window.Event)) {
	macro.Rewrite("$_.onshippingoptionchange")
	return onshippingoptionchange
}

// SetOnshippingoptionchange prop
// js:"onshippingoptionchange"
func (*PaymentRequest) SetOnshippingoptionchange(onshippingoptionchange func(window.Event)) {
	macro.Rewrite("$_.onshippingoptionchange = $1", onshippingoptionchange)
}

// ShippingAddress prop
// js:"shippingAddress"
func (*PaymentRequest) ShippingAddress() (shippingAddress *paymentaddress.PaymentAddress) {
	macro.Rewrite("$_.shippingAddress")
	return shippingAddress
}

// ShippingOption prop
// js:"shippingOption"
func (*PaymentRequest) ShippingOption() (shippingOption string) {
	macro.Rewrite("$_.shippingOption")
	return shippingOption
}

// ShippingType prop
// js:"shippingType"
func (*PaymentRequest) ShippingType() (shippingType *paymentshippingtype.PaymentShippingType) {
	macro.Rewrite("$_.shippingType")
	return shippingType
}
