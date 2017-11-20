package serviceworker

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/abstractworker"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/serviceworkerstate"
	"github.com/matthewmueller/golly/js"
)

type ServiceWorker struct {
	*eventtarget.EventTarget
	*abstractworker.AbstractWorker
}

func (*ServiceWorker) PostMessage(message interface{}, transfer []interface{}) {
	js.Rewrite("$<.postMessage($1, $2)", message, transfer)
}

func (*ServiceWorker) GetOnstatechange() (statechange *event.Event) {
	js.Rewrite("$<.onstatechange")
	return statechange
}

func (*ServiceWorker) SetOnstatechange(statechange *event.Event) {
	js.Rewrite("$<.onstatechange = $1", statechange)
}

func (*ServiceWorker) GetScriptURL() (scriptURL string) {
	js.Rewrite("$<.scriptURL")
	return scriptURL
}

func (*ServiceWorker) GetState() (state *serviceworkerstate.ServiceWorkerState) {
	js.Rewrite("$<.state")
	return state
}
