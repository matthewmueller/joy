package serviceworkermessageevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/messageport"
	"github.com/matthewmueller/golly/js"
)

type ServiceWorkerMessageEvent struct {
	*event.Event
}

func (*ServiceWorkerMessageEvent) GetData() (data interface{}) {
	js.Rewrite("$<.data")
	return data
}

func (*ServiceWorkerMessageEvent) GetLastEventID() (lastEventId string) {
	js.Rewrite("$<.lastEventId")
	return lastEventId
}

func (*ServiceWorkerMessageEvent) GetOrigin() (origin string) {
	js.Rewrite("$<.origin")
	return origin
}

func (*ServiceWorkerMessageEvent) GetPorts() (ports []*messageport.MessagePort) {
	js.Rewrite("$<.ports")
	return ports
}

func (*ServiceWorkerMessageEvent) GetSource() (source interface{}) {
	js.Rewrite("$<.source")
	return source
}
