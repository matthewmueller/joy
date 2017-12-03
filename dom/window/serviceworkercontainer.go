package window

import (
	"github.com/matthewmueller/joy/dom/registrationoptions"
	"github.com/matthewmueller/joy/macro"
)

var _ EventTarget = (*ServiceWorkerContainer)(nil)

// ServiceWorkerContainer struct
// js:"ServiceWorkerContainer,omit"
type ServiceWorkerContainer struct {
}

// GetRegistration fn
// js:"getRegistration"
func (*ServiceWorkerContainer) GetRegistration(clientURL *string) (i interface{}) {
	macro.Rewrite("await $_.getRegistration($1)", clientURL)
	return i
}

// GetRegistrations fn
// js:"getRegistrations"
func (*ServiceWorkerContainer) GetRegistrations() (s []*ServiceWorkerRegistration) {
	macro.Rewrite("await $_.getRegistrations()")
	return s
}

// Register fn
// js:"register"
func (*ServiceWorkerContainer) Register(scriptURL string, options *registrationoptions.RegistrationOptions) (s *ServiceWorkerRegistration) {
	macro.Rewrite("await $_.register($1, $2)", scriptURL, options)
	return s
}

// AddEventListener fn
// js:"addEventListener"
func (*ServiceWorkerContainer) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*ServiceWorkerContainer) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*ServiceWorkerContainer) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Controller prop
// js:"controller"
func (*ServiceWorkerContainer) Controller() (controller *ServiceWorker) {
	macro.Rewrite("$_.controller")
	return controller
}

// Oncontrollerchange prop
// js:"oncontrollerchange"
func (*ServiceWorkerContainer) Oncontrollerchange() (oncontrollerchange func(Event)) {
	macro.Rewrite("$_.oncontrollerchange")
	return oncontrollerchange
}

// SetOncontrollerchange prop
// js:"oncontrollerchange"
func (*ServiceWorkerContainer) SetOncontrollerchange(oncontrollerchange func(Event)) {
	macro.Rewrite("$_.oncontrollerchange = $1", oncontrollerchange)
}

// Onmessage prop
// js:"onmessage"
func (*ServiceWorkerContainer) Onmessage() (onmessage func(*ServiceWorkerMessageEvent)) {
	macro.Rewrite("$_.onmessage")
	return onmessage
}

// SetOnmessage prop
// js:"onmessage"
func (*ServiceWorkerContainer) SetOnmessage(onmessage func(*ServiceWorkerMessageEvent)) {
	macro.Rewrite("$_.onmessage = $1", onmessage)
}

// Ready prop
// js:"ready"
func (*ServiceWorkerContainer) Ready() (ready *ServiceWorkerRegistration) {
	macro.Rewrite("$_.ready")
	return ready
}
