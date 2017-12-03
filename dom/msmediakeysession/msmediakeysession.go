package msmediakeysession

import (
	"github.com/matthewmueller/joy/dom/msmediakeyerror"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.EventTarget = (*MSMediaKeySession)(nil)

// MSMediaKeySession struct
// js:"MSMediaKeySession,omit"
type MSMediaKeySession struct {
}

// Close fn
// js:"close"
func (*MSMediaKeySession) Close() {
	macro.Rewrite("$_.close()")
}

// Update fn
// js:"update"
func (*MSMediaKeySession) Update(key []uint8) {
	macro.Rewrite("$_.update($1)", key)
}

// AddEventListener fn
// js:"addEventListener"
func (*MSMediaKeySession) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MSMediaKeySession) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MSMediaKeySession) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Error prop
// js:"error"
func (*MSMediaKeySession) Error() (err *msmediakeyerror.MSMediaKeyError) {
	macro.Rewrite("$_.error")
	return err
}

// KeySystem prop
// js:"keySystem"
func (*MSMediaKeySession) KeySystem() (keySystem string) {
	macro.Rewrite("$_.keySystem")
	return keySystem
}

// SessionID prop
// js:"sessionId"
func (*MSMediaKeySession) SessionID() (sessionId string) {
	macro.Rewrite("$_.sessionId")
	return sessionId
}
