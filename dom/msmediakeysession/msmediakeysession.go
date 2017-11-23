package msmediakeysession

import (
	"github.com/matthewmueller/golly/dom2/msmediakeyerror"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// MSMediaKeySession struct
// js:"MSMediaKeySession,omit"
type MSMediaKeySession struct {
	window.EventTarget
}

// Close fn
func (*MSMediaKeySession) Close() {
	js.Rewrite("$<.close()")
}

// Update fn
func (*MSMediaKeySession) Update(key []uint8) {
	js.Rewrite("$<.update($1)", key)
}

// Error prop
func (*MSMediaKeySession) Error() (err *msmediakeyerror.MSMediaKeyError) {
	js.Rewrite("$<.error")
	return err
}

// KeySystem prop
func (*MSMediaKeySession) KeySystem() (keySystem string) {
	js.Rewrite("$<.keySystem")
	return keySystem
}

// SessionID prop
func (*MSMediaKeySession) SessionID() (sessionId string) {
	js.Rewrite("$<.sessionId")
	return sessionId
}
