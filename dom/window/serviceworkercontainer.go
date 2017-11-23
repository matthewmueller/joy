package window

import (
	"github.com/matthewmueller/golly/dom2/registrationoptions"
	"github.com/matthewmueller/golly/js"
)

// ServiceWorkerContainer struct
// js:"ServiceWorkerContainer,omit"
type ServiceWorkerContainer struct {
	EventTarget
}

// GetRegistration fn
func (*ServiceWorkerContainer) GetRegistration(clientURL *string) (i interface{}) {
	js.Rewrite("await $<.getRegistration($1)", clientURL)
	return i
}

// GetRegistrations fn
func (*ServiceWorkerContainer) GetRegistrations() (s []*ServiceWorkerRegistration) {
	js.Rewrite("await $<.getRegistrations()")
	return s
}

// Register fn
func (*ServiceWorkerContainer) Register(scriptURL string, options *registrationoptions.RegistrationOptions) (s *ServiceWorkerRegistration) {
	js.Rewrite("await $<.register($1, $2)", scriptURL, options)
	return s
}

// Controller prop
func (*ServiceWorkerContainer) Controller() (controller *ServiceWorker) {
	js.Rewrite("$<.controller")
	return controller
}

// Oncontrollerchange prop
func (*ServiceWorkerContainer) Oncontrollerchange() (oncontrollerchange func(Event)) {
	js.Rewrite("$<.oncontrollerchange")
	return oncontrollerchange
}

// Oncontrollerchange prop
func (*ServiceWorkerContainer) SetOncontrollerchange(oncontrollerchange func(Event)) {
	js.Rewrite("$<.oncontrollerchange = $1", oncontrollerchange)
}

// Onmessage prop
func (*ServiceWorkerContainer) Onmessage() (onmessage func(*ServiceWorkerMessageEvent)) {
	js.Rewrite("$<.onmessage")
	return onmessage
}

// Onmessage prop
func (*ServiceWorkerContainer) SetOnmessage(onmessage func(*ServiceWorkerMessageEvent)) {
	js.Rewrite("$<.onmessage = $1", onmessage)
}

// Ready prop
func (*ServiceWorkerContainer) Ready() (ready *ServiceWorkerRegistration) {
	js.Rewrite("$<.ready")
	return ready
}
