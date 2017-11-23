package window

import (
	"github.com/matthewmueller/golly/dom/serviceworkerstate"
	"github.com/matthewmueller/golly/js"
)

// ServiceWorker struct
// js:"ServiceWorker,omit"
type ServiceWorker struct {
	EventTarget
}

// PostMessage fn
func (*ServiceWorker) PostMessage(message interface{}, transfer *[]interface{}) {
	js.Rewrite("$<.postMessage($1, $2)", message, transfer)
}

// Onerror prop
func (*ServiceWorker) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*ServiceWorker) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onstatechange prop
func (*ServiceWorker) Onstatechange() (onstatechange func(Event)) {
	js.Rewrite("$<.onstatechange")
	return onstatechange
}

// Onstatechange prop
func (*ServiceWorker) SetOnstatechange(onstatechange func(Event)) {
	js.Rewrite("$<.onstatechange = $1", onstatechange)
}

// ScriptURL prop
func (*ServiceWorker) ScriptURL() (scriptURL string) {
	js.Rewrite("$<.scriptURL")
	return scriptURL
}

// State prop
func (*ServiceWorker) State() (state *serviceworkerstate.ServiceWorkerState) {
	js.Rewrite("$<.state")
	return state
}
