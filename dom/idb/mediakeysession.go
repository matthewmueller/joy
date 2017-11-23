package idb

import (
	"github.com/matthewmueller/golly/dom2/mediakeystatusmap"
	"github.com/matthewmueller/golly/js"
)

// MediaKeySession struct
// js:"MediaKeySession,omit"
type MediaKeySession struct {
	EventTarget
}

// Close
func (*MediaKeySession) Close() {
	js.Rewrite("await $<.Close()")
}

// GenerateRequest
func (*MediaKeySession) GenerateRequest(initDataType string, initData []byte) {
	js.Rewrite("await $<.GenerateRequest($1, $2)", initDataType, initData)
}

// Load
func (*MediaKeySession) Load(sessionId string) (b bool) {
	js.Rewrite("await $<.Load($1)", sessionId)
	return b
}

// Remove
func (*MediaKeySession) Remove() {
	js.Rewrite("await $<.Remove()")
}

// Update
func (*MediaKeySession) Update(response []byte) {
	js.Rewrite("await $<.Update($1)", response)
}

func (*MediaKeySession) Closed() {
	js.Rewrite("await $<.Closed")
}

// Expiration
func (*MediaKeySession) Expiration() (expiration float32) {
	js.Rewrite("$<.Expiration")
	return expiration
}

// KeyStatuses
func (*MediaKeySession) KeyStatuses() (keyStatuses *mediakeystatusmap.MediaKeyStatusMap) {
	js.Rewrite("$<.KeyStatuses")
	return keyStatuses
}

// SessionID
func (*MediaKeySession) SessionID() (sessionId string) {
	js.Rewrite("$<.SessionID")
	return sessionId
}
