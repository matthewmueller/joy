package window

import "github.com/matthewmueller/golly/js"

// js:"ServiceWorkerMessageEvent,omit"
type ServiceWorkerMessageEvent struct {
	Event
}

// Data
func (*ServiceWorkerMessageEvent) Data() (data interface{}) {
	js.Rewrite("$<.Data")
	return data
}

// LastEventID
func (*ServiceWorkerMessageEvent) LastEventID() (lastEventId string) {
	js.Rewrite("$<.LastEventID")
	return lastEventId
}

// Origin
func (*ServiceWorkerMessageEvent) Origin() (origin string) {
	js.Rewrite("$<.Origin")
	return origin
}

// Ports
func (*ServiceWorkerMessageEvent) Ports() (ports []*MessagePort) {
	js.Rewrite("$<.Ports")
	return ports
}

// Source
func (*ServiceWorkerMessageEvent) Source() (source interface{}) {
	js.Rewrite("$<.Source")
	return source
}
