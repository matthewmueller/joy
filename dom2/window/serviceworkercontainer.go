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

// GetRegistration
func (*ServiceWorkerContainer) GetRegistration(clientURL *string) (i interface{}) {
	js.Rewrite("await $<.GetRegistration($1)", clientURL)
	return i
}

// GetRegistrations
func (*ServiceWorkerContainer) GetRegistrations() (s []*ServiceWorkerRegistration) {
	js.Rewrite("await $<.GetRegistrations()")
	return s
}

// Register
func (*ServiceWorkerContainer) Register(scriptURL string, options *registrationoptions.RegistrationOptions) (s *ServiceWorkerRegistration) {
	js.Rewrite("await $<.Register($1, $2)", scriptURL, options)
	return s
}

// Controller
func (*ServiceWorkerContainer) Controller() (controller *ServiceWorker) {
	js.Rewrite("$<.Controller")
	return controller
}

// Oncontrollerchange
func (*ServiceWorkerContainer) Oncontrollerchange() (oncontrollerchange func(Event)) {
	js.Rewrite("$<.Oncontrollerchange")
	return oncontrollerchange
}

// Oncontrollerchange
func (*ServiceWorkerContainer) SetOncontrollerchange(oncontrollerchange func(Event)) {
	js.Rewrite("$<.Oncontrollerchange = $1", oncontrollerchange)
}

// Onmessage
func (*ServiceWorkerContainer) Onmessage() (onmessage func(*ServiceWorkerMessageEvent)) {
	js.Rewrite("$<.Onmessage")
	return onmessage
}

// Onmessage
func (*ServiceWorkerContainer) SetOnmessage(onmessage func(*ServiceWorkerMessageEvent)) {
	js.Rewrite("$<.Onmessage = $1", onmessage)
}

// Ready
func (*ServiceWorkerContainer) Ready() (ready *ServiceWorkerRegistration) {
	js.Rewrite("$<.Ready")
	return ready
}
