package window

import (
	"github.com/matthewmueller/joy/dom/mediakeystatusmap"
	"github.com/matthewmueller/joy/js"
)

var _ EventTarget = (*MediaKeySession)(nil)

// MediaKeySession struct
// js:"MediaKeySession,omit"
type MediaKeySession struct {
}

// Close fn
// js:"close"
func (*MediaKeySession) Close() {
	js.Rewrite("await $_.close()")
}

// GenerateRequest fn
// js:"generateRequest"
func (*MediaKeySession) GenerateRequest(initDataType string, initData []byte) {
	js.Rewrite("await $_.generateRequest($1, $2)", initDataType, initData)
}

// Load fn
// js:"load"
func (*MediaKeySession) Load(sessionId string) (b bool) {
	js.Rewrite("await $_.load($1)", sessionId)
	return b
}

// Remove fn
// js:"remove"
func (*MediaKeySession) Remove() {
	js.Rewrite("await $_.remove()")
}

// Update fn
// js:"update"
func (*MediaKeySession) Update(response []byte) {
	js.Rewrite("await $_.update($1)", response)
}

// AddEventListener fn
// js:"addEventListener"
func (*MediaKeySession) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MediaKeySession) DispatchEvent(evt Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MediaKeySession) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Closed prop
// js:"closed"
func (*MediaKeySession) Closed() {
	js.Rewrite("await $_.closed")
}

// Expiration prop
// js:"expiration"
func (*MediaKeySession) Expiration() (expiration float32) {
	js.Rewrite("$_.expiration")
	return expiration
}

// KeyStatuses prop
// js:"keyStatuses"
func (*MediaKeySession) KeyStatuses() (keyStatuses *mediakeystatusmap.MediaKeyStatusMap) {
	js.Rewrite("$_.keyStatuses")
	return keyStatuses
}

// SessionID prop
// js:"sessionId"
func (*MediaKeySession) SessionID() (sessionId string) {
	js.Rewrite("$_.sessionId")
	return sessionId
}
