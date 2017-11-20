package messageport

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/messageevent"
	"github.com/matthewmueller/golly/js"
)

type MessagePort struct {
	*eventtarget.EventTarget
}

func (*MessagePort) Close() {
	js.Rewrite("$<.close()")
}

func (*MessagePort) PostMessage(message interface{}, transfer []interface{}) {
	js.Rewrite("$<.postMessage($1, $2)", message, transfer)
}

func (*MessagePort) Start() {
	js.Rewrite("$<.start()")
}

func (*MessagePort) GetOnmessage() (message *messageevent.MessageEvent) {
	js.Rewrite("$<.onmessage")
	return message
}

func (*MessagePort) SetOnmessage(message *messageevent.MessageEvent) {
	js.Rewrite("$<.onmessage = $1", message)
}
