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

// Close
func (*MSMediaKeySession) Close() {
	js.Rewrite("$<.Close()")
}

// Update
func (*MSMediaKeySession) Update(key []uint8) {
	js.Rewrite("$<.Update($1)", key)
}

// Error
func (*MSMediaKeySession) Error() (err *msmediakeyerror.MSMediaKeyError) {
	js.Rewrite("$<.Error")
	return err
}

// KeySystem
func (*MSMediaKeySession) KeySystem() (keySystem string) {
	js.Rewrite("$<.KeySystem")
	return keySystem
}

// SessionID
func (*MSMediaKeySession) SessionID() (sessionId string) {
	js.Rewrite("$<.SessionID")
	return sessionId
}
