package window

import "github.com/matthewmueller/golly/js"

// NewServiceWorkerMessageEvent fn
func NewServiceWorkerMessageEvent(kind string, eventinitdict *ServiceWorkerMessageEventInit) *ServiceWorkerMessageEvent {
	js.Rewrite("ServiceWorkerMessageEvent")
	return &ServiceWorkerMessageEvent{}
}

// ServiceWorkerMessageEvent struct
// js:"ServiceWorkerMessageEvent,omit"
type ServiceWorkerMessageEvent struct {
	Event
}

// Data prop
func (*ServiceWorkerMessageEvent) Data() (data interface{}) {
	js.Rewrite("$<.data")
	return data
}

// LastEventID prop
func (*ServiceWorkerMessageEvent) LastEventID() (lastEventId string) {
	js.Rewrite("$<.lastEventId")
	return lastEventId
}

// Origin prop
func (*ServiceWorkerMessageEvent) Origin() (origin string) {
	js.Rewrite("$<.origin")
	return origin
}

// Ports prop
func (*ServiceWorkerMessageEvent) Ports() (ports []*MessagePort) {
	js.Rewrite("$<.ports")
	return ports
}

// Source prop
func (*ServiceWorkerMessageEvent) Source() (source interface{}) {
	js.Rewrite("$<.source")
	return source
}
