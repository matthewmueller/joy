package msmediakeysession

import (
	"github.com/matthewmueller/golly/dom/msmediakeyerror"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.EventTarget = (*MSMediaKeySession)(nil)

// MSMediaKeySession struct
// js:"MSMediaKeySession,omit"
type MSMediaKeySession struct {
}

// Close fn
// js:"close"
func (*MSMediaKeySession) Close() {
	js.Rewrite("$_.close()")
}

// Update fn
// js:"update"
func (*MSMediaKeySession) Update(key []uint8) {
	js.Rewrite("$_.update($1)", key)
}

// AddEventListener fn
// js:"addEventListener"
func (*MSMediaKeySession) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MSMediaKeySession) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MSMediaKeySession) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Error prop
// js:"error"
func (*MSMediaKeySession) Error() (err *msmediakeyerror.MSMediaKeyError) {
	js.Rewrite("$_.error")
	return err
}

// KeySystem prop
// js:"keySystem"
func (*MSMediaKeySession) KeySystem() (keySystem string) {
	js.Rewrite("$_.keySystem")
	return keySystem
}

// SessionID prop
// js:"sessionId"
func (*MSMediaKeySession) SessionID() (sessionId string) {
	js.Rewrite("$_.sessionId")
	return sessionId
}
