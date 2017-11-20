package worker

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/abstractworker"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/messageevent"
	"github.com/matthewmueller/golly/js"
)

type Worker struct {
	*eventtarget.EventTarget
	*abstractworker.AbstractWorker
}

func (*Worker) PostMessage(message interface{}, transfer []interface{}) {
	js.Rewrite("$<.postMessage($1, $2)", message, transfer)
}

func (*Worker) Terminate() {
	js.Rewrite("$<.terminate()")
}

func (*Worker) GetOnmessage() (message *messageevent.MessageEvent) {
	js.Rewrite("$<.onmessage")
	return message
}

func (*Worker) SetOnmessage(message *messageevent.MessageEvent) {
	js.Rewrite("$<.onmessage = $1", message)
}
