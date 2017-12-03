package window

import (
	"github.com/matthewmueller/joy/dom/registrationoptions"
	"github.com/matthewmueller/joy/js"
)

var _ EventTarget = (*ServiceWorkerContainer)(nil)

// ServiceWorkerContainer struct
// js:"ServiceWorkerContainer,omit"
type ServiceWorkerContainer struct {
}

// GetRegistration fn
// js:"getRegistration"
func (*ServiceWorkerContainer) GetRegistration(clientURL *string) (i interface{}) {
	js.Rewrite("await $_.getRegistration($1)", clientURL)
	return i
}

// GetRegistrations fn
// js:"getRegistrations"
func (*ServiceWorkerContainer) GetRegistrations() (s []*ServiceWorkerRegistration) {
	js.Rewrite("await $_.getRegistrations()")
	return s
}

// Register fn
// js:"register"
func (*ServiceWorkerContainer) Register(scriptURL string, options *registrationoptions.RegistrationOptions) (s *ServiceWorkerRegistration) {
	js.Rewrite("await $_.register($1, $2)", scriptURL, options)
	return s
}

// AddEventListener fn
// js:"addEventListener"
func (*ServiceWorkerContainer) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*ServiceWorkerContainer) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*ServiceWorkerContainer) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Controller prop
// js:"controller"
func (*ServiceWorkerContainer) Controller() (controller *ServiceWorker) {
	js.Rewrite("$_.controller")
	return controller
}

// Oncontrollerchange prop
// js:"oncontrollerchange"
func (*ServiceWorkerContainer) Oncontrollerchange() (oncontrollerchange func(Event)) {
	js.Rewrite("$_.oncontrollerchange")
	return oncontrollerchange
}

// SetOncontrollerchange prop
// js:"oncontrollerchange"
func (*ServiceWorkerContainer) SetOncontrollerchange(oncontrollerchange func(Event)) {
	js.Rewrite("$_.oncontrollerchange = $1", oncontrollerchange)
}

// Onmessage prop
// js:"onmessage"
func (*ServiceWorkerContainer) Onmessage() (onmessage func(*ServiceWorkerMessageEvent)) {
	js.Rewrite("$_.onmessage")
	return onmessage
}

// SetOnmessage prop
// js:"onmessage"
func (*ServiceWorkerContainer) SetOnmessage(onmessage func(*ServiceWorkerMessageEvent)) {
	js.Rewrite("$_.onmessage = $1", onmessage)
}

// Ready prop
// js:"ready"
func (*ServiceWorkerContainer) Ready() (ready *ServiceWorkerRegistration) {
	js.Rewrite("$_.ready")
	return ready
}
