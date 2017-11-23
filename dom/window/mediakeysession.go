package window

import (
	"github.com/matthewmueller/golly/dom2/mediakeystatusmap"
	"github.com/matthewmueller/golly/js"
)

// MediaKeySession struct
// js:"MediaKeySession,omit"
type MediaKeySession struct {
	EventTarget
}

// Close fn
func (*MediaKeySession) Close() {
	js.Rewrite("await $<.close()")
}

// GenerateRequest fn
func (*MediaKeySession) GenerateRequest(initDataType string, initData []byte) {
	js.Rewrite("await $<.generateRequest($1, $2)", initDataType, initData)
}

// Load fn
func (*MediaKeySession) Load(sessionId string) (b bool) {
	js.Rewrite("await $<.load($1)", sessionId)
	return b
}

// Remove fn
func (*MediaKeySession) Remove() {
	js.Rewrite("await $<.remove()")
}

// Update fn
func (*MediaKeySession) Update(response []byte) {
	js.Rewrite("await $<.update($1)", response)
}

func (*MediaKeySession) Closed() {
	js.Rewrite("await $<.closed")
}

// Expiration prop
func (*MediaKeySession) Expiration() (expiration float32) {
	js.Rewrite("$<.expiration")
	return expiration
}

// KeyStatuses prop
func (*MediaKeySession) KeyStatuses() (keyStatuses *mediakeystatusmap.MediaKeyStatusMap) {
	js.Rewrite("$<.keyStatuses")
	return keyStatuses
}

// SessionID prop
func (*MediaKeySession) SessionID() (sessionId string) {
	js.Rewrite("$<.sessionId")
	return sessionId
}
