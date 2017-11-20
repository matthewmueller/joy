package paymentrequest

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/paymentaddress"
	"github.com/matthewmueller/golly/internal/gendom/dom/paymentresponse"
	"github.com/matthewmueller/golly/internal/gendom/dom/paymentshippingtype"
	"github.com/matthewmueller/golly/js"
)

type PaymentRequest struct {
	*eventtarget.EventTarget
}

func (*PaymentRequest) Abort() {
	js.Rewrite("await $<.abort()")
}

func (*PaymentRequest) Show() (p *paymentresponse.PaymentResponse) {
	js.Rewrite("await $<.show()")
	return p
}

func (*PaymentRequest) GetOnshippingaddresschange() (shippingaddresschange *event.Event) {
	js.Rewrite("$<.onshippingaddresschange")
	return shippingaddresschange
}

func (*PaymentRequest) SetOnshippingaddresschange(shippingaddresschange *event.Event) {
	js.Rewrite("$<.onshippingaddresschange = $1", shippingaddresschange)
}

func (*PaymentRequest) GetOnshippingoptionchange() (shippingoptionchange *event.Event) {
	js.Rewrite("$<.onshippingoptionchange")
	return shippingoptionchange
}

func (*PaymentRequest) SetOnshippingoptionchange(shippingoptionchange *event.Event) {
	js.Rewrite("$<.onshippingoptionchange = $1", shippingoptionchange)
}

func (*PaymentRequest) GetShippingAddress() (shippingAddress *paymentaddress.PaymentAddress) {
	js.Rewrite("$<.shippingAddress")
	return shippingAddress
}

func (*PaymentRequest) GetShippingOption() (shippingOption string) {
	js.Rewrite("$<.shippingOption")
	return shippingOption
}

func (*PaymentRequest) GetShippingType() (shippingType *paymentshippingtype.PaymentShippingType) {
	js.Rewrite("$<.shippingType")
	return shippingType
}
