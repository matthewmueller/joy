package window

import (
	"github.com/matthewmueller/joy/dom/serviceworkerstate"
	"github.com/matthewmueller/joy/macro"
)

var _ EventTarget = (*ServiceWorker)(nil)

// ServiceWorker struct
// js:"ServiceWorker,omit"
type ServiceWorker struct {
}

// PostMessage fn
// js:"postMessage"
func (*ServiceWorker) PostMessage(message interface{}, transfer *[]interface{}) {
	macro.Rewrite("$_.postMessage($1, $2)", message, transfer)
}

// AddEventListener fn
// js:"addEventListener"
func (*ServiceWorker) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*ServiceWorker) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*ServiceWorker) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Onstatechange prop
// js:"onstatechange"
func (*ServiceWorker) Onstatechange() (onstatechange func(Event)) {
	macro.Rewrite("$_.onstatechange")
	return onstatechange
}

// SetOnstatechange prop
// js:"onstatechange"
func (*ServiceWorker) SetOnstatechange(onstatechange func(Event)) {
	macro.Rewrite("$_.onstatechange = $1", onstatechange)
}

// ScriptURL prop
// js:"scriptURL"
func (*ServiceWorker) ScriptURL() (scriptURL string) {
	macro.Rewrite("$_.scriptURL")
	return scriptURL
}

// State prop
// js:"state"
func (*ServiceWorker) State() (state *serviceworkerstate.ServiceWorkerState) {
	macro.Rewrite("$_.state")
	return state
}

// Onerror prop
// js:"onerror"
func (*ServiceWorker) Onerror() (onerror func(Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*ServiceWorker) SetOnerror(onerror func(Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}
