package serviceworkercontainer

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/registrationoptions"
	"github.com/matthewmueller/golly/internal/gendom/dom/serviceworker"
	"github.com/matthewmueller/golly/internal/gendom/dom/serviceworkermessageevent"
	"github.com/matthewmueller/golly/internal/gendom/dom/serviceworkerregistration"
	"github.com/matthewmueller/golly/js"
)

type ServiceWorkerContainer struct {
	*eventtarget.EventTarget
}

func (*ServiceWorkerContainer) GetRegistration(clientURL string) (i interface{}) {
	js.Rewrite("await $<.getRegistration($1)", clientURL)
	return i
}

func (*ServiceWorkerContainer) GetRegistrations() (s []*serviceworkerregistration.ServiceWorkerRegistration) {
	js.Rewrite("await $<.getRegistrations()")
	return s
}

func (*ServiceWorkerContainer) Register(scriptURL string, options *registrationoptions.RegistrationOptions) (s *serviceworkerregistration.ServiceWorkerRegistration) {
	js.Rewrite("await $<.register($1, $2)", scriptURL, options)
	return s
}

func (*ServiceWorkerContainer) GetController() (controller *serviceworker.ServiceWorker) {
	js.Rewrite("$<.controller")
	return controller
}

func (*ServiceWorkerContainer) GetOncontrollerchange() (controllerchange *event.Event) {
	js.Rewrite("$<.oncontrollerchange")
	return controllerchange
}

func (*ServiceWorkerContainer) SetOncontrollerchange(controllerchange *event.Event) {
	js.Rewrite("$<.oncontrollerchange = $1", controllerchange)
}

func (*ServiceWorkerContainer) GetOnmessage() (message *serviceworkermessageevent.ServiceWorkerMessageEvent) {
	js.Rewrite("$<.onmessage")
	return message
}

func (*ServiceWorkerContainer) SetOnmessage(message *serviceworkermessageevent.ServiceWorkerMessageEvent) {
	js.Rewrite("$<.onmessage = $1", message)
}

func (*ServiceWorkerContainer) GetReady() (ready *serviceworkerregistration.ServiceWorkerRegistration) {
	js.Rewrite("$<.ready")
	return ready
}
