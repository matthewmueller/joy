package window

import (
	"github.com/matthewmueller/golly/dom2/serviceworkerstate"
	"github.com/matthewmueller/golly/js"
)

// ServiceWorker struct
// js:"ServiceWorker,omit"
type ServiceWorker struct {
	EventTarget
}

// PostMessage
func (*ServiceWorker) PostMessage(message interface{}, transfer *[]interface{}) {
	js.Rewrite("$<.PostMessage($1, $2)", message, transfer)
}

// Onerror
func (*ServiceWorker) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*ServiceWorker) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onstatechange
func (*ServiceWorker) Onstatechange() (onstatechange func(Event)) {
	js.Rewrite("$<.Onstatechange")
	return onstatechange
}

// Onstatechange
func (*ServiceWorker) SetOnstatechange(onstatechange func(Event)) {
	js.Rewrite("$<.Onstatechange = $1", onstatechange)
}

// ScriptURL
func (*ServiceWorker) ScriptURL() (scriptURL string) {
	js.Rewrite("$<.ScriptURL")
	return scriptURL
}

// State
func (*ServiceWorker) State() (state *serviceworkerstate.ServiceWorkerState) {
	js.Rewrite("$<.State")
	return state
}
